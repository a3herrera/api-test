# api-test

### Ejecucion

Actualmente se dispone de 3 formas de poder estar iniciando el servidor http. Este servidor http se ejecuta en el puerto 8080:

#### Linux execution
Por medio del comando make que genera la compilacion y ejecucion del proyecto.

Si no se tiene instalado el comando make se puede estar ejecutando
``apt-get -y install make ``

Y dentro del directorio del proyecto se ejecuta el comando ``make`` que compila y ejecuta el proyecto

#### Docker-compose
Por medio del comando ``docker-compose up`` que se encarga de estar compilando y ejecutando el proyecto

#### Assets de github
Dentro del repositorio se tiene el tag main que contiene diferentes assets dependiendo del sistema operativo que se posea
se descarga cualesquiera de los ejecutables. Y ya con el ejecutable disponible se dispone a realizarse:

Windows
```shell
  ./searcher.exe server
```

Linux / MacOS

```shell
  ./searcher server
```

El comando extra que se le brinda al ejecutable permite estar levantando el servidor http.


#### Compilacion

Si se desea estar compilando directamente el proyecto para su correspondiente ejecucion se puede estar ejecutando el 
siguiente comando:

``
    go build -ldflags "-s -w -extldflags '-static'" -o searcher
``

Teniendo en consideracion una instalacion previa de Go como requerimiento version 1.16
Luego de la compilacion se puede estar haciendo la ejecucion segun como se encuentra explicado dentro del apartado de assets de github

### API

Se pone a disposicion 2 servicios http, siendo estos:


| METODO | path | descripcion |
| :----: | :---- | : ---- |
| GET | /health/ping | permite conocer si el servicio esta respondiendo |
| GET | /search     | realiza una busqueda por el termine en las diferentes integraciones con otras API, permitiendo centralizar la informacion de las mismas |


#### GET _/health/ping_

Permite conocer si el servidor http se encuentra escuchando solicitudes http. 

Actualmente no realizar accion alguna dentro del sistema mas que estar retornando como respuesta 
```json
"pong"
```


#### GET _/search_

Realiza una busqueda por el termine en las diferentes integraciones con otras API, permitiendo centralizar la informacion de las mismas |

Recibe como query param siendo esta la key *`value`* el cual sel valor que contenga ese el termino de busqueda que se utiliza en cada una
de las diferentes apis integradas para realizar busquedas.

Al no estarse enviando dicho valor, no se indica un valor http de error, mas que solo se retorna un resultado vacio sin
estar realizando una busqueda en las diferentes apis

Ejemplo de utilizacion:

- Sin terminos de busqueda
```
GET /search
Response: 
{
"results": []
}
```

- Incluyendo un termino de busqueda
```
GET /search?value=b
Response: 
{
"results": [
        {
            "uri": "https://itunes.apple.com/search?term=jim",
            "results": [...]
        },
        {
            "uri": "http://api.tvmaze.com/shows?q=jim",
            "results": [...]
        },
        {
            "uri": "http://www.crcind.com/csp/samples/SOAP.Demo.cls?WSDL",
            "results": [...]
        }
    ]
}
```

Actualmente se tienen 3 integraciones, siendo estas:
- itunes
- tvmaze
- crcind

Cuando se realiza una busqueda por un termino y en alguna de las integraciones no se tienen resultados
la misma no es retornada en los resutados de la busqueda

