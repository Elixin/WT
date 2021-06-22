package util

import (
	"WT/entry"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
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
		index := strings.Index(s, "=")
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
		strings.Index(value,"，")
		if index > 0 {
			_ = fmt.Errorf("不能使用中文表点")
		}
		strings.Index(value,",")

		if index > 0 {
			aValue := strings.Split(value,",")
			config[key] = aValue
		}else {
			config[key] = value
		}

	}
	marshal, err := json.Marshal(config)
	Errors(err)
	deploy := ConfigToDeploy(marshal)
	return deploy
}
// 获取




// InputEntry 将配置文件写入实体类
func InputEntry(config map[string]string) (*entry.Deploy,error) {


	return nil,nil
}