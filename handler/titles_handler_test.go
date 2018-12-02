package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mvanbrummen/go-rest/models"
)

type StubTitlesRepository struct{}

func (*StubTitlesRepository) SearchByTitle(title string, limit int) ([]*models.Title, error) {
	return []*models.Title{
		&models.Title{
			"E12345",
			"movie",
			"Labryinth",
			"Labryinth",
			true,
			new(int64),
			nil,
			nil,
			[]string{"action"},
		},
	}, nil
}

func (*StubTitlesRepository) FetchTitle(id string) (*models.Title, error) {
	return &models.Title{
		"E12345",
		"movie",
		"Labryinth",
		"Labryinth",
		true,
		new(int64),
		nil,
		nil,
		[]string{"action"},
	}, nil
}

func TestGetTitle(t *testing.T) {
	req, err := http.NewRequest("GET", "/titles/poo", nil)
	if err != nil {
		t.Fatal(err)
	}

	titlesHandler := NewTitlesHandler(&StubTitlesRepository{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(titlesHandler.GetTitle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":"E12345","title_type":"movie","primary_title":"Labryinth","original_title":"Labryinth","is_adult":true,"start_year":0,"end_year":null,"runtime_minutes":null,"genres":["action"]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
