package main

import (
	"fmt"
	"net/http"
	"os"

	driver "project_user-management_ms/common"

	calificacionescontrol "project_user-management_ms/controllers/calificaciones"
	comentarioscontrol "project_user-management_ms/controllers/comentarios"
	estadocuentascontrol "project_user-management_ms/controllers/estadocuentas"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	/*
		dbName := os.Getenv("DB_NAME")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
	*/

	connection, err := driver.ConnectSQL("10.128.0.2", "3005", "root", "123", "userManagement")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	comentariosHandler := comentarioscontrol.NewComentarioHandler(connection)
	calificacionesHandler := calificacionescontrol.NewCalificaionHandler(connection)
	estadocuentasHandler := estadocuentascontrol.NewEstadoCuentaHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/comentarios", ComentarioRouter(comentariosHandler))
		rt.Mount("/calificaciones", CalificacionRouter(calificacionesHandler))
		rt.Mount("/estadocuentas", EstadoCuentaRouter(estadocuentasHandler))
	})

	fmt.Println("Server listen at :5005")
	http.ListenAndServe(":5005", r)
}

// A completely separate router for posts routes
func ComentarioRouter(comentariosHandler *comentarioscontrol.Comentarios) http.Handler {
	r := chi.NewRouter()
	r.Get("/", comentariosHandler.Fetch)
	r.Get("/{idcomento:[0-9]+}/{idcomentado:[0-9]+}", comentariosHandler.GetByID)
	r.Get("/{idcomentado:[0-9]+}", comentariosHandler.GetAllByID)
	r.Post("/", comentariosHandler.Create)
	r.Put("/{idcomento:[0-9]+}/{idcomentado:[0-9]+}", comentariosHandler.Update)
	r.Delete("/{idcomento:[0-9]+}/{idcomentado:[0-9]+}", comentariosHandler.Delete)

	return r
}

func CalificacionRouter(calificacionesHandler *calificacionescontrol.Calificaciones) http.Handler {
	r := chi.NewRouter()
	r.Get("/", calificacionesHandler.Fetch)
	r.Get("/{idcalifico:[0-9]+}/{idcalificado:[0-9]+}", calificacionesHandler.GetByID)
	//r.Get("/{idcalificado:[0-9]+}", calificacionesHandler.GetAVGByID)
	r.Get("/{idcalificado:[0-9]+}", calificacionesHandler.GetAllByID)
	r.Post("/", calificacionesHandler.Create)
	r.Put("/{idcalifico:[0-9]+}/{idcalificado:[0-9]+}", calificacionesHandler.Update)
	r.Delete("/{idcalifico:[0-9]+}/{idcalificado:[0-9]+}", calificacionesHandler.Delete)

	return r
}

func EstadoCuentaRouter(estadocuentasHandler *estadocuentascontrol.EstadoCuentas) http.Handler {
	r := chi.NewRouter()
	r.Get("/", estadocuentasHandler.Fetch)
	r.Get("/{id:[0-9]+}", estadocuentasHandler.GetByID)
	r.Post("/", estadocuentasHandler.Create)
	r.Put("/{id:[0-9]+}", estadocuentasHandler.Update)
	r.Delete("/{id:[0-9]+}", estadocuentasHandler.Delete)

	return r
}
