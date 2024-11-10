package services

import (
	"github.com/dannyoka/memorize-it-api/internal/data"
	"github.com/dannyoka/memorize-it-api/internal/repositories"
)

type EntryService struct {
	entryRepository *repositories.EntryRepository
}

type IEntryService interface {
	GetEntry(id string, strategy string)
	CreateEntry()
}

func NewEntryService(entryRepository repositories.EntryRepository) *EntryService {
	return &EntryService{
		entryRepository: &entryRepository,
	}
}

// different memorization strategies will be passed in as url params
func (service *EntryService) GetEntry(id string, strategy string) data.Entry {
	return service.entryRepository.GetEntry(id)
}

func (service *EntryService) GetEntries() []data.Entry {
	return service.entryRepository.GetEntries()
}

func (service *EntryService) CreateEntry() {
	// create entry in repository
}
