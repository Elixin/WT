package main

import (
	"WT/entry"
	"WT/util"
)

func errorPrint(err error) {
	if err != nil {
		print(err.Error())
		return
	}
}

func main() {

	deploy, err := util.ToDeploy()
	errorPrint(err)

	week := util.ObtainWorkDayToCurrentWeek()

	//文件读取
	acquisition, err := FileAcquisition(deploy.FilePath)
	errorPrint(err)
	// 文本提取
	parsing, allContext := TextParsingRewrite(acquisition)
	tableContent := entry.TableContent{}
	tableContent = tableContent.SetValueInEntry(deploy, week, parsing, allContext)
	util.ExcelRead(&tableContent,*deploy)
	errorPrint(err)
	// 创建邮箱发送地址
	// 写入发送内容模板
	// 指定发送人
	// 写入定时器

}
