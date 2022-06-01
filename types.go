package main

import(
	"fmt"
)

//Declaración de funciones
//func identificador(parametros) tipo_de_retorno {}

//Recibe 2 arreglos de tamaño 3
func dot_product(v1,v2 [3]float64) (float64, bool) {
	if len(v1) != len(v2) {
		return 0.0, false
	}

	var result float64 = 0.0
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result, true
}

//Recibe 2 slices de x tamaño
func dot_product2(v1,v2 []float64) (float64,bool) {
	//Verificar que los slices tengan el mismo tamaño
	if len(v1) != len(v2) {
		return 0.0, false
	}

	var result float64 = 0.0
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result, true
}
/*
//Metodos
func (c celsius) kelvin() float64 {
	return c + 273.15
}
*/

func main(){
	var valor = 3.141592654
	valor2 := 3.141592654
	var valor3 float64 = 3.141592654

	fmt.Println(valor,valor2,valor3)

	var cadena = "Hola Mundo"
	cadena2 := "Hola Mundo"
	var cadena3 string = "Hola Mundo"

	fmt.Println(cadena,cadena2,cadena3)

	//***************************************

	//Declarar un arreglo vacio con tamaño
	var arreglo [5]int

	//Ciclos for
	//For normal estilo python
	for i := range arreglo {
		arreglo[i] = i + 1
		fmt.Printf("El %d-esimo valor del arreglo es: %d\n", i, arreglo[i])
	}

	//For que extrae el valor en la posicion i
	for value,i := range arreglo {
		fmt.Printf("\nSoy el elemento %d y el valor es: %d\n", i, value)
	}

	//Otras maneras de declarar un arreglo

	//Arreglo con valores iniciales y tamaño
	fmt.Println("\nArreglo con valores iniciales y tamaño")
	arreglo2 := [5]int{10,20,30,40,50}
	fmt.Println(arreglo2)

	//Arreglo con valores iniciales y sin declarar el tamaño
	fmt.Println("\nArreglo con valores iniciales y sin declarar el tamaño")
	arreglo3 := [...]int{10,20,30,40,50}
	fmt.Println(arreglo3)

	v1 := [3]float64{1.0, 2.0, 3.0}
	v2 := [3]float64{4.0, 5.0, 6.0}

	//Uso de la funcion dot_product
	result, error := dot_product(v1, v2)
	fmt.Println(v1, " dot ", v2, " = ", result, " (", error, ")")

	//fmt.Println(arreglo)
	s1 := arreglo[1:3]
	/*s2 := arreglo[1:]
	s3 := arreglo[:3]
	s4 := arreglo[:] */

	//Declaración de una variable de tipo slice
	slice_directo := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Slice directo:", slice_directo)

	//Uso de funcion dot_product2 para slices
	vector1 := []float64{1.0, 2.0, 3.0}
	vector2 := []float64{1,2}
	result, error = dot_product2(vector1, vector2)
	fmt.Println(vector1, " dot ", vector2, " = ", result, " (", error, ")")

	fmt.Println(arreglo,s1)
	s1 = append(s1, 8)
	
	/*
	var grados_c celsius
	grados_c = 23
	fmt.Println("Hoy hay ", grados_c, "°C")
	var kelvin1 kelvin = grados_c.kelvin()
	*/
}

