/*
@Time : 2019/9/6 12:30 
@Author : liaoyizhong
@File : wgin_com
@Software: GoLand
*/
package wgin

type WginCom struct{
	Rinit map[string]interface{}
}

func(this *WginCom)Register(name string, objProject interface{}){
	this.Rinit[name] = objProject
}
