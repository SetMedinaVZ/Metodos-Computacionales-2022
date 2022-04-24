import re

strings = []

def limpiaLista(array):
    aux = []
    for i in array:
        for j in i:
            if j != '':
                aux.append(j)
                
    return aux

def imprimeResultado(lista, tipo):
    for i in lista:
            print(i +'\t' +tipo)
            
def eliminaElemento(lista):
    for i in lista:
        n = 0
        for j in strings:
            if i in j:
                strings[n] = strings[n].replace(i, '')
            n = n + 1

def isOperador(string):
    x = re.findall("([+\-*/\^=])", string)
    
    if not x:
        return False
    else:
        lista = limpiaLista(x)
        for i in lista:
            if i == "+":
                imprimeResultado(i, 'Suma')
            elif i == "-":
                imprimeResultado(i, 'Resta') 
            elif i == "/":
                imprimeResultado(i, 'Division') 
            elif i == "*":
                imprimeResultado(i, 'Multiplicacion') 
            elif i == "^":
                imprimeResultado(i, 'Potencia') 

def isAsign(string):
    x = re.findall("([=])", string)
    if not x:
        return False
    else:
        imprimeResultado(x, 'Asignación') 

def isAPar(string):
    x = re.findall("([(])", string)
    if not x:
        return False
    else:
        imprimeResultado(x, 'Abre')

def isCPar(string):
    x = re.findall("([)])", string)
    if not x:
        return False
    else:
        imprimeResultado(x, 'Cierra')


def isComentario(string):
    x = re.findall("(/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/)|(//.*)", string)
    if not x:
        return False
    else:
        lista = limpiaLista(x)
        imprimeResultado(lista, 'Comentario')
        
        #Eliminamos el comentario del string en el que esta
        eliminaElemento(lista)

def isVariable(string):
    x = re.findall("(?![_0-9])([a-zA-Z_$][a-zA-Z_$0-9]*)", string)
    if not x:
        return False
    else:
        imprimeResultado(x, 'Variable') 


def isInt(array):
    aux = []
    for i in array:
        x = re.search("^[+-]?[0-9]*(?![.0-9])+$", i)
        if x:
            aux.append(x.group())
    return aux        

def isReal(string):
    x = re.findall("([+-]?[0-9]+\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.?[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.[0-9]+)|([+-]?\.[0-9]+)|([+-]?[0-9]+\.?)", string)
    if not x:
        return False
    else:
        lista = limpiaLista(x)
        enteros = isInt(lista)
        if len(enteros) > 0:
            for i in enteros:
                n = 0
                for j in lista:
                    if i in j:
                        lista.pop(n)
                        n = n + 1
            imprimeResultado(enteros, 'Entero')
            imprimeResultado(lista, 'Real')
            eliminaElemento(lista)
        else:
            imprimeResultado(lista, 'Real') 
            eliminaElemento(lista)
        

def lexerAritmetico(nombreArchivo):

    with open(nombreArchivo + ".txt") as f:
        lineas = f.readlines()
        for linea in lineas:
            strings.append(linea.strip('\n'))
        
            
    for i in range(len(strings)):
        print("\nLinea " + str(i+1));
        print("Token\tTipo")
        isComentario(strings[i]);
        isReal(strings[i]);
        isVariable(strings[i]);
        isAsign(strings[i]);
        isOperador(strings[i]);        
        isAPar(strings[i]);
        isCPar(strings[i]);
    
    
        
lexerAritmetico('expresiones')