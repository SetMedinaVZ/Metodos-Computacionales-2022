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
            
    print("\nToken\tTipo")
    
    for i in range(len(strings)):
        isComentario(strings[i]);
        isVariable(strings[i]);
        isReal(strings[i]);


lexerAritmetico('expresiones')