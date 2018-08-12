FICHEROS
===

## Leer ficheros

Para leer ficheros debemos hacer uso de la libreria **io/ioutil**

    $ import "io/ioutil"
    
    $ texto, error := ioutil.ReadFile("fichero.txt")

## Escribir ficheros

Para escribir texto en el archivo se debe tener en cuenta lo siguiente

1. Leer valores ingresados

        $ nuevoTexto := os.Args[1] + "\n"

2. Abrir el fichero

        $ fichero, err := os.OpenFile("fichero.txt", os.O_APPEND|os.O_WRONLY, 0777)

3. Escribir en el fichero

        $ escribir, err := fichero.WriteString(nuevoTexto)

4. Cerrar el fichero

        $ fichero.Close()