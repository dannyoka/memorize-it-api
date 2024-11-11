package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dannyoka/memorize-it-api/internal/data"
	"github.com/dannyoka/memorize-it-api/internal/services"
	"github.com/gorilla/mux"
)

type EntryController struct {
	entryService *services.EntryService
}

type IEntryController interface {
	HandleGetEntries(w http.ResponseWriter, r *http.Request)
	HandleGetEntry(w http.ResponseWriter, r *http.Request)
	HandleCreateEntry(w http.ResponseWriter, r *http.Request)
}

func NewEntryController(service services.EntryService) IEntryController {
	return &EntryController{
		entryService: &service,
	}
}

func (controller *EntryController) HandleGetEntries(w http.ResponseWriter, r *http.Request) {
	entries := controller.entryService.GetEntries()
	jsonData, err := json.Marshal(entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonData))
}

func (controller *EntryController) HandleGetEntry(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	strategy := r.URL.Query().Get("strategy")
	n := r.URL.Query().Get("n")
	if n == "" {
		n = "2"
	}
	nInt, err := strconv.Atoi(n)
	if err != nil {
		nInt = 0
	}
	entry := controller.entryService.GetEntry(id, strategy, nInt)
	jsonData, err := json.Marshal(entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonData))
}

func (controller *EntryController) HandleCreateEntry(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var entry data.EntryPayload
	err := decoder.Decode(&entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newEntry := controller.entryService.CreateEntry(entry)
	jsonData, err := json.Marshal(newEntry)
	w.Write([]byte(jsonData))
}
