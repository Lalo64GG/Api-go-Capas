package main

import (
	"fmt"
	"firstAPI-go/src/server"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main(){
	srv := server.NewServer(HOST, PORT)

	if err := srv.Run(); err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
	}

	fmt.Println("Hello, CompileDaemon!")

}