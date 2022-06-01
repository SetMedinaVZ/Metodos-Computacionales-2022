package main

import "fmt"

func sub(A []int, final int, c chan int){
	var r int 
	for i := 0; i < final; i++ {
		r += A[i]
	}
	c <- r
}

func main() {
	A := [6]int{3,1,19,2,4,5}
	c := make(chan int)

	for i:=0; i < len(A); i++ {
		go sub(A[:], i, c)
	}

	var resultado int
	