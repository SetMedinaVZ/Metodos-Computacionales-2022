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

def isOperador(string):
    x = re.findall("([+\-*/^=])", string)
    
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

def isPot(string):
    x = re.findall("([/^])", string)
    if not x:
        return False
    else:
        imprimeResultado(x, 'Potencia')

def isAPar(string):
    x = re.findall("([()])", string)
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
        for i in lista:
            n = 0
            for j in strings:
                if i in j:
                    strings[n] = strings[n].replace(i, '')
                n = n + 1
       
def isVariable(string):
    x = re.findall("(?<![_0-9])[a-zA-Z]", string)
    if not x:
        return False
    else:
        imprimeResultado(x, 'Variable') 

def isReal(string):
    x = re.findall("([+-]?[0-9]+\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.?[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.[0-9]+)|([+-]?\.[0-9]+)|([+-]?[0-9]+\.?)", string)
    if not x:
        return False
    else:
        lista = limpiaLista(x)
        imprimeResultado(lista, 'Real')  

def lexerAritmetico(nombreArchivo):

    with open(nombreArchivo + ".txt") as f:
        lineas = f.readlines()
        for linea in lineas:
            strings.append(linea.strip('\n'))
        
            
    for i in range(len(strings)):
        print("\nLinea " + str(i+1));
        print("Token\tTipo")
        isComentario(strings[i]);
        isVariable(strings[i]);
        isAsign(strings[i]);
        #isInt
        isReal(strings[i]);
        isOperador(strings[i]);        
        isPot(strings[i]);
        isAPar(strings[i]);
        isCPar(strings[i]);
        
    


lexerAritmetico('expresiones')
