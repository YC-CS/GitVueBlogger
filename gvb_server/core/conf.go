package core

import (
	"encoding/json"
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.json"
	c := &config.Config{}
	jsonConfig, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get jsonConfig error: %s", err))
	}
	err = json.Unmarshal(jsonConfig, c)
	if err != nil {
		log.Fatalf("config init Unmarashal: %s", err)
	}
	log.Println("config jsonFile load Init success.")
	global.Config = c
}
