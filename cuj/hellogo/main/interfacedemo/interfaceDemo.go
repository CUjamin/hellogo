package interfacedemo

import(
	"fmt"
)

type Birds interface{
	Twitter() string
	Fly(hight int ) bool
}

type Chicken interface{
	Birds
	Walk()
}

type Sparrow struct{
	name string 
}

func(s *Sparrow) Fly(hight int ) bool{
	fmt.Println("fly:",hight)
	return true
}
func (s *Sparrow) Twitter() string {
    return fmt.Sprintf("%s,jojojo", s.name)
}


func BirdAnimation(bird Birds, high int) {
    fmt.Printf("BirdAnimation of %T\n", bird)
    bird.Twitter()
    bird.Fly(high)
}
