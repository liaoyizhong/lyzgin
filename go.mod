module lyzgin

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/gin-gonic/gin v1.4.0
	github.com/judwhite/go-svc v1.1.2
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	golang.org/x/crypto v0.0.0
	golang.org/x/net v0.0.0
	golang.org/x/sys v0.0.0
)

replace golang.org/x/crypto v0.0.0 => github.com/golang/crypto v0.0.0-20190923035154-9ee001bba392

replace golang.org/x/sys v0.0.0 => github.com/golang/sys v0.0.0-20190924154521-2837fb4f24fe

replace golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190923162816-aa69164e4478
