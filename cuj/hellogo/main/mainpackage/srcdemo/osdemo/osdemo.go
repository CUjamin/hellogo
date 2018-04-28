package osdemo

import (
	"os"
	"fmt"
)
func Test(){
	var x , error = os.Hostname()
	if error!= nil{
		fmt.Printf("Host Name : %s \n",x)
	}

	fmt.Printf("Get PageSize : %d \n",os.Getpagesize())

	fmt.Printf("Environ : %s \n",os.Environ())

	fmt.Printf("GOPATH : %s \n",os.Getenv("GOPATH"))

	fmt.Printf("Getppid : %s \n", os.Getppid())

	x , error = os.Getwd()
	if error!=nil{
		fmt.Printf("Getwd ：%s %s\n",x,error)
	}
}
