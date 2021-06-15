package main

import (
	"WT/entry"
	util "WT/util"
)

func errorPrint(err error)  {
	if err!=nil {
		print(err.Error())
		return
	}
}

func main() {

	//s := "今日完成:\r\n1.日常巡检6台服务器6个信息系统(100%)1.5小时\r\n2.测试服务器完成管线中心OA系统信息化需求确认功能(100%)3小时\r\n3.编写管线中心OA系统项目预算模块(5%)3.5小时\r\n今日完成:\r\n1.日常巡检6台服务器6个信息系统(100%)1.5小时\r\n2.测试服务器完成管线中心OA系统信息化需求确认功能(100%)3小时\n3.编写管线中心OA系统项目预算模块(5%)3.5小时\r\n今日完成:"
	//TextParsingRewrite(s)

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
	util.ExcelRead(&tableContent)
	errorPrint(err)
	// 创建邮箱发送地址
	// 写入发送内容模板
	// 指定发送人
	// 写入定时器


}
