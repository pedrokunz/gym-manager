package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedrokunz/gym-manager/backend/internal/db"
	"github.com/pedrokunz/gym-manager/backend/internal/handlers"
	"github.com/pedrokunz/gym-manager/backend/internal/repository"
)

func TestMain(m *testing.M) {
	db.InitDB()
	db.DB.Exec("DELETE FROM members") // Clean start
	db.DB.Exec("INSERT INTO members (name, email, status, joined_at) VALUES ('Test User', 'test@example.com', 'active', '2024-01-01')")
	m.Run()
}

func TestListMembers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/members", nil)
	rr := httptest.NewRecorder()

	repo := repository.NewSQLiteRepository(db.DB)
	handler := handlers.NewMemberHandler(repo)

	// Create a handler function calling the method
	httpHandler := http.HandlerFunc(handler.ListMembers)

	httpHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var members []interface{}
	json.Unmarshal(rr.Body.Bytes(), &members)

	if len(members) == 0 {
		t.Errorf("Empty members list returned")
	}
}

func TestGetPlans(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/plans/getall", nil)
	rr := httptest.NewRecorder()

	repo := repository.NewSQLiteRepository(db.DB)
	handler := handlers.NewPlanHandler(repo)
	httpHandler := http.HandlerFunc(handler.GetPlans)

	httpHandler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Expected 200, got %v", rr.Code)
	}
}
