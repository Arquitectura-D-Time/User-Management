package models

type Comentarios struct {
	IDComento   int64  `json:"idcomento"`
	IDComentado int64  `json:"idcomentado"`
	Comentario  string `json:"comentario"`
	Fecha       string `json:"fecha"`
	Hora        string `json:"hora"`
}

type Calificaciones struct {
	IDCalifico   int64 `json:"idcalifico"`
	IDCalificado int64 `json:"idcalificado"`
	Calificacion int   `json:"calificacion"`
}

type EstadoCuentas struct {
	ID          int64  `json:"idcomento"`
	Estado      int64  `json:"idcomentado"`
	FechaInicio string `json:"comentario"`
	FechaFinal  string `json:"fecha"`
}
