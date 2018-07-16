package main 
import ("fmt")

var a = test() 
func test() int {
	fmt.Println("test")
	return 10
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
