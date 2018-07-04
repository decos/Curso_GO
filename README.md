INTRODUCCION
===

## DEFINICION

Es un lenguaje de programacion de codigo abierto.

Desarrollado por Google 2009.

Tipado Estatico, indicar a cada variable que tipo de dato sera.

Lenguaje compilado: Debe compilarse para poder ejecutarse.

## VENTAJAS

Es sencillo de aprender, bastante parecido a C.

Soporta miles de conexiones simultaneas.

Velocidad y Rendimiento

## PARA QUE NOS SIRVE GO?

Crear scripts para el sistema

Backend / APIs RESTful

Sockets, aplicaciones simultaneas

PRIMEROS PASOS
===

## INSTALAR GO EN LINUX
		
Descargar GO de la pagina oficial para Linux

	https://golang.org/doc/install?download=go1.10.3.linux-amd64.tar.gz

Ejecutar los siguientes comandos
		
	$ tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
		
	$ sudo gedit $HOME/.profile

Añadir al final del documento lo siguiente

	$ export PATH=$PATH:/usr/local/go/bin

Guardar los cambios, cerrar y ejecutar para refrescar el perfil que se esta ejecutando

	$ source ~/.profile
	
Crear el siguiente directorio dentro de $HOME
		
	Documentos/GO/prueba-00

Crear un fichero llamado hello.go

Ingresar el codigo para imprimir "hello, world"

Compilar el fichero, ejecutando:
		
	$ go build

Ejecutar el archivo compilado
		
	$ ./hello -> hello, world
		
Que version de GO tengo instalada?
	
	$ go version

## DOCUMENTACION Y HERRAMIENTAS

Documentacion de GO
	
	https://golang.org/doc/cmd

	
		
	
