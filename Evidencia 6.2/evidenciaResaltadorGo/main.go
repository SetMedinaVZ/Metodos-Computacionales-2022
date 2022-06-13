//Alejandro Daniel Dennis Hernández----A00831138
//Sebastián Medina Veraza--------------A00827461
//Mauricio Ian Gómez Flores------------A01023943

package main

import (
	rs "evidenciaResaltadorGo/resaltadorSintaxis"
	"fmt"
	"time"
)

var nombresDeArchivoEntrada = []string{
	"test/test1.txt", "test/test2.txt",
	"test/test3.txt","test/test4.txt",
	"test/test5.txt", "test/test6.txt",
	"test/test7.txt","test/test8.txt",
	"test/test9.txt","test/test10.txt",
} 


var nombresDeArchivoSalidaSecuencial = []string{
	"salida1secuencial.html", "salida2secuencial.html",
	"salida3secuencial.html","salida4secuencial.html",
	"salida5secuencial.html", "salida6secuencial.html",
	"salida7secuencial.html","salida8secuencial.html",
	"salida9secuencial.html","salida10secuencial.html",
}

var nombresDeArchivoSalidaParalelo = []string{
	"salida1paralelo.html", "salida2paralelo.html",
	"salida3paralelo.html","salida4paralelo.html",
	"salida5paralelo.html", "salida6paralelo.html",
	"salida7paralelo.html","salida8paralelo.html",
	"salida9paralelo.html","salida10paralelo.html",
}

func main() {
	//SECUENCIAL
	tiempoInicioS := time.Now()
	//Creamos un for para mandar a llamar a la función de manera concurrente
	for i := 0; i < 10; i++ {
		rs.ResaltadorSecuencial(nombresDeArchivoEntrada[i], nombresDeArchivoSalidaSecuencial[i])
	}
	tiempoFinS := time.Now()
	tiempoTotalS := tiempoFinS.Sub(tiempoInicioS)
	fmt.Println("Tiempo de ejecución concurrente: ", tiempoTotalS)

	//PARALELO
	tiempoInicioP := time.Now()
	c := make(chan string)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[0], nombresDeArchivoSalidaParalelo[0], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[1], nombresDeArchivoSalidaParalelo[1], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[2], nombresDeArchivoSalidaParalelo[2], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[3], nombresDeArchivoSalidaParalelo[3], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[4], nombresDeArchivoSalidaParalelo[4], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[5], nombresDeArchivoSalidaParalelo[5], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[6], nombresDeArchivoSalidaParalelo[6], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[7], nombresDeArchivoSalidaParalelo[7], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[8], nombresDeArchivoSalidaParalelo[8], c)
	go rs.ResaltadorParalelo(nombresDeArchivoEntrada[9], nombresDeArchivoSalidaParalelo[9], c)
	for i := 0; i < 10; i++ {
		<-c
	}

	tiempoFinP := time.Now()
	tiempoTotalP := tiempoFinP.Sub(tiempoInicioP)
	fmt.Println("Tiempo de ejecución paralelo: ", tiempoTotalP)

	//COMPARACIÓN DE TIEMPOS
	fmt.Println("Speedup: ", tiempoTotalS.Seconds()/tiempoTotalP.Seconds())

}