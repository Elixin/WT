package util

import (
	"WT/entry"
	"encoding/json"
)

// ToDeploy 配置文件读取
func ConfigToDeploy(jsonByte []byte) (*entry.Deploy) {
	deploy := &entry.Deploy{}
	err := json.Unmarshal(jsonByte,&deploy)
	Errors(err)
	return deploy
}


