package main

import (
	"log"

	"github.com/Nicol68/Nicopre/bd"
	"github.com/Nicol68/Nicopre/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
