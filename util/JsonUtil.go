package util

import (
	"WT/entry"
	"encoding/json"
	"os"
)

// ToDeploy 配置文件读取
func ToDeploy() (*entry.Deploy,error) {
	open, err := os.ReadFile("conf.json")
	if err != nil {
		return nil,err
	}
	print(string(open))
	deploy := &entry.Deploy{}
	err = json.Unmarshal(open,&deploy)
	if err != nil {
		return nil,err
	}
	return deploy,nil
}


