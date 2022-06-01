package main

import (
	"fmt"
	"math"
	"time"
)

//Funcion que genera el array de numeros primos con un rango determinado
//y devuelve la suma de los primos de manera paralela
func primeNumbersSum(n int) int{
	if n < 2 {
		return 0
	}
	//Array de numeros primos
	var primes []int

	//For que genera los numeros primos dentro del rango
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i % j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	//Sumamos los primos
	sum := 0
	for _, v := range primes {
		sum += v
	}
	return sum
}

//Funcion que genera el array de numeros primos con un rango determinado
//y devuelve la suma de los primos de manera paralela
//n1 y n2 son los limites de la criba
func primeNumbersSum_Parallel(n1, n2 int, c chan int){
	if n1 < 2  || n2 < 2 {
		return
	}
	//Array de numeros primos
	var primes []int

	//For que genera los numeros primos dentro del rango
	for n1 <= n2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(n1))); i++ {
			if n1 % i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, n1)
		}
		n1++
	}

	//Sumamos los primos
	sum := 0
	for _, v := range primes {
		sum += v
	}
	
	//Enviamos el resultado al canal
	c <- sum
}

func main(){
	//Tiempo inicial
	startParallel := time.Now()

	//Creamos un canal
	c := make(chan int)

	//Cantidad de hilos a usar
	const nThreads = 10

	//Por cada hilo, se ejecuta la funcion sieveOfEratosthenes.
	go primeNumbersSum_Parallel(2, 500_000, c)
	go primeNumbersSum_Parallel(500_001, 1_000_000, c)
	go primeNumbersSum_Parallel(1_000_001, 1_500_000, c)
	go primeNumbersSum_Parallel(1_500_001, 2_000_000, c)
	go primeNumbersSum_Parallel(2_000_001, 2_500_000, c)
	go primeNumbersSum_Parallel(2_500_001, 3_000_000, c)
	go primeNumbersSum_Parallel(3_000_001, 3_500_000, c)
	go primeNumbersSum_Parallel(3_500_001, 4_000_000, c)
	go primeNumbersSum_Parallel(4_000_001, 4_500_000, c)
	go primeNumbersSum_Parallel(4_500_001, 5_000_000, c)

	sum := 0
	//Recibimos los resultados de cada hilo y los sumamos
	for i := 0; i < nThreads; i++ {
		sum += <-c
	}

	//Tiempo final
	endParallel := time.Now()

	//Imprimimos el resultado
	fmt.Println("La suma de los primos de manera paralela es:", sum, "en", endParallel.Sub(startParallel))

	//Tiempo inicial
	start := time.Now()

	sum2 := primeNumbersSum(5_000_000)

	//Tiempo final
	end := time.Now()

	//Imprimimos el resultado
	fmt.Println("La suma de los primos de manera secuencial es:", sum2, "en", end.Sub(start))

}

