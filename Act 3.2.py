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
    n = len(lista)
    for i in range(len(strings)):
        elemString = strings[i]
        j = 0
        #Si la lista no esta vacia
        while n > 0 and j != n:
            elemLista = lista[j]
            if elemLista in elemString:
                strings[i] = strings[i].replace(lista[j], '')
                # lista[j] = lista[j].replace(lista[j],'')
                lista.pop(j)
                n = len(lista)
                j = 0
            else:
                j = j + 1
        if len(lista) == 0:
            break

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

def eliminaIntFromReal(arrayInt, arrayReals):
    aux = arrayInt.copy()
    n = len(aux)
    n_2 = len(arrayReals) 
    if len(arrayInt) > 0:
        i = 0
        while i <= n_2:
            j = 0
            flag = True
            while n >= 0 and flag:
                if n == 0:
                    return arrayReals
                elif n == j:
                    flag = False
                else:
                    elemReal = arrayReals[i]
                    elemInt = aux[j]
                    if elemInt == elemReal:
                        #arrayReals[i] = arrayReals[i].replace(arrayInt[j] ,'')
                        aux.pop(j)
                        arrayReals.pop(i)
                        n = len(aux)
                        n_2 = len(arrayReals)
                        j = 0
                    else:
                        j = j + 1
            i = i + 1

def isReal(string):
    x = re.findall("([+-]?[0-9]+\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.?[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.[0-9]+)|([+-]?\.[0-9]+)|([+-]?[0-9]+\.?)", string)
    if not x:
        return False
    else:
        lista = limpiaLista(x)
        enteros = isInt(lista)
        if len(enteros) > 0:
            lista = eliminaIntFromReal(enteros, lista)
            imprimeResultado(enteros, 'Entero')
            eliminaElemento(enteros)
            if len(lista) > 0:
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