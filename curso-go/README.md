Hola Mundo
===

## Pasos

Crear el directorio "curso-go"

Crear el fichero "hola-mundo.go"

Ejecutar el fichero creado:
    
    $ go run hola-mundo.go

## Librerias y Sleep

Para importar librerias debemos usar la palabra **import**

Importar la libreria la libreria **time**

Hacer un retardo en el programa de 5 segundos

    $ time.Sleep(time.Second * 5)

Variables y Tipos
===

## Comandos Utiles

Ejecutar un fichero 

    $ go run [file]

Obtener informacion de un paquete o libreria

    $ godoc [paquete|libreria]

Tabular y dar un formato correcto a un fichero.go

    $ go fmt [file]

Compilar y obtener un ejecutable de nuestro programa

    $ go build [file]

## Mas variables

Crear variables con la palabra reservada **var**

En GO es opcional poner **;** o no

Tambien se pueden crear variables de otra manera **:=**

    $ pais := "Peru"

De esta manera GO va inferir que tipo de variable tengo

## Constantes

Para definir una constante usamos la palabra reservada **const**

## Operadores Aritmeticos

Suma (+)
Resta (-)
Multiplicacion (*)
Division (/)
Resto (%)

## Tipos personalizados. Structs

Para crear un tipo de dato personalizado usar la palabra reservada **type**

Funciones
===

## Crear funciones

Para definir una funcion tenemos la palabra reservada **func**

## Retorno de datos

Se puede devolver mas de un dato

