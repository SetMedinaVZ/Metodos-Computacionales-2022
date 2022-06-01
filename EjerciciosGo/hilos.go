package main

import (
	"fmt"
	"time"
)

func hola_hilo(id int, c chan int) {
	fmt.Printf("Hola, soy Panchito %d concurrente\n", id)
	c <- id //Enviamos un mensaje al canal
}

func main(){
	/* El canal es compartido por todos los hilos.
	El canal es como una tuberia que nos lleva al resultado
	y los hilos son los que le inyectan contenido.
	*/

	//Creamos un canal 
	c := make(chan int)
	for i := 0; i < 20; i++ {
		//Con la funcion go, creamos un hilo
		go hola_hilo(i, c) 
	}

	for i := 0; i < 20; i++ {
		//Recibimos un mensaje del canal
		resultado := <-c
		fmt.Println("Hilo terminado", resultado)
	}

	time.Sleep(100 * time.Second)
}