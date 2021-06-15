package main

import (
	"io"
	"os"
	strings "strings"
)

// FileAcquisition 文件内容获取
func FileAcquisition(path string) (string, error) {
	// 指定用户输入文件路径
	fileUrl := path
	file, err := os.Open(fileUrl)
	if err != nil {
		return "", err
	}
	// 延迟file Close
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	// 获取输出流
	all, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(all), err
}

// TextParsing 字符串内容解析
func TextParsing(report string) []string {
// 分割字符串 f:\\1.txt
	reportArr := strings.Split(report,"今日完成:")
	var context = make([]string,5)
	i:=0
	for _ , lineOfText:= range reportArr {
		if lineOfText !="" {
			lineOfText = strings.Replace(lineOfText, "\r\n", "", 1)
			context[i]= lineOfText
			i+=1
		}
	}
	for i := 0; i < len(context); i++ {
		println(context[i])
		println(i)
	}
	return context
}

// TextParsingRewrite 采用分割获取内容
func TextParsingRewrite(report string) ([]string, string) {
	// 分割字符串 f:\\1.txt
	reportArr := strings.Split(report,"\r\n")
	context := make([]string,5)
	tag:=0
	for i := 0; i < len(reportArr); i++ {
		// 定义起点
		if reportArr[i] == "今日完成:" {
			for j := i+1; j <len(reportArr); j++ {
				//定义终点
				if reportArr[j] == "今日完成:" {
					i=j-1
					break
				}
				if context[tag] == "" {
					context[tag]+=reportArr[j]
					continue
				}
				context[tag]=context[tag]+"\r\n"+reportArr[j]
			}
			tag++
		}
	}

	allContext:=""
	for index, c := range context {
		if index==0 {
			allContext+=c
			continue
		}
		allContext+="\r\n"+c
	}
	return context,allContext
}


