package main

import "fmt"

func main(){
	const lightspeed = 299792 //km/s
	var distance = 56000000 //km
	fmt.Println(distance / lightspeed, "seconds")
	distance = 401000000
	fmt.Println(distance / lightspeed, "seconds")
}