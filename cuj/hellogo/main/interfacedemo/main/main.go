package main
import(
    ifd "github.com/CUjamin/hellogo/cuj/hellogo/main/interfacedemo"
)


func main() {
    var bird ifd.Birds
    sparrow := &ifd.Sparrow{}
    bird = sparrow
    ifd.BirdAnimation(bird, 1000)
    // 或者将sparrow直接作为参数
    ifd.BirdAnimation(sparrow, 1000)
}
