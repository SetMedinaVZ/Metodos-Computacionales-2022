import re

def lexerAritmetico(strings):
    print("\nToken\tTipo")
    

def main():
    strings = []
    with open("expresiones.txt") as f:
        for lineas in f:
            strings.extend(lineas.split())
    print (strings)

    lexerAritmetico(strings)

main()