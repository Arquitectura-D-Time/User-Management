package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "User-Management/common"

	models "User-Management/models"

	repository "User-Management/data"

	comentarios "User-Management/data/comentarios"

	"github.com/go-chi/chi"
)

func NewComentarioHandler(db *driver.DB) *Comentarios {
	return &Comentarios{
		repo: comentarios.NewSQLComentario(db.SQL),
	}
}

// Comentarios ...
type Comentarios struct {
	repo repository.Comentarios
}

// Fetch all comentarios data
func (c *Comentarios) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := c.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new comentarios
func (c *Comentarios) Create(w http.ResponseWriter, r *http.Request) {
	comentarios := models.Comentarios{}
	json.NewDecoder(r.Body).Decode(&comentarios)

	newID, err := c.repo.Create(r.Context(), &comentarios)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a comentarios by id
func (c *Comentarios) Update(w http.ResponseWriter, r *http.Request) {
	idcomento, _ := strconv.Atoi(chi.URLParam(r, "IDComento"))
	idcomentado, _ := strconv.Atoi(chi.URLParam(r, "IDComentado"))
	data := models.Comentarios{
		IDComento: int64(idcomento)
		IDComentado: int64(idcomentado)
	}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := c.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a comentarios details
func (c *Comentarios) GetByID(w http.ResponseWriter, r *http.Request) {
	idcomento, _ := strconv.Atoi(chi.URLParam(r, "IDComento"))
	idcomentado, _ := strconv.Atoi(chi.URLParam(r, "IDComentado"))
	payload, err := c.repo.GetByID(r.Context(), int64(idcomento), int64(idcomentado))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a comentarios
func (c *Comentarios) Delete(w http.ResponseWriter, r *http.Request) {
	idcomento, _ := strconv.Atoi(chi.URLParam(r, "IDComento"))
	idcomentado, _ := strconv.Atoi(chi.URLParam(r, "IDComentado"))
	_, err := c.repo.Delete(r.Context(), int64(idcomento), int64(idcomentado))

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
