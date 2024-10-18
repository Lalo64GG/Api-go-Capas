package server

import (
	"firstAPI-go/src/config"
	"firstAPI-go/src/middlewares"
	"firstAPI-go/src/routes"
	"firstAPI-go/src/services"
	"log"

	"github.com/gin-gonic/gin"
)

//* Struct de Server
type Server struct {
	engine   *gin.Engine //* engine es una instancia de *gin.Engine, que es el motor de Gin para manejar las rutas y el middleware
	host 	 string      //* host almacena la dirección IP o el dominio donde se alojará el servidor
	port 	 string      //* port almacena el puerto en el que correrá el servidor
	httpAddr string      //* httpAddr es la combinación de host y port (host:port), se usa como la dirección completa del servidor
}

//* NewServer es un constructor que inicializa una nueva instancia de Server
func NewServer(host, port string) Server {

	//* Inicializamos el struct Server y asignamos valores a sus campos
	srv := Server{
        engine: gin.Default(),  //* Se crea una instancia del engine de Gin con middlewares por defecto, como logging y recuperación de pánicos
        host:   host,           //* Se asigna el host recibido por parámetro
        port:   port,           //* Se asigna el puerto recibido por parámetro
		httpAddr: host + ":" + port, //* httpAddr es una concatenación del host y puerto para formar la dirección completa
    }

	db := config.InitDB() //* Inicializamos la conexión a la base de datos

	//* Aquí es donde se pueden declarar las rutas o middlewares que se usarán en la aplicación.
	//* Ejemplo de cómo podría declararse una ruta: s.engine.GET("/ruta", handlerFunc)

	//* Ejemplo de cómo podría declararse un middleware: s.engine.Use(middleware.Logger())
	srv.engine.Use(middlewares.ConfigurationCors())

	srv.engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
            "message": "Pong!",
        })
	})


	//* Crear el servicio de productos

	userService := services.NewUserService(db)

	routes.UserRoutes(srv.engine.Group("/users"), userService)

    return srv  //* Retornamos la instancia de Server ya configurada
}

func (s *Server) Run() error {
	log.Println("Server running on" + s.host + ":" + s.port)
	return s.engine.Run(s.httpAddr) //* Iniciamos el server en la dirección httpAddr
}
