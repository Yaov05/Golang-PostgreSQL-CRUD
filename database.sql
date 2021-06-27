CREATE TABLE Estudiantes (
    id SERIAL NOT NULL,
    nombre VARCHAR(50) NOT NULL,
    edad SMALLINT NOT NULL,
    active BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);