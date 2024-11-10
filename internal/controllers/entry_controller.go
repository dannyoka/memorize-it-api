package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dannyoka/memorize-it-api/internal/services"
)

type EntryController struct {
	entryService *services.EntryService
}

type IEntryController interface {
	HandleGetEntries(w http.ResponseWriter, r *http.Request)
}

func NewEntryController(service services.EntryService) *EntryController {
	return &EntryController{
		entryService: &service,
	}
}

func (controller *EntryController) HandleGetEntries(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	strategy := r.URL.Query().Get("strategy")
	fmt.Println("id", id)
	fmt.Println("strategy", strategy)
	entries := controller.entryService.GetEntries()
	jsonData, err := json.Marshal(entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonData))
}
