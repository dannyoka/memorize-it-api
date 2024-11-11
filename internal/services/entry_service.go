package services

import (
	"github.com/dannyoka/memorize-it-api/internal/data"
	"github.com/dannyoka/memorize-it-api/internal/repositories"
	"github.com/dannyoka/memorize-it-api/internal/strategies"
)

type EntryService struct {
	entryRepository *repositories.EntryRepository
}

type IEntryService interface {
	GetEntry(id string, strategy string, n int) data.Entry
	CreateEntry(payload data.EntryPayload)
}

func NewEntryService(entryRepository repositories.EntryRepository) *EntryService {
	return &EntryService{
		entryRepository: &entryRepository,
	}
}

// different memorization strategies will be passed in as url params
func (service *EntryService) GetEntry(id string, strategy string, n int) data.Entry {
	passage := service.entryRepository.GetEntry(id)
	switch strategy {
	case "every_nth_word":
		return data.Entry{
			Id:      passage.Id,
			Name:    passage.Name,
			Content: strategies.EveryNthWord(passage.Content, n),
		}
	case "first_letter_of_every_word":
		return data.Entry{
			Id:      passage.Id,
			Name:    passage.Name,
			Content: strategies.FirstLetterOfEveryWord(passage.Content),
		}
	default:
		return data.Entry{
			Id:      passage.Id,
			Name:    passage.Name,
			Content: passage.Content,
		}
	}
}

func (service *EntryService) GetEntries() []data.Entry {
	return service.entryRepository.GetEntries()
}

func (service *EntryService) CreateEntry(entry data.EntryPayload) data.Entry {
	// create entry in repository
	return service.entryRepository.CreateEntry(entry)
}
