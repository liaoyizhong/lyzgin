/*
@Time : 2019/9/23 12:35 
@Author : liaoyizhong
@File : registermanager
@Software: GoLand
*/
package wgin

import(
	"github.com/judwhite/go-svc/svc"
	"log"
	"sync"
)

type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

type InitFunc struct{
	Fun func()error
	Name string
}

var InitFuncList []InitFunc

// 注册器初始化
func (p *program) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())

	for _,v := range InitFuncList{
		err := v.Fun()
		if err != nil{

		}
	}

	return nil
}

// 执行方法
func (p *program) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.

	p.quit = make(chan struct{})

	p.wg.Add(1)
	go func() {
		log.Println("Starting...")
		<-p.quit
		log.Println("Quit signal received...")
		p.wg.Done()
	}()

	return nil
}

// 退出方法
func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.

	log.Println("Stopping...")
	close(p.quit)
	p.wg.Wait()
	log.Println("Stopped.")
	return nil
}

func Run(){
	prg := &program{}

	// Call svc.Run to start your program/service.
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}
