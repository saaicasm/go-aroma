package mocks

import (
	"github/saaicasm/snipbox/internal/models"
	"time"
)

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "Mock",
	Content: "Fake, Stub",
	Created: time.Now(),
	Expires: <-time.After(10),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(id int, title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord
	}

}

func (m *SnippetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}
