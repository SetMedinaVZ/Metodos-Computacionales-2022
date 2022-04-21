import re

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
                   

def isReal(string):
    x = re.findall("([+-]?[0-9]+\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.?[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.[0-9]+)|([+-]?\.[0-9]+)|([+-]?[0-9]+\.?)", string)
    if not x:
        return False
    else:
        lista = limpiaLista(x)
        imprimeResultado(lista, 'Real')  

def lexerAritmetico(strings):
    print("\nToken\tTipo")
    for i in strings:
        isReal(i)


def main():
    strings = []
    with open("expresiones.txt") as f:
        lineas = f.readlines()
        for linea in lineas:
            strings.append(linea.strip('\n'))
    lexerAritmetico(strings)

main()