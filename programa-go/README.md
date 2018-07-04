Estructuras de Control
===

## Condicional if

Si algo se cumple, haces una opcion, si algo no se cumple, haces otra opcion

    $ if edad >= 18 || edad == 17

## Argumentos de la consola

Para recibir valores desde la consola debemos usar la libreria **os**

    $ os.Args

Para convertir un string a entero debemos usar la libreria **strconv**

    $ edad, _ := strconv.Atoi(os.Args[2])