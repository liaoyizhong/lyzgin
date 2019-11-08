/*
@Time : 2019/9/29 19:51 
@Author : liaoyizhong
@File : register
@Software: GoLand
*/
package wgin

//注册程序
func RegisterFun(name string, ff func()error){
	temFuc := InitFunc{Fun:ff, Name:name}
	InitFuncList = append(InitFuncList, temFuc)
}


