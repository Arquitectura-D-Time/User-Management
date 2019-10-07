DROP DATABASE IF exists userManagement;
CREATE DATABASE userManagement character set utf8;

USE userManagement;

CREATE TABLE Comentarios (
	IDComento BIGINT NOT NULL,
    IDComentado BIGINT NOT NULL,
    Comentario VARCHAR(255),
    Fecha VARCHAR(255),
    Hora VARCHAR(255)
);

CREATE TABLE Calificaciones (
	IDCalifico BIGINT NOT NULL,
    IDCalificado BIGINT NOT NULL,
    Calificacion INT
);

CREATE TABLE EstadoCuentas (
	ID BIGINT NOT NULL,
	Estado VARCHAR(255),
    FechaInicio VARCHAR(255),
    FechaFinal VARCHAR(255)
);