package main 
import "fmt"

func test() {
	fmt.Println("test")
}

type typea func() 

func (aa typea)testb () {
	fmt.Println("testb")
}

func main() {
	 a := test
	 a()
	fmt.Printf("a is %T", a)
}
