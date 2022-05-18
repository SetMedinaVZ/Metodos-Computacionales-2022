package main

import "fmt"

func main(){
	fmt.Print("My weight on the surface of Mars is: ")
	fmt.Print(149.0 * 0.3738)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Println(" years old.")
	fmt.Println((40 + (3-31) / 4) % 9)

	const lightspeed = 299792 
	var distance = 50 //Declaración larga de variable
	distance2 := 10 //Declaración corta de variable
	var (  //Declaración de variable en bloque
		k = 30
		m = 40
	)
	var i , j = 10 , 20 //Declaración multiple de variables
	fmt.Println(distance)
	fmt.Println(distance2)
	fmt.Println(k)
	fmt.Println(m)
	fmt.Println(i)
	fmt.Println(j)

}
