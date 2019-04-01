// hello project main.go
package main

import (
	"fmt"
	"github.com/CUjamin/hellogo/cuj/hellogo/main/mainpackage/point"
	"github.com/CUjamin/hellogo/cuj/hellogo/main/mainpackage/lockdemo"
	"github.com/CUjamin/hellogo/cuj/hellogo/main/mainpackage/srcdemo/bufiodemo"
	"github.com/CUjamin/hellogo/cuj/hellogo/main/mainpackage/srcdemo/runtimedome"
	"github.com/CUjamin/hellogo/cuj/hellogo/main/mainpackage/srcdemo/osdemo"
)

// func main() {
// 	fmt.Println("Hello World!!!")
// }

func main() {
	runtimedome.Test()
	osdemo.Test()
	bufiodemo.Inputdemo()
	lockdemo.Start()
	var p point.Point = point.Point{2,2}
	var q point.Point = point.Point{1,1}
	
	fmt.Println(point.Discount(p,q))

	var s collection.Stack
	s.Push("test")
	fmt.Println(s.Pop())
	fmt.Println(s.Size())

	fmt.Println("END")
}
