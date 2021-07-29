package util

import (
	"WT/entry"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"os"
	reflect "reflect"
	"strconv"
	"strings"
	"time"
)

// ExcelCreate 创建Excel操作文件类
func ExcelCreate(path string) (*excelize.File,string) {
	files, err := ioutil.ReadDir(path)
	Errors(err)
	var max int
	var changeTime time.Time
	for i, file := range files {
		split := strings.Split(file.Name(), ".")
		if split[1] == "xlsx" || split[1] == "xls" {
			if time.Time.IsZero(changeTime) {
				changeTime = file.ModTime()
				max = i
			} else {
				if file.ModTime().Before(changeTime) {
					changeTime = file.ModTime()
					max = i
				}
			}
		}
	}
	path += "/" + files[max].Name()
	file, err := excelize.OpenFile(path)
	Errors(err)
	return file,path
}

func ExcelCreateMode(fileName string) *excelize.File {
	file, err := excelize.OpenFile(fileName)
	Errors(err)
	return file
}

func fileName(firstDay string,endDay string,currentMonth string,currentWeek string,author string) string {
	return firstDay+"-"+endDay+"-"+currentMonth+currentWeek+"工作周报("+author+").xlsx"
}

func ExcelRead(content *entry.TableContent, deploy entry.Deploy) {
	// 获取模板文件
	var create *excelize.File
	path:=""
	if deploy.ModePath == "" {
		create ,path= ExcelCreate(deploy.OutPath)
	} else {
		create = ExcelCreateMode(deploy.ModePath)
	}

	// 获取写入文件
	nowMonth :=  strconv.Itoa(int(time.Now().Month())) + "月"
	fileName := fileName(content.WorkingDay[0],content.WorkingDay[len(content.WorkingDay)-1],nowMonth,content.NowWeek,deploy.Author)
	nowSheetName := nowMonth
	defaultSheetName := "Sheet1"
	// 获取模板所有数据
	rows, err := create.GetRows(defaultSheetName)
	Errors(err)
	// 获取现有文件所有数据
	create.NewSheet(nowSheetName)
	getRows, err := create.GetRows(nowSheetName)
	Errors(err)
	// 获取起点
	starting := len(getRows) + 1
	// 获取模板样式以及数据
	for i, row := range rows {
		for j, colCell := range row {
			axis, err := excelize.CoordinatesToCellName(j+1, i+1)
			Errors(err)
			// 获取当前单元格样式
			style, modeHeight, colName, modeWidth, err := getModeCellStyle(create, defaultSheetName, axis, i, j)
			Errors(err)
			cpyAxis, err := excelize.CoordinatesToCellName(j+1, starting+i+1)
			Errors(err)
			nowRowNum := starting + i + 1
			setCellStyle(nowSheetName, create, nowRowNum, modeHeight, colName, modeWidth)
			// 写入模板对应数据
			var value interface{}
			if strings.ContainsAny(colCell, "$") {
				value = WriteFileContent(content, colCell)
			} else {
				value = colCell
			}

			err = create.SetCellValue(nowSheetName, cpyAxis, value)
			Errors(err)
			// 写入模板对应样式
			err = create.SetCellStyle(nowSheetName, cpyAxis, cpyAxis, style)
			Errors(err)

		}
	}
	// 获取模板文件合并单元格文件
	mergeCells(err, create, defaultSheetName, starting, nowSheetName)

	// 设置sheet为默认
	create.SetActiveSheet(create.GetSheetIndex(nowSheetName))
	err = create.SetSheetVisible(defaultSheetName, false)
	Errors(err)
	if deploy.ModePath == "" {
		err = create.Save()
		err = os.Rename(path, deploy.OutPath+"/"+fileName)
		Errors(err)
	}
	if deploy.ModePath != "" {
		err := create.SaveAs(deploy.OutPath + "/" + fileName)
		Errors(err)
	}

}

// 获取模板单元格样式
func getModeCellStyle(create *excelize.File, defaultSheetName string, axis string, i int, j int) (int, float64, string, float64, error) {
	style, err := create.GetCellStyle(defaultSheetName, axis)
	Errors(err)
	modeHeight, err := create.GetRowHeight(defaultSheetName, i+1)
	Errors(err)
	colName, err := excelize.ColumnNumberToName(j + 1)
	Errors(err)
	modeWidth, err := create.GetColWidth(defaultSheetName, colName)
	return style, modeHeight, colName, modeWidth, nil
}

// 设置单元格合并
func mergeCells(err error, create *excelize.File, defaultSheetName string, starting int, nowSheetName string) {
	cells, err := create.GetMergeCells(defaultSheetName)
	Errors(err)
	// 设置模板合并单元格到写入文件
	for _, cell := range cells {
		x1, y1, err := excelize.CellNameToCoordinates(cell.GetStartAxis())
		Errors(err)
		x2, y2, err := excelize.CellNameToCoordinates(cell.GetEndAxis())
		Errors(err)
		y1 += starting
		y2 += starting
		start, err := excelize.CoordinatesToCellName(x1, y1)
		Errors(err)
		end, err := excelize.CoordinatesToCellName(x2, y2)
		Errors(err)
		err = create.MergeCell(nowSheetName, start, end)
		Errors(err)
	}
}

// 设置单元格样式
func setCellStyle(sheetName string, create *excelize.File, nowRowNum int, modeHeight float64, colName string, modeWidth float64) {
	err := create.SetRowHeight(sheetName, nowRowNum, modeHeight)
	Errors(err)
	err = create.SetColWidth(sheetName, colName, colName, modeWidth)
	Errors(err)

}

// WriteFileContent 文件内容写入
func WriteFileContent(content *entry.TableContent, tableFieldName string) (value interface{}) {
	// 反射获取实体类字段名
	elem := reflect.ValueOf(content).Elem()
	tableFieldName = strings.Split(tableFieldName, "$")[1]
	// 从结构体中获取切片
	if strings.ContainsAny(tableFieldName, "_") {
		split := strings.Split(tableFieldName, "_")
		index, err := strconv.Atoi(split[1])
		Errors(err)
		value := elem.FieldByName(split[0]).Interface()
		of := reflect.ValueOf(value)
		return of.Index(index - 1)

	}
	// 通过字段名获取属性值
	field := elem.FieldByName(tableFieldName)
	return field
}
