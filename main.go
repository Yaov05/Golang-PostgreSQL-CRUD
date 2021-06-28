package main

import (
	"fmt"
	"log"
)

func main() {

	opt := 4

	switch opt {
	case 1: //Registrar estudiante
		e := Estudiante{
			Nombre: "Alejandro",
			Edad:   21,
			Active: true,
		}
		err := createEstudiante(e)
		if err != nil {
			log.Fatal(err)
		}
	case 2: //Obtener estudiantes
		es, err := getEstudiantes()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(es)
	case 3: //Consultar estudiante por ID
		est, err := getEstudianteByID(5)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(est)
	case 4: //Actualizar estudiante
		e := Estudiante{
			ID:     6,
			Nombre: "Pedro Pérez",
			Edad:   24,
			Active: true,
		}
		err := updateEstudiante(e)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Actualizado correctamente.")
	case 5: //Eliminar estudiante
		err := deleteEstudiante(3)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Eliminado correctamente.")
	}
}
