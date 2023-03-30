# Bookmark API con Fiber

Este repositorio contiene el código para una API de marcadores desarrollada con el framework de Go llamado [Fiber](https://github.com/gofiber/fiber). La API permite a los usuarios registrar y gestionar sus marcadores, así como realizar búsquedas en ellos.

## Instalación

Antes de comenzar, asegúrate de tener instalado Go y Fiber. Si aún no los tienes, puedes descargar Go desde el sitio web oficial de [Go](https://golang.org/doc/install) y Fiber desde su [repositorio en Github](https://github.com/gofiber/fiber#installation).

Una vez que tengas Go y Fiber instalados, puedes clonar este repositorio con el siguiente comando:

git clone https://github.com/Jarbram/bookmark-api-fiber.git


Una vez que se haya clonado el repositorio, puedes instalar las dependencias utilizando el siguiente comando:


## Ejecución

Una vez que hayas instalado las dependencias y configurado las variables de entorno, puedes ejecutar la aplicación con el siguiente comando:


La aplicación se ejecutará en el puerto especificado en la variable de entorno `PORT` del archivo `.env`. Por defecto, el puerto es 3000.

## Endpoints

La API proporciona los siguientes endpoints:

- `GET /bookmarks`: Devuelve una lista de todos los marcadores.
- `GET /bookmarks/:id`: Devuelve un marcador específico según su ID.
- `POST /bookmarks`: Crea un nuevo marcador.
- `PUT /bookmarks/:id`: Actualiza un marcador específico según su ID.
- `DELETE /bookmarks/:id`: Elimina un marcador específico según su ID.

Puedes encontrar más detalles sobre los parámetros y las respuestas en el archivo `routes/bookmark.go`.

## Contribuir

Si deseas contribuir a este proyecto, por favor crea un `pull request`. Todos los cambios son bienvenidos.


