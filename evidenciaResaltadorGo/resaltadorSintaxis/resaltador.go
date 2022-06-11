//Alejandro Daniel Dennis Hernández----A00831138
//Sebastián Medina Veraza--------------A00827461
//Mauricio Ian Gómez Flores------------A01023943

//Resaltador de sintaxis de Python con html y css

package resaltadorSintaxis

import (
	"bufio"
	"fmt"
	"os"
)

//Palabras reservadas
var reservadas = []string{
	"True", "False", "None", "and", "as", "assert",
	"break", "class", "continue", "def", "del", "elif",
	"else", "except", "finally", "for", "from", "global",
	"if", "import", "in", "is", "lambda", "nonlocal", "not", "or",
	"pass", "raise", "return", "try", "while", "with", "yield", "print",
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
	"+=", "-=", "*=", "/=", "%=", "**=", "//=", "&=", "|=", "^=", ">>=", "<<=", "!",
}

//Delimitadores
var delimitadores = []string{
	"(", ")", "[", "]", "{", "}", ",", ":", ";", ".",
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

//Funcion que retorna el html del nombre del archivo con el final del mismo
func inyectarHTML(nombre string) string {
	html := "<div class='nombreArchivo'>\n<h3>" + nombre + "</h4>\n</div>\n"
	html += htmlFinal
	return html
}

//Funcion que realiza el resaltado de sintaxis de un archivo de texto recibido
func ResaltadorSecuencial(nombreArchivo, nombreSalida string) {
	//Abrimos el archivo
	archivo, err1 := os.Open(nombreArchivo)

	//Creamos el archivo de salida
	salida, err2 := os.Create(nombreSalida)

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
	line := ""   //Linea actual
	str := "<p>" //linea que se va a agregar al html
	char := ""   //Caracter actual
	lexema := "" //Lexema actual
	count := 0   //Contador de lineas

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
		for i <= len(line) {
			//Verificamos que ya se haya recorrido la linea
			if i == len(line) {
				//Si ya se ha recorrido la linea cerramos el parrafo
				str += "</p>\n"
				//Agregamos la linea al html
				salida.WriteString(str)
				//Reiniciamos la variable
				str = "<p>"
				//Salimos del ciclo
				break

			} else {
				//Obtenemos el caracter actual
				char = string(line[i])
				//Si es comentario
				if char == "#" {
					//Guardamos todo el resto de la linea
					str += "<span class = comentario>" + line[i:] + "</span>"
					//Salimos del ciclo
					i = len(line)

					//Si es un espacio
				} else if char == " " {
					j = 0 //En este caso J cuenta los espacios
					//Recorremos la linea hasta que no haya espacios
					//for string(line[i]) == " " {
					//si el siguiente caracter no es un espacio aumentamos i
					if string(line[i]) == " " {
						str += "&nbsp;"
						i++
						j++
					} else {
						//Si no es un espacio salimos del ciclo
						break
					}
					//}
					//Si el caracter actual es una letra
				} else if estaEn(alfabeto, char) {
					//Guardamos el caracter
					lexema += char
					j = i + 1

					//Recorremos la linea desde i
					for j <= len(line) {
						//Verificamos si j es igual a la longitud de la linea
						if j == len(line) {
							//Si es igual a la longitud de la linea verificamos si el lexema es una palabra reservada o un identificador
							if estaEn(reservadas, lexema) {
								str += "<span class = palabraReservada>" + lexema + "</span>"
								lexema = ""
								i++
								break
							} else {
								str += "<span class = identificador>" + lexema + "</span>"
								lexema = ""
								i++
								break
							}
						}
						//Obtenemos el caracter actual
						char = string(line[j])
						//Si el caracter es una letra o un numero
						if estaEn(alfabeto, char) || estaEn(numeros, char) {
							//Agregamos el caracter
							lexema += char
							//Aumentamos j
							j++

							//Si no es letra o numero
						} else {
							//Verificamos que sea palabra reservada
							if estaEn(reservadas, lexema) {
								//Agregamos el span
								str += "<span class = palabraReservada>" + lexema + "</span>"
								//Reiniciamos el lexema
								lexema = ""

								//Salimos del ciclo
								break

								//Si es delimitador
							} else {
								//Agregamos el span
								str += "<span class = identificador>" + lexema + "</span>"
								//Reiniciamos el lexema
								lexema = ""
								i = j //Dejamos i en donde termino j
								//Salimos del ciclo
								break
							}

						}
						i = j //Dejamos i en donde termino j
					}
					//Si es un numero
				} else if estaEn(numeros, char) {
					//Guardamos el caracter
					lexema += char
					//Igualamos j a i + 1
					j = i + 1

					//Recorremos la linea desde i
					for j <= len(line) {
						//Verificamos si j es igual a la longitud de la linea
						if j == len(line) {
							str += "<span class = numero>" + lexema + "</span>"
							lexema = ""
							i++
							break
						}
						//Obtenemos el caracter actual
						char = string(line[j])
						//Si el caracter es un numero, punto o una E o una e
						if estaEn(numeros, char) || char == "." || char == "E" || char == "e" {
							//Agregamos el caracter
							lexema += char
							//Aumentamos j
							j++

							//Si hay una letra
						} else if estaEn(alfabeto, char) {
							for estaEn(alfabeto, char) || estaEn(numeros, char) {
								//Agregamos el caracter
								lexema += char
								//Aumentamos j
								j++
							}
							//Significa que no es un numero
							str += "<span class = error>" + lexema + "</span>"
							//Reiniciamos el lexema
							lexema = ""
							//Salimos del ciclo
							break

							//Si no es numero o letra
						} else {
							//Significa que es un numero
							str += "<span class = numero>" + lexema + "</span>"
							//Reiniciamos el lexema
							lexema = ""
							//Salimos del ciclo
							break
						}
					}
					i = j //Dejamos i en donde termino j

					//Si es una comilla
				} else if char == "\"" {
					//Guardamos el caracter
					lexema += char
					//Igualamos j a i + 1
					j = i + 1
					//Recorremos la linea desde i
					for j < len(line) {
						//Obtenemos el caracter actual
						char = string(line[j])
						//Si el caracter es una comilla creamos el span
						if char == "\"" {
							lexema += char
							str += "<span class = cadena>" + lexema + "</span>"
							//Reiniciamos el lexema
							lexema = ""
							//Aumentamos j
							j++
							//Salimos del ciclo
							break
						} else {
							//Agregamos el caracter
							lexema += char
							//Aumentamos j
							j++
						}
					}
					i = j //Dejamos i en donde termino j

					//Si es un operador
				} else if estaEn(operadores, char) {
					//Guardamos el caracter y cremos el span
					str += "<span class = operador>" + char + "</span>"
					//Aumentamos i
					i++

					//Si es un delimitador
				} else if estaEn(delimitadores, char) {
					//Guardamos el caracter y cremos el span
					str += "<span class = delimitador>" + char + "</span>"
					//Aumentamos i
					i++

				} else {
					//Si no es nada de lo anterior, es un error
					str += "<span class = error>" + char + "</span>"
					//Aumentamos i
					i++
				}
			}

		} // Fin del for que recorre la linea

	} //Fin del for que recorre el archivo

	//Escribimos el string final en el archivo
	salida.WriteString(inyectarHTML(nombreArchivo))
	salida.Sync()
}

func ResaltadorParalelo(nombreArchivo, nombreSalida string, c chan string) {
	//Abrimos el archivo
	archivo, err1 := os.Open(nombreArchivo)

	//Creamos el archivo de salida
	salida, err2 := os.Create(nombreSalida)

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
	line := ""   //Linea actual
	str := "<p>" //linea que se va a agregar al html
	char := ""   //Caracter actual
	lexema := "" //Lexema actual
	count := 0   //Contador de lineas

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
		for i <= len(line) {
			//Verificamos que ya se haya recorrido la linea
			if i == len(line) {
				//Si ya se ha recorrido la linea cerramos el parrafo
				str += "</p>\n"
				//Agregamos la linea al html
				salida.WriteString(str)
				//Reiniciamos la variable
				str = "<p>"
				//Salimos del ciclo
				break

			} else {
				//Obtenemos el caracter actual
				char = string(line[i])
				//Si es comentario
				if char == "#" {
					//Guardamos todo el resto de la linea
					str += "<span class = comentario>" + line[i:] + "</span>"
					//Salimos del ciclo
					i = len(line)

					//Si es un espacio
				} else if char == " " {
					j = 0 //En este caso J cuenta los espacios
					//Recorremos la linea hasta que no haya espacios
					//for string(line[i]) == " " {
					//si el siguiente caracter no es un espacio aumentamos i
					if string(line[i]) == " " {
						str += "&nbsp;"
						i++
						j++
					} else {
						//Si no es un espacio salimos del ciclo
						break
					}
					//}
					//Si el caracter actual es una letra
				} else if estaEn(alfabeto, char) {
					//Guardamos el caracter
					lexema += char
					j = i + 1

					//Recorremos la linea desde i
					for j <= len(line) {
						//Verificamos si j es igual a la longitud de la linea
						if j == len(line) {
							//Si es igual a la longitud de la linea verificamos si el lexema es una palabra reservada o un identificador
							if estaEn(reservadas, lexema) {
								str += "<span class = palabraReservada>" + lexema + "</span>"
								lexema = ""
								i++
								break
							} else {
								str += "<span class = identificador>" + lexema + "</span>"
								lexema = ""
								i++
								break
							}
						}
						//Obtenemos el caracter actual
						char = string(line[j])
						//Si el caracter es una letra o un numero
						if estaEn(alfabeto, char) || estaEn(numeros, char) {
							//Agregamos el caracter
							lexema += char
							//Aumentamos j
							j++

							//Si no es letra o numero
						} else {
							//Verificamos que sea palabra reservada
							if estaEn(reservadas, lexema) {
								//Agregamos el span
								str += "<span class = palabraReservada>" + lexema + "</span>"
								//Reiniciamos el lexema
								lexema = ""

								//Salimos del ciclo
								break

								//Si es delimitador
							} else {
								//Agregamos el span
								str += "<span class = identificador>" + lexema + "</span>"
								//Reiniciamos el lexema
								lexema = ""
								i = j //Dejamos i en donde termino j
								//Salimos del ciclo
								break
							}

						}
						i = j //Dejamos i en donde termino j
					}
					//Si es un numero
				} else if estaEn(numeros, char) {
					//Guardamos el caracter
					lexema += char
					//Igualamos j a i + 1
					j = i + 1

					//Recorremos la linea desde i
					for j <= len(line) {
						//Verificamos si j es igual a la longitud de la linea
						if j == len(line) {
							str += "<span class = numero>" + lexema + "</span>"
							lexema = ""
							i++
							break
						}
						//Obtenemos el caracter actual
						char = string(line[j])
						//Si el caracter es un numero, punto o una E o una e
						if estaEn(numeros, char) || char == "." || char == "E" || char == "e" {
							//Agregamos el caracter
							lexema += char
							//Aumentamos j
							j++

							//Si hay una letra
						} else if estaEn(alfabeto, char) {
							for estaEn(alfabeto, char) || estaEn(numeros, char) {
								//Agregamos el caracter
								lexema += char
								//Aumentamos j
								j++
							}
							//Significa que no es un numero
							str += "<span class = error>" + lexema + "</span>"
							//Reiniciamos el lexema
							lexema = ""
							//Salimos del ciclo
							break

							//Si no es numero o letra
						} else {
							//Significa que es un numero
							str += "<span class = numero>" + lexema + "</span>"
							//Reiniciamos el lexema
							lexema = ""
							//Salimos del ciclo
							break
						}
					}
					i = j //Dejamos i en donde termino j

					//Si es una comilla
				} else if char == "\"" {
					//Guardamos el caracter
					lexema += char
					//Igualamos j a i + 1
					j = i + 1
					//Recorremos la linea desde i
					for j < len(line) {
						//Obtenemos el caracter actual
						char = string(line[j])
						//Si el caracter es una comilla creamos el span
						if char == "\"" || char == "'" {
							lexema += char
							str += "<span class = cadena>" + lexema + "</span>"
							//Reiniciamos el lexema
							lexema = ""
							//Aumentamos j
							j++
							//Salimos del ciclo
							break
						} else {
							//Agregamos el caracter
							lexema += char
							//Aumentamos j
							j++
						}
					}
					i = j //Dejamos i en donde termino j

					//Si es un operador
				} else if estaEn(operadores, char) {
					//Guardamos el caracter y cremos el span
					str += "<span class = operador>" + char + "</span>"
					//Aumentamos i
					i++

					//Si es un delimitador
				} else if estaEn(delimitadores, char) {
					//Guardamos el caracter y cremos el span
					str += "<span class = delimitador>" + char + "</span>"
					//Aumentamos i
					i++

				} else {
					//Si no es nada de lo anterior, es un error
					str += "<span class = error>" + char + "</span>"
					//Aumentamos i
					i++
				}
			}

		} // Fin del for que recorre la linea

	} //Fin del for que recorre el archivo

	//Escribimos el string final en el archivo
	salida.WriteString(inyectarHTML(nombreArchivo))
	salida.Sync()
	c <- "El archivo ha acabado!"
}
