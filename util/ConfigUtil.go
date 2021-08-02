package util

import (
	"WT/entry"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	strconv "strconv"
	strings "strings"
)

//读取key=value类型的配置文件
func InitConfig(path string) *entry.Deploy {
	config := make(map[string]interface{})

	f, err := os.Open(path)
	defer f.Close()
	Errors(err)

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		// 是否是注释
		index := strings.Index(s, "#")
		if index>0 {
			continue
		}
		index = strings.Index(s, "=")
		if index < 0 {
			continue
		}


		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		hasCP := strings.Index(value, "，")
		if hasCP > 0 {
			_ = fmt.Errorf("不能使用中文表点")
		}
		hasComma := strings.Index(value, ",")
		prefixs := strings.HasPrefix(value, "...")
		prefixn := strings.HasPrefix(value, "F...")
		prefixf := strings.HasPrefix(value, "N...")
		if hasComma > 0 || prefixs||prefixn||prefixf{
			array := getArray(value)
			if array==nil {
				continue
			}
			config[key] = array
		}else {
			num := isNum(value)
			if num!=nil {
				config[key] = num
				continue
			}
			config[key] = value
		}


	}

	marshal, err := json.Marshal(config)


	deploy := ConfigToDeploy(marshal)
	return deploy
}

func getArray(value string) interface{}{
	split := strings.Split(value, ",")
	if len(split)>1 {
		switch split[0] {
		case "N":
			// 整形数组
			size := len(split)-1
			intArray:=make([]int,size)
			for i := 0; i < len(split); i++ {
				atoi, err := strconv.Atoi(split[i+1])
				if err != nil {
					Errors(err)
				}
				intArray[i] = atoi
			}
			return intArray
		case "F":
			// 浮点数组
			size := len(split) - 1
			floatArray := make([]float64, size)
			for i := 0; i < len(split); i++ {
				float, err := strconv.ParseFloat(split[i+1], 64)
				if err != nil {
					Errors(err)
				}
				floatArray[i] = float
			}
			return floatArray
		}
	}else {
		// 建立重复字段数组
		split := strings.Split(value, "...")
		switch split[0] {
		case "":
			// 字符串
			stringArray := make([]string, 5)
			for i := 0; i < len(stringArray); i++ {
				stringArray[i] = split[1]
			}
			return stringArray
		case "N":
			// 整形
			intArray := make([]int, 5)
			for i := 0; i < len(intArray); i++ {
				atoi, err := strconv.Atoi(split[1])
				if err != nil {
					Errors(err)
				}
				intArray[i] = atoi
			}
			return intArray
		case "F":
			// 浮点数
			floatArray := make([]float64, 5)
			for i := 0; i < len(floatArray); i++ {
				float, err := strconv.ParseFloat(split[1],64)
				if err != nil {
					Errors(err)
				}
				floatArray[i] = float
			}
			return floatArray
		}
	}
	return split
}


// 获取

func isNum(str string) interface{} {
	if strings.HasSuffix(str, "-s") {
		return strings.Split(str,"-s")[0]
	}
	Int, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return Int
	}
	Float, err := strconv.ParseFloat(str,64)
	if err != nil {
		return nil
	}
	return Float
}





// InputEntry 将配置文件写入实体类
func InputEntry(config map[string]string) (*entry.Deploy,error) {


	return nil,nil
}