package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hola Mundo")
	var flag1 = true
	var flag2 = false

	var a, b = 1, 2
	fmt.Println("a+b ", a+b)
	fmt.Println("a = ", a)
	a++
	fmt.Println("a++ ", a)
	fmt.Println("f1 == f2", (flag1 == flag2))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Mi nombre es: ")
	name, _ := reader.ReadString('\n')
	fmt.Print(name)
	if strings.Compare(name, "Anubis") == 1 {
		fmt.Println("Claro")
	} else if strings.Compare(name, "Enki") == 1 {
		fmt.Println("Oscuro")
	} else {
		fmt.Println("Claront")
	}
	if rand_num := rand.Intn(10); rand_num == 0 {
		fmt.Println("El numero aleatorio es: ", rand_num)
	}else{
		fmt.Println("No es 0")
	}
}