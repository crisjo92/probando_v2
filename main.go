package main

import (
	"log"
	"probando2/server"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	srv := server.New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
