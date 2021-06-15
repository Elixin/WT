package util

import (
	"WT/entry"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	reflect "reflect"
	"strconv"
	"strings"
	"time"
)

// ExcelCreate 创建Excel操作文件类
func ExcelCreate(path string) *excelize.File {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
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
	path += "/"+files[max].Name()
	file, err := excelize.OpenFile(path)
	if err != nil {
		print(err)
		return nil
	}
	return file
}
func ExcelCreateMode(fileName string) *excelize.File {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		print(err)
		return nil
	}
	return file
}
func ExcelRead(content *entry.TableContent, deploy entry.Deploy) {
	// 获取模板文件
	var create *excelize.File
	if deploy.ModePath == "" {
		print(deploy.OutPath)
		create = ExcelCreate(deploy.OutPath)
	}else {
		print(deploy.ModePath)
		create = ExcelCreateMode(deploy.ModePath)
	}

	// 获取写入文件
	nowMonth := strconv.Itoa(int(time.Now().Month())) + "月"
	fileName := content.WorkingDay[0] + "-" + content.WorkingDay[len(content.WorkingDay)-1] + " - " + nowMonth + content.NowWeek + "工作周报(" + deploy.Author + ").xlsx"
	nowSheetName := nowMonth
	defaultSheetName := "Sheet1"
	// 获取模板所有数据
	rows, err := create.GetRows(defaultSheetName)
	if err != nil {
		return
	}
	// 获取现有文件所有数据
	create.NewSheet(nowSheetName)
	getRows, err := create.GetRows(nowSheetName)
	if err != nil {
		return
	}
	// 获取起点
	starting := len(getRows) + 1
	// 获取模板样式以及数据
	for i, row := range rows {
		for j, colCell := range row {
			axis, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				print(err)
				return
			}
			// 获取当前单元格样式
			style, err := create.GetCellStyle(defaultSheetName, axis)
			if err != nil {
				print(err)
				return
			}
			modeHeight, err := create.GetRowHeight(defaultSheetName, i+1)
			if err != nil {
				return
			}
			colName, err := excelize.ColumnNumberToName(j + 1)
			if err != nil {
				return
			}
			modeWidth, err := create.GetColWidth(defaultSheetName, colName)
			if err != nil {
				return
			}
			cpyAxis, err := excelize.CoordinatesToCellName(j+1, starting+i+1)
			if err != nil {
				print(err.Error())
				return
			}
			nowRowNum := starting + i + 1
			err = setCellStyle(nowSheetName, create, nowRowNum, modeHeight, colName, modeWidth)
			if err != nil {
				return
			}
			// 写入模板对应数据
			var value interface{}
			if strings.ContainsAny(colCell, "$") {
				err, value = WriteFileContent(content, colCell)
				if err != nil {
					return
				}
			} else {
				value = colCell
			}

			err = create.SetCellValue(nowSheetName, cpyAxis, value)
			if err != nil {
				print(err.Error())
				return
			}
			// 写入模板对应样式
			err = create.SetCellStyle(nowSheetName, cpyAxis, cpyAxis, style)
			if err != nil {
				print(err.Error())
				return
			}

		}
	}
	// 获取模板文件合并单元格文件
	cells, err := create.GetMergeCells(defaultSheetName)
	if err != nil {
		return
	}
	// 设置模板合并单元格到写入文件
	for _, cell := range cells {
		x1, y1, err := excelize.CellNameToCoordinates(cell.GetStartAxis())
		if err != nil {
			return
		}
		x2, y2, err := excelize.CellNameToCoordinates(cell.GetEndAxis())
		if err != nil {
			return
		}
		y1 += starting
		y2 += starting
		start, err := excelize.CoordinatesToCellName(x1, y1)
		if err != nil {
			return
		}
		end, err := excelize.CoordinatesToCellName(x2, y2)
		if err != nil {
			print(err)
			return
		}
		err = create.MergeCell(nowSheetName, start, end)
		if err != nil {
			return
		}
	}
	// 设置sheet为默认
	create.SetActiveSheet(create.GetSheetIndex(nowSheetName))
	err = create.SetSheetVisible(defaultSheetName, false)
	if err != nil {
		return
	}
	if deploy.ModePath == "" {
		err = create.Save()
		if err != nil {
			return
		}
	} else {
		err := create.SaveAs(deploy.OutPath+"/"+fileName)
		if err != nil {
			print(err.Error())
			return
		}
	}

}

// 设置单元格样式
func setCellStyle(sheetName string, create *excelize.File, nowRowNum int, modeHeight float64, colName string, modeWidth float64) error {
	err := create.SetRowHeight(sheetName, nowRowNum, modeHeight)
	if err != nil {
		print(err.Error())
		return nil
	}
	err = create.SetColWidth(sheetName, colName, colName, modeWidth)
	if err != nil {
		print(err.Error())
		return nil
	}
	return err
}

// WriteFileContent 文件内容写入
func WriteFileContent(content *entry.TableContent, tableFieldName string) (err error, value interface{}) {
	// 反射获取实体类字段名
	elem := reflect.ValueOf(content).Elem()
	typeInfo := elem.Type()
	tableFieldName = strings.Split(tableFieldName, "$")[1]
	// 从结构体中获取切片
	println(tableFieldName)
	if strings.ContainsAny(tableFieldName, "_") {
		split := strings.Split(tableFieldName, "_")
		index, err := strconv.Atoi(split[1])
		if err != nil {
			return err, ""
		}
		value := elem.FieldByName(split[0]).Interface()
		of := reflect.ValueOf(value)
		return nil, of.Index(index - 1)

	}
	name, b := typeInfo.FieldByName(tableFieldName)
	if b {
		switch name.Type.Kind().String() {
		case "int":
			value := elem.FieldByName(tableFieldName).Interface().(int)
			print(name.Name, value)
			return nil, value
		case "float64":
			value := elem.FieldByName(tableFieldName).Interface().(float64)
			print(name.Name, value)
			return nil, value
		case "string":
			value := elem.FieldByName(tableFieldName).Interface().(string)
			print(name.Name, value)
			return nil, value
		}
	}
	return nil, nil
}
