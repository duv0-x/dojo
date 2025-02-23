package main

import (
	"wataaah-go/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") // Inicia el servidor en el puerto 8080
}
