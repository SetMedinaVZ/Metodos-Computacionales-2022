//Alejandro Daniel Dennis Hernández----A00831138
//Sebastián Medina Veraza--------------A00827461
//Mauricio Ian Gómez Flores------------A01023943
//Resaltador de sintaxis de python con html y css

package main

import (
	"bufio"
	"fmt"
	"strings"
)

//Palabras reservadas
var reservadas = []string{
	"True", "False", "None", "and", "as", "assert", 
	"break", "class", "continue", "def", "del", "elif", 
	"else", "except", "finally", "for", "from", "global", 
	"if", "import", "in", "is", "lambda", "nonlocal", "not", "or", 
	"pass", "raise", "return", "try", "while", "with", "yield",
}

//Operadores aritmeticos
var operadoresAritmeticos = []string{
	"+", "-", "*", "/", "%", "**",
}

