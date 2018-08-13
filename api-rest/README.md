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

## Parametros URL

- Configuracion de la ruta

        $ router.HandleFunc("/pelicula/{id}", MovieShow)

- Recoger parametros por la URL

        $ params := mux.Vars(r)
        $ movie_id := params["id"]
        $ fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)

## Modelos

- El modelo va ser una pequeña clase, dentro de GO llamado *struct*

- Todo lo que este dentro del paquete *main* estara disponible desde cualquier otro fichero

- Importar la libreria *encoding/json* que nos permitira devolver la respuesta en un formato JSON

## Modelos y JSON

- Por buenas practicas los nombre por defecto a la hora de devolver el JSON deben ser en minuscula

- Para evitar modificar el Struct, generamos un alias

        type Movie struct {
            Name     string `json:"name"`
            Year     int    `json:"year"`
            Director string `json:"director"`
        }

## Arquitectura

- La API ya esta optimizada, mejor encapsulada y abstraida

## Metodo POST para crear elementos

1. Definir el la ruta *MovieAdd*

2. Definir la accion *MovieAdd*

3. Definir la variable global *movies*

5. Utilizar POSTMAN para agregar una pelicular a la aplicación

        URL: http://localhost:8080/pelicula
        
        Headers: Content-Type:application/json

        Body (raw)
        {
            "name": "El Señor de los Anillos",
            "year": 2015,
            "director": "J.K. Rowling"
        }

        