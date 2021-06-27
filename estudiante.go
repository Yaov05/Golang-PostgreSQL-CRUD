package main

import (
	"errors"
	"time"
)

//Estructura del estudiante
type Estudiante struct {
	ID        int
	Nombre    string
	Edad      int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Registro de estudiantes
func createEstudiante(e Estudiante) error {
	q := `INSERT INTO 
			Estudiantes (nombre, edad, active) 
			VALUES ($1, $2, $3)`
	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(q) //statement
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(e.Nombre, e.Edad, e.Active)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba 1 línea afectada")
	}
	return nil
}

//Consulta de estudiantes
func getEstudiantes() {

}

//Consulta estudiante en particular
func getEstudianteByID() {

}

//Actualizar estudiante
func updateEstudiante() {

}

//Eliminar estudiante
func deleteEstudiante() {

}
