package main

import (
	"fmt"
	"log"
)

func main() {
	e := Estudiante{
		Nombre: "Alejandro",
		Edad:   30,
		Active: true,
	}
	err := createEstudiante(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado exitosamente")
}
