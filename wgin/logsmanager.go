package wgin

import (
	"github.com/astaxie/beego/logs"
)

func LogsInfo(textInfo string){
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/test.log"}`)
	logs.Info(textInfo)
}
