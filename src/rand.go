package main

import "fmt"

func BubbleSort(a []int) {
	len := len(a)
	for i := 0; i < len-1; i++ {
        for j := 0; j < len-1-i; j++ {
            if a[j] > a[j+1] {
                a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := [...]int{5, 3, 8, 100, 6, 9}
	BubbleSort(a[:])
	fmt.Println("a= ", a)

}