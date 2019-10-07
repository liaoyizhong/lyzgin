package wgin

type lyzRoute struct{
	Name string
	Fun func()error
}

var ListRoute []lyzRoute

func InitRoute(){
	for _,i := range ListRoute{
		err := i.Fun()
		if err != nil{

		}
	}
}
