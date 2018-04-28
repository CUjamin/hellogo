package bufiodemo

import (
	"bufio"
	"os"
	"fmt"
)

func Inputdemo() {
	inputReader := bufio.NewReader(os.Stdin)
	inputString , error := inputReader.ReadString('\n')
	if error !=nil {
		fmt.Println(" [ ERROR ] ")
	}else{
		fmt.Println(" [ "+inputString[1:len(inputString)-1]+ " ] ")
	}
}