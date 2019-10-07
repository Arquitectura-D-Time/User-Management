package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "User-Management/common"

	models "User-Management/models"

	repository "User-Management/data"

	estadocuentas "User-Management/data/estadocuentas"

	"github.com/go-chi/chi"
)

func NewEstadoCuentaHandler(db *driver.DB) *EstadoCuentas {
	return &EstadoCuentas{
		repo: estadocuentas.NewSQLEstadoCuentas(db.SQL),
	}
}

// EstadoCuentas ...
type EstadoCuentas struct {
	repo repository.EstadoCuentas
}

// Fetch all estadocuentas data
func (c *EstadoCuentas) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := c.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new estadocuentas
func (c *EstadoCuentas) Create(w http.ResponseWriter, r *http.Request) {
	estadocuentas := models.EstadoCuentas{}
	json.NewDecoder(r.Body).Decode(&estadocuentas)

	newID, err := c.repo.Create(r.Context(), &estadocuentas)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a estadocuentas by id
func (c *EstadoCuentas) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "ID"))
	data := models.EstadoCuentas{
		ID: int64(id),
	}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := c.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a estadocuentas details
func (c *EstadoCuentas) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "ID"))
	payload, err := c.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a estadocuentas
func (c *EstadoCuentas) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "ID"))
	_, err := c.repo.Delete(r.Context(), int64(id))

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
