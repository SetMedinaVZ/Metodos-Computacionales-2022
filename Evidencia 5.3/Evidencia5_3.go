//Alejandro Daniel Dennis Hernández----A00831138
//Sebastián Medina Veraza--------------A00827461
//Mauricio Ian Gómez Flores------------A01023943
//Resaltador de sintaxis de python con html y css

package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

//Palabras reservadas
var reservadas = []string{
	"True", "False", "None", "and", "as", "assert", 
	"break", "class", "continue", "def", "del", "elif", 
	"else", "except", "finally", "for", "from", "global", 
	"if", "import", "in", "is", "lambda", "nonlocal", "not", "or", 
	"pass", "raise", "return", "try", "while", "with", "yield",
}

//Alfabeto
var alfabeto = []string{
	"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n", "o", "p",
	"q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "_",
}

//Numeros
var numeros = []string{
	"1", "2", "3", "4", "5",
	"6", "7", "8", "9", "0",
}

//Operadores
var operadores = []string{
	"+", "-", "*", "/", "%", "**",
	"!=", "==", "<", ">", "<=", ">=",
	"=", "|", "&", "^", "~", "<<", ">>",
	"//", ">>=", "<<=", "**=", "//=", "&=", "|=", "^=", ">>=", "<<=",
	"+=", "-=", "*=", "/=", "%=", "**=", "//=", "&=", "|=", "^=", ">>=", "<<=",
}

const htmlInitial = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Resaltador de sintaxis</title>
    <!-- Font Awesome -->
    <link
    href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.1/css/all.min.css"
    rel="stylesheet"
    />
    <!-- Google Fonts -->
    <link
    href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
    rel="stylesheet"
    />
    <!-- MDB -->
    <link
    href="https://cdnjs.cloudflare.com/ajax/libs/mdb-ui-kit/3.3.0/mdb.min.css"
    rel="stylesheet"
    />
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <!-- As a heading -->
    <nav class="navbar navbar-dark bg-primary sticky-top">
        <div class="container-fluid">
        <span class="navbar-brand mb-0 h1">Resaltador léxico</span>
        </div>
    </nav>
    <div class="parent">
        <div class="d-flex justify-content-center bg-white">
            <img
                src="./img/Python.jpg"
                class="img-fluid"
                alt=""
                width="100px"
            />
            <img
                src="./img/Go.jpg"
                class="img-fluid"
                alt=""
                width="100px"
            />
        </div>
        <div class="div5">

`

const htmlFinal = `
		</div>
	</div>
</body>
</html>
`

//Funcion que verifica si el string dado esta en un arreglo
func estaEn(arreglo []string, string string) bool {
	for _, elemento := range arreglo {
		if elemento == string {
			return true
		}
	}
	return false
}

func main(){
	//Abrimos el archivo
	archivo, err1 := os.Open("test.txt")

	//Creamos el archivo de salida
	salida, err2 := os.Create("salida.html")

	//Cerramos los archivos
	defer archivo.Close()
	defer salida.Close()

	//Verificamos si hay error en los archivos
	if err1 != nil || err2 != nil {
		fmt.Println("Error al abrir el archivo")
		return
	}

	//Escribimos el html inicial
	salida.WriteString(htmlInitial)

	//Variables para leer el archivo
	line := "" //Linea actual
	str := "<p>" //linea que se va a agregar al html
	char := "" //Caracter actual
	count := 0 //Contador de lineas
	

	//Creamos un buffer para leer el archivo
	scanner := bufio.NewScanner(archivo)

	//Recorremos el archivo linea por linea
	for scanner.Scan() {
		count++

		//Obtenemos la linea actual
		line = scanner.Text()

		//Establecemos i & j
		i := 0 //para recorrer la linea inicio
		j := 0 //auxiliar para recorrer la linea stop

		//Recorremos la linea actual
		for i < len(line) {
			//Obtenemos el caracter actual
			char = string(line[i])

			//Verificamos que ya se haya recorrido la linea
			if i == len(line)-1 {
				//Si ya se ha recorrido la linea cerramos el parrafo 
				str += "</p>"
				//Agregamos la linea al html
				salida.WriteString(str)
				//Reiniciamos la variable
				str = "<p>"
				//Salimos del ciclo
				break

				//Si es comentario
			} else if char == "#" {
				//Guardamos todo el resto de la linea 
				str += "<span class = comentario>" + line[i:] + "</span>" 
				//Salimos del ciclo
				break
				
				//Si es un espacio
			}	else if char == " " {
				j = 0 //En este caso J cuenta los espacios
				//Recorremos la linea hasta que no haya espacios
				for char == " " {
					//si el siguiente caracter no es un espacio aumentamos i
					if string(line[i+1]) != " " {
						str += char 
						i++
					} else{
						//si el siguiente caracter es un espacio aumentamos j e i
						i++
						j++
						//Si j llega a 3 agregamos tab
						if j == 3 {
							//Agregamos el tab
							str += "&nbsp;&nbsp;&nbsp;"
							j = 0 //Reiniciamos j
						}
					}
				}
				//Si el caracter actual es una letra
			} else if estaEn(alfabeto, char) {
				//Guardamos el caracter
				str += char
				j = i + 1

				//Recorremos la linea desde i 
				for j < len(line) {
					//Obtenemos el caracter actual
					char = string(line[j])
					//Si el caracter es una letra o un numero
					if estaEn(alfabeto, char) || estaEn(numeros, char) {
						//Agregamos el caracter
						str += char
						//Aumentamos j
						j++

					  //Si no es letra o numero
					} else {
						//Verificamos que sea palabra reservada
						if estaEn(reservadas, line[i:j]) {
							//Agregamos el span
							str += "<span class = palabraReservada>" + line[i:j] + "</span>"
							//Salimos del ciclo
							break

							//Si no es palabra reservada quiere decir que es un identificador
						} else {
							//Agregamos el span
							str += "<span class = identificador>" + line[i:j] + "</span>"
							//Salimos del ciclo
							break
						}

					}
				}
				i = j //Dejamos i en donde termino j

				//Si es un numero
			} else if estaEn(numeros, char) {
				//Guardamos el caracter
				str += char
				//Igualamos j a i + 1
				j = i + 1

				//Recorremos la linea desde i
				for j < len(line) {
					//Obtenemos el caracter actual
					char = string(line[j])
					//Si el caracter es un numero, punto o una E o una e
					if estaEn(numeros, char) || char == "." || char == "E" || char == "e" {
						//Agregamos el caracter
						str += char
						//Aumentamos j
						j++
					}
				}

			}
		}
	}
}