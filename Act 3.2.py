import re

def lexerAritmetico(strings):
    print("\nToken\tTipo")
    

def main():
    strings = []
    with open("expresiones.txt") as f:
        lineas = f.readlines()
        for linea in lineas:
            strings.append(linea.strip('\n'))
    print (strings)

    lexerAritmetico(strings)

main()