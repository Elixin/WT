package main

import (
	"WT/entry"
	"WT/util"
)



func main() {

	deploy, err := util.ToDeploy()
	util.Errors(err)

	week := util.ObtainWorkDayToCurrentWeek()

	//文件读取
	acquisition, err := FileAcquisition(deploy.FilePath)
	util.Errors(err)
	// 文本提取
	parsing, allContext := TextParsingRewrite(acquisition)
	tableContent := entry.TableContent{}
	tableContent = tableContent.SetValueInEntry(deploy, week, parsing, allContext)
	util.ExcelRead(&tableContent,*deploy)
	errorPrint(err)

	util.Errors(err)
	// 创建邮箱发送地址
	// 写入发送内容模板
	// 指定发送人
	// 写入定时器

}
