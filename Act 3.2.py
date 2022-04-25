import re

strings = []

#Nombre funcion: limpiaLista
#Descripcion: Retorna una lista sin elementos vacios
#Parametros: array -> Lista de elementos a la cual queremos quitarle elementos vacios
def limpiaLista(array):
    #Generamos un array donde guardaremos los elementos NO vacios
    aux = []
    #Añadimos los elementos al array
    for i in array:
        for j in i:
            if j != '':
                aux.append(j)
                
    return aux

#Nombre funcion: imprimeResultado
#Descripcion: Imprime el token y su tipo
#Parametros: lista -> Lista de elementos a imprimir
#            tipo -> tipo del elemento
def imprimeResultado(lista, tipo):
    #Imprimos los elementos de la lista junto con su tipo
    for i in lista:
            print(i +'\t' +tipo)
  
#Nombre funcion: eliminaElemento
#Descripcion: Borra los elementos de la lista recibida en el array global de strings
#Parametros: lista -> Lista de elementos a borrar en el array global         
def eliminaElemento(lista):
    n = len(lista)
    #Recorremos la lista global
    for i in range(len(strings)):
        #Obtenemos el elemetno del string en la posición i
        elemString = strings[i]
        #Establecemos un contador que recorrera la lista
        j = 0
        #Si la lista no esta vacia y el contador es diferente de la longitud de la lista
        while n > 0 and j != n:
            elemLista = lista[j]
            #Si el elemento que queremos borrar esta en la lista global
            if elemLista in elemString:
                #Elimanamos el elemento que queremos borrar de la lista global
                strings[i] = strings[i].replace(lista[j], '')
                #Eliminamos el elemento ubicado en la posición j
                lista.pop(j)
                #Actualizamos la longitud de la lista
                n = len(lista)
                #Restablecemos el contador
                j = 0
            else:
                #Aumentamos el contador en 1
                j = j + 1
        #Una vez que se haya recorrido toda la linea y se elimina nos salimos de la función
        if len(lista) == 0:
            break

#Nombre funcion: isOperador
#Descripcion: Obtiene la lista de operador(+, -, /, *, ^) de la linea recibida
# e imprime los operadores
#Parametros: string -> string recibido por parte de la funcion principal
def isOperador(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("([+\-*/\^=])", string)
    
    if not x:
        return False
    else:
        #Limpiamos los valores vacios de la tupla obtenida
        lista = limpiaLista(x)
        #Imprimos el operador correspondiente
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

#Nombre funcion: isAsign
#Descripcion: Obtiene la lista de simbolos igual de la linea recibida
#Parametros: string -> string recibido por parte de la funcion principal
def isAsign(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("([=])", string)
    if not x:
        return False
    else:
        #Imprmimos las asignaciones (simbolos igual)
        imprimeResultado(x, 'Asignación') 

#Nombre funcion: isAPar
#Descripcion: Obtiene la lista de parentesis que abren de la linea recibida
#Parametros: string -> string recibido por parte de la funcion principal
def isAPar(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("([(])", string)
    if not x:
        return False
    else:
        #Imprimimos los parentesis que abren
        imprimeResultado(x, 'Abre')

#Nombre funcion: isCPar
#Descripcion: Obtiene la lista de parentesis que cierran de la linea recibida
#Parametros: string -> string recibido por parte de la funcion principal
def isCPar(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("([)])", string)
    if not x:
        return False
    else:
        #Imprimos los parentesis que cierran
        imprimeResultado(x, 'Cierra')

#Nombre funcion: isComentario
#Descripcion: Obtiene la lista de comentarios de la linea recibida
#Parametros: string -> string recibido por parte de la funcion principal
def isComentario(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("(/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/)|(//.*)", string)
    if not x:
        return False
    else:
        #Limpiamos los valores vacios de la lista obtenida de la expresion regular
        lista = limpiaLista(x)
        #Imprimos los comentarios de la linea
        imprimeResultado(lista, 'Comentario')
        #Eliminamos el comentario del string en el que esta
        eliminaElemento(lista)

#Nombre funcion: isVariable
#Descripcion: Obtiene la lista de variables de la linea recibida
#Parametros: string -> string recibido por parte de la funcion principal
def isVariable(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("(?![_0-9])([a-zA-Z_$][a-zA-Z_$0-9]*)", string)
    #Si la tupla esta vacia
    if not x:
        return False
    else:
        #Imprimos las variables de la linea
        imprimeResultado(x, 'Variable') 

#Nombre funcion: isInt
#Descripcion: retorna la lista de numeros enteros 
#Parametros: array -> array recibido por la funcion isReal
def isInt(array):
    #Generamos un array auxiliar donde se guardaran los numeros enteros
    aux = []
    #Recorremos el array de numeros reales
    for i in array:
        #Buscamos el numero entero
        x = re.search("^[+-]?[0-9]*(?![.0-9])+$", i)
        #Si es un numero entero lo agregamos al array aux
        if x:
            aux.append(x.group())
    return aux        

#Nombre funcion: eliminaIntFromReal
#Descripcion: retorna la lista de numeros reales sin los numeros enteros
#Parametros: arrayInt -> array de numeros enteros recibido por la funcion isReal
#            arrayReals -> array de numeros reales recibido por la funcion isReal
def eliminaIntFromReal(arrayInt, arrayReals):
    #Generamos una copia del array de numeros enteros
    aux = arrayInt.copy()
    #Obtenemos la longitud del array de numeros enteros
    n = len(aux)
    #Obtenemos la longitud del array de numeros reales
    n_2 = len(arrayReals) 
    #Si el array de enteros no esta vacio
    if len(arrayInt) > 0:
        #Inicialiazamos contador para recorrer el array de numeros reales
        i = 0
        #Mientras el contador sea menor o igual a la longitud del array de numeros reales
        while i <= n_2:
            #Inicialiazamos contador para recorrer el array de numeros enteros
            j = 0
            #Inicializamos un booleano para poder salirnos del proximo while
            flag = True
            #Mientras la longitud del array de numeros enteros sea mayor o igual a cero
            #y la bandera sea true
            while n >= 0 and flag:
                #Si la longitud del array de numeros enteros es cero retornamos la lista de reales
                if n == 0:
                    return arrayReals
                #Si el contador llego al ultimo elemento nos salimos del while
                elif n == j:
                    flag = False
                else:
                    #Obtenemos el numero real en la posicion i
                    elemReal = arrayReals[i]
                    #Obtenemos el numero entero en la posicion j
                    elemInt = aux[j]
                    #Si el numero entero y el real es el mismo
                    if elemInt == elemReal:
                        #Eliminamos el numero entero de su array
                        aux.pop(j)
                        #Eliminamos el numero entero del array de reales
                        arrayReals.pop(i)
                        #Actualizamos la longitud del array de numeros enteros
                        n = len(aux)
                        #Actualizamos la longitud del array de numeros reales
                        n_2 = len(arrayReals)
                        #Restablecemos el contador para recorrer el array de numeros enteros
                        j = 0
                    else:
                        #Incrementamos j
                        j = j + 1
            #Incrementamos i
            i = i + 1

#Nombre funcion: isReal
#Descripcion: Obtiene la lista de numeros reales y enteros de la linea recibida
#Parametros: string -> string recibido por parte de la funcion principal
def isReal(string):
    #Obtenemos la tupla de los strings que cumplen con la expresion regular
    x = re.findall("([+-]?[0-9]+\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?\.[0-9]+[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.?[eE][+-]?[0-9]{1,3})|([+-]?[0-9]+\.[0-9]+)|([+-]?\.[0-9]+)|([+-]?[0-9]+\.?)", string)
    #Si el array esta vacio retorna false
    if not x:
        return False
    else:
        #Quitamos elementos vacios de la lista
        lista = limpiaLista(x)
        #Obtenemos los numeros enteros de la lista de numeros reales
        enteros = isInt(lista)
        #Si hay numeros enteros
        if len(enteros) > 0:
            #Eliminamos los numeros enteros de la lista de numeros reales
            lista = eliminaIntFromReal(enteros, lista)
            #Imprimimos los numeros enteros
            imprimeResultado(enteros, 'Entero')
            #Eliminamos los numeros enteros de la linea a la que pertenecen
            eliminaElemento(enteros)
            #Verificamos que la lista de numeros reales no este vacia
            if len(lista) > 0:
                #Imprimirmos los numeros reales
                imprimeResultado(lista, 'Real')
                #Eliminamos los numeros reales de la linea que pertenecen
                eliminaElemento(lista)
        #Si no hay numeros enteros
        else:
            #Imprimimos los numeros reales
            imprimeResultado(lista, 'Real') 
            #Eliminamos los numeros reales de la linea que pertenecen
            eliminaElemento(lista)
        
#Nombre funcion: lexerArimetico
#Descripcion: Función principal la cual administra las demas funciones.
#Parametros: nombreArchivo -> es el nombre del archivo el cual buscara en la misma
#                             ubicación donde se encuentra el .py
def lexerAritmetico(nombreArchivo):
    #Abrimos el archivo segun el nombre recibido
    with open(nombreArchivo + ".txt") as f:
        lineas = f.readlines()
        #Guardamos las lineas del txt en un array
        for linea in lineas:
            strings.append(linea.strip('\n'))
        
    #Mandamos linea por linea del txt a las funciones correspondientes       
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