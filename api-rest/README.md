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

## Instalar MongoDB

1. Descargar MongoDB de la pagina web
            
        https://www.mongodb.com/download-center?jmp=nav#community

2. Importar la llave publica de MongoDB
    
        sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4

3. Actualizar la base de datos del administrador de paquetes

        sudo apt-get update

4. Instalar los paquetes de MongoDB

        sudo apt-get install -y mongodb-org

5. Habilitar el servicio de MongoDB

        sudo systemctl enable mongod
        
        sudo service mongod stop
        sudo service mongod start

6. Para conectarse a MongoDB 

        mongo

7. Verificar el funcionamiento

        > use mydb;
        > db.test.save( { tecadmin: 100 } )
        > db.test.find()
            { "_id" : ObjectId("52b0dc8285f8a8071cbb5daf"), "tecadmin" : 100 }

## Mongo GUI (Robo3t)

Para instalar Robo3t en Ubuntu, realizar las siguientes acciones:

1. Descargar Robo3t de la siguiente página

        https://robomongo.org/

2. Ejecutar los siguientes comandos

        $ tar -xvzf robo3t-1.2.1-linux-x86_64-3e50a65.tar.gz
        $ sudo mkdir /usr/local/bin/robo3t
        $ mv robo3t-1.2.1-linux-x86_64-3e50a65/* /usr/local/bin/robo3t
        $ cd /usr/local/bin/robo3t/bin
        $ chmod +x robo3t ./robo3t

3. Verificar que el demonio **mongod** este levantado

4. Iniciar la aplicacion

        $ ./robo3t

5. Crear la conexión local a mongodb

## Conexión a MongoDB desde GO

Para hacer la conexión utilizaremos la librería **mgo** (pronunciado como *mango*)

1. Ingresar a la siguiente URL

        https://labix.org/mgo

2. Ir al proyecto en GO y ejecutar el siguiente comando para instalar la libreria

        $ go get gopkg.in/mgo.v2

3. Además instalar la librería que nos permitira trabajar con los JSON y documentos binarios (**bson**) 

        $ go get gopkg.in/mgo.v2/bson

4. Crear la conexión a MongoDB en el proyecto (*actions.go*)

        $ func getSession() *mgo.Session {
        $       session, err := mgo.Dial("mongodb://localhost")
        $       if err != nil {
        $               panic(err)        
        $       }
        $       return session
        $ }

## Guardar en la base de datos

1. Crear la variable global **collection**

        $ var collection = getSession().DB("curso_go").C("movies")

2. Modificar la accion **MovieAdd** para que almacene en la base de datos **curso_go**

3. Hacer un POST usando *Postman*

- La base de datos no necesita ser creada, al hacer el POST la crea automaticamente y se crea la coleccion **movies**

## Accion para el listado

Modificar la accion *MovieList* para que liste todos los registros de la base de datos creada en MongoDB

## Un solo elemento

Modificar la accion *MovieShow* para que liste un solo elemento en la base de datos creada en MongoDB

## Responses

Crear un par de metodos *responseMovie* & *responseMovies* para evitar codear la mismas respuestas mas de una vez

## Actualizar Documento

1. Crear la nueva ruta en el fichero de rutas

        $ Route{"MovieUpdate", "PUT", "/pelicula/{id}", MovieUpdate}

2. Codear la accion *MovieUpdate* en el fichero de acciones

3. Por POSTMAN utilizar la siguiente configuracion

- **Headers**: Content-Type:application/json
- **Metodo**: PUT
- **Ruta**: http://localhost:8080/pelicula/5b7ac150ab466df8d7ceaf7d
- **Body**: {
	"name": "Harry Potter - La camara Secreta",
	"year": 2012,
	"director": "J.K. Rowling"
}

## Eliminar Documentos

1. Crear la nueva ruta en el fichero de rutas

        $ Route{"MovieRemove", "DELETE", "/pelicula/{id}", MovieRemove},

2. Codear la accion *MovieRemove* en el fichero de acciones

3. Por POSTMAN utilizar la siguiente configuracion

- **Headers**: Content-Type:application/json
- **Metodo**: DELETE
- **Ruta**: http://localhost:8080/pelicula/5b7ac150ab466df8d7ceaf7d