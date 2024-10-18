package middlewares

import (
	"time"

	"github.com/gin-contrib/cors" //? Importamos la librería cors de gin-contrib para gestionar CORS
	"github.com/gin-gonic/gin"     //? Importamos gin para definir el middleware como un handler
)

//* ConfigurationCors retorna una función middleware que configura CORS para la aplicación Gin
func ConfigurationCors() gin.HandlerFunc {
	//* Retornamos una configuración personalizada de CORS usando gin-contrib/cors
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"}, //* Lista de orígenes permitidos
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},         //* Métodos HTTP permitidos
        AllowHeaders:     []string{"Content-Type", "Authorization"},                 //* Headers permitidos en las solicitudes
        ExposeHeaders:    []string{"Content-Length"},                                //* Headers que se pueden exponer en las respuestas
        AllowCredentials: true,                                                      //* Permitir el uso de cookies o credenciales en las solicitudes
        MaxAge:           12 * time.Hour,                                            //* Tiempo que se almacenará en caché la configuración CORS
	})
}
