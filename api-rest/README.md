Servidor web y API RESTful con Golang y MongoDB
===

## Crear Servidor

- Uno de sus puntos fuertes es hacer API REST muy rapida

- Importar el paquete **net/http**

## Router

1. El paquete **gorilla/mux** implementa un enrutador de solicitud y un despachador para hacer coincidir las solicitudes entrantes con su manejador respectivo
    - https://github.com/gorilla/mux

2. Para instalar en el proyecto debemos ejecutar lo siguiente en consola: 
    - go get -u github.com/gorilla/mux

3. Crear un sistema de rutas

        $ router := mux.NewRouter().StrictSlash(true)
        $ router.HandleFunc("/", Index)
        $ router.HandleFunc("/contacto", Contact)