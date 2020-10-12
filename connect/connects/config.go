package connects

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct{
	Apikey string
	Apisecret string
}

var Config ConfigList

func Configini(){
	cfg,err := ini.Load("connects/config.ini")
	if err != nil{
		fmt.Println(err)
	}
	Config = ConfigList{
		Apikey : cfg.Section("web").Key("api_key").String(),
		Apisecret : cfg.Section("web").Key("api_secret").String(),
	}
	fmt.Println(Config)
}
