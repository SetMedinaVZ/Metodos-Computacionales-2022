package main

import (
	"fmt"
)

func parallel_binarySearch(arreglo []int, low, high, value int, c chan int){
	for i := low; i < high; i++ {
		if arreglo[i] == value {
			c <- i
		}
	}
	c <- -1
}

func main() {
	c := make(chan int)
	arreglo := [8]int{7,3,18,8,10,9,20,31}
	
	//Creamos 4 hilos paar buscar el valor 18 en el arreglo
	go parallel_binarySearch(arreglo[:], 0, 2, 18, c)
	go parallel_binarySearch(arreglo[:], 2, 4, 18, c)
	go parallel_binarySearch(arreglo[:], 4, 6, 18, c)
	go parallel_binarySearch(arreglo[:], 6, 8, 18, c)

	for i := 0;  i < 4; i++{
		index := <-c
		if index != -1 {
			fmt.Println("El elemento esta en la posicion", index)
		}
	}

}