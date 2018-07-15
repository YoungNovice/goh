package main 

import (
	"fmt"
	"math/rand"
	"time"
) 

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		r :=rand.Intn(100) + 1
		fmt.Println(r)
		if	r == 99 {
			break;
		}
	}
	
}