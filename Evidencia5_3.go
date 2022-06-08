//Alejandro Daniel Dennis Hernández----A00831138
//Sebastián Medina Veraza--------------A00827461
//Mauricio Ian Gómez Flores------------A01023943
//Resaltador de sintaxis de python con html y css

package main

import (
	"fmt"
	"bufio"
	"strings"
)

//Tokens
const (
	//Palabras reservadas
	KW = 501
	//Variables
	ID = 502
	//Numeros
	INT = 503
	FLOAT = 504
	//Comentarios
	COMM = 505
	//Operadores aritmeticos
	OPA = 506
	//Asignacion
	ASI = 507
	//Operadores logicos
	OPL = 508
	//Delimitadores
	DEL = 509
	//Error
	ERR = 510
	//Blanco
	BLA = 511
	//EOF
	EOF = 512
)

//Objeto que contiene el lexema y el token
type tokens struct {
	lexema string
	token int
}

func esReservada(lexema string) bool {
	reservadas := []string{"True", "False", "None", "and", "as", "assert", "break", "class", "continue", "def", "del", "elif", "else", "except", "finally", "for", "from", "global", "if", "import", "in", "is", "lambda", "nonlocal", "not", "or", "pass", "raise", "return", "try", "while", "with", "yield"}
	for _, r := range reservadas {
		if lexema == r {
			return true
		}
	}
	return false
}

func esEntero(lexema string) bool {
	for _, r := range lexema {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func esFloat(lexema string) bool {
	var contador int
	for _, r := range lexema {
		if r == '.' {
			contador++
		}
		if r < '0' || r > '9' {
			return false
		}
	}
	if contador > 1 {
		return false
	}
	return true
} 

func scanner() {
	//Variables para el analisis
	lexema := "" //Palabra que genera el token
	tokens := []tokens{} //Lista de tokens
	leer := true //Bandera para leer caracter
	
	//Leer el archivo
	archivo, err := bufio.NewReader(strings.NewReader(texto)).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	char := string(archivo[0]) //Caracter inicial

	//Analisis lexico
	for leer {
		
}
}