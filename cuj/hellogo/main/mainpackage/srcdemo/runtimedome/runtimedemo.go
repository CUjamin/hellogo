package runtimedome

import (
	"fmt"
	"runtime"
)
func Test(){
	fmt.Printf("GOOS : %s ; GOARCH : %s ； Version : %s ; GOROOT ：%s \n",runtime.GOOS,runtime.GOARCH,runtime.Version(),runtime.GOROOT())
}

func main() {
	fmt.Println("Hello World!")
}