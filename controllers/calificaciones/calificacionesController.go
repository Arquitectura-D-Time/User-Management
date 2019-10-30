package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "project_user-management_ms/common"

	models "project_user-management_ms/models"

	repository "project_user-management_ms/data"

	calificaiones "project_user-management_ms/data/calificaciones"

	"github.com/go-chi/chi"
)

func NewCalificaionHandler(db *driver.DB) *Calificaciones {
	return &Calificaciones{
		repo: calificaiones.NewSQLCalificacion(db.SQL),
	}
}

// Calificaciones ...
type Calificaciones struct {
	repo repository.Calificaciones
}

// Fetch all calificaiones data
func (c *Calificaciones) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := c.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new calificaiones
func (c *Calificaciones) Create(w http.ResponseWriter, r *http.Request) {
	calificaiones := models.Calificaciones{}
	json.NewDecoder(r.Body).Decode(&calificaiones)

	newID, err := c.repo.Create(r.Context(), &calificaiones)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a calificaiones by id
func (c *Calificaciones) Update(w http.ResponseWriter, r *http.Request) {
	idcalifico, _ := strconv.Atoi(chi.URLParam(r, "idcalifico"))
	idcalificado, _ := strconv.Atoi(chi.URLParam(r, "idcalificado"))
	data := models.Calificaciones{
		IDCalifico:   int64(idcalifico),
		IDCalificado: int64(idcalificado),
	}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := c.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a calificaiones details
func (c *Calificaciones) GetByID(w http.ResponseWriter, r *http.Request) {
	idcalifico, _ := strconv.Atoi(chi.URLParam(r, "idcalifico"))
	idcalificado, _ := strconv.Atoi(chi.URLParam(r, "idcalificado"))
	payload, err := c.repo.GetByID(r.Context(), int64(idcalifico), int64(idcalificado))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (c *Calificaciones) GetAVGByID(w http.ResponseWriter, r *http.Request) {
	idcalificado, _ := strconv.Atoi(chi.URLParam(r, "idcalificado"))
	payload, _ := c.repo.Fetch(r.Context(), int64(idcalificado))

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a calificaiones
func (c *Calificaciones) Delete(w http.ResponseWriter, r *http.Request) {
	idcalifico, _ := strconv.Atoi(chi.URLParam(r, "idcalifico"))
	idcalificado, _ := strconv.Atoi(chi.URLParam(r, "idcalificado"))
	_, err := c.repo.Delete(r.Context(), int64(idcalifico), int64(idcalificado))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
