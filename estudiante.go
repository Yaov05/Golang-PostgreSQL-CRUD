package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
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

	intNull := sql.NullInt64{}
	strNull := sql.NullString{}

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(q) //statement
	if err != nil {
		return err
	}
	defer stmt.Close()

	//Valores nullos
	if e.Edad == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Edad)
	}
	if e.Nombre == "" {
		strNull.Valid = false
	} else {
		strNull.Valid = true
		strNull.String = e.Nombre
	}

	r, err := stmt.Exec(strNull, intNull, e.Active)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("error: Se esperaba 1 fila afectada")
	}
	fmt.Println("Creado exitosamente")
	return nil
}

//Consulta de estudiantes
func getEstudiantes() (estudiantes []Estudiante, err error) {
	q := `SELECT id, nombre, edad, active, created_at, updated_at 
			FROM estudiantes`

	timeNull := pq.NullTime{}
	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	boolNull := sql.NullBool{}

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Estudiante{}
		err = rows.Scan(
			&e.ID,
			&strNull,
			&intNull,
			&boolNull,
			&e.CreatedAt,
			&timeNull,
		)
		if err != nil {
			return
		}

		e.UpdatedAt = timeNull.Time
		e.Nombre = strNull.String
		e.Edad = int16(intNull.Int64)
		e.Active = boolNull.Bool

		estudiantes = append(estudiantes, e)
	}
	return estudiantes, nil //devuelve los estudiantes sin algún error
}

//Consulta estudiante en particular
func getEstudianteByID(id int) (estudiante []Estudiante, err error) {
	q := `SELECT id, nombre, edad, active, created_at, updated_at 
			FROM Estudiantes 
			WHERE id = $1`

	timeNull := pq.NullTime{}
	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	boolNull := sql.NullBool{}

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		e := Estudiante{}
		err = rows.Scan(
			&e.ID,
			&strNull,
			&intNull,
			&boolNull,
			&e.CreatedAt,
			&timeNull,
		)
		if err != nil {
			return
		}

		e.UpdatedAt = timeNull.Time
		e.Nombre = strNull.String
		e.Edad = int16(intNull.Int64)
		e.Active = boolNull.Bool

		estudiante = append(estudiante, e)
	}
	return estudiante, nil
}

//Actualizar estudiante
func updateEstudiante(e Estudiante) error {
	q := `UPDATE estudiantes 
			SET nombre = $1, edad = $2, active = $3, updated_at = now()
			WHERE id = $4`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(e.Nombre, e.Edad, e.Active, e.ID)
	if err != nil {
		return err
	}
	i, _ := rows.RowsAffected()
	if i != 1 {
		return errors.New("error: Se esperaba 1 fila afectada")
	}
	return nil
}

//Eliminar estudiante
func deleteEstudiante(id int) error {
	q := `DELETE FROM Estudiantes WHERE id = $1`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("error: Se esperaba 1 fila afectada")
	}
	return nil
}
