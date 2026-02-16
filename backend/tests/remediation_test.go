package tests

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pedrokunz/gym-manager/backend/internal/db"
	"github.com/pedrokunz/gym-manager/backend/internal/handlers"
)

func TestSQLInjection(t *testing.T) {
	// Setup localized DB for testing
	var err error
	db.DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}
	// Initialize minimal schema
	db.DB.Exec("CREATE TABLE members (id INTEGER PRIMARY KEY, name TEXT, email TEXT, status TEXT, joined_at DATETIME)")
	db.DB.Exec("INSERT INTO members (name, email, status) VALUES ('Injected', 'test@test.com', 'active')")

	req, _ := http.NewRequest("GET", "/api/members?status=' OR '1'='1", nil)
	rr := httptest.NewRecorder()

	handlers.ListMembers(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// If vulnerable, it would return all members (including 'active' ones even if we filtered for something else,
	// or empty if the syntax error crashed it).
	// But specifically with ' OR '1'='1, it returns everything.
	// We want to ensure parameterized query handles the input literally.
	// So status="' OR '1'='1" should match NOTHING (status is usually 'active' or 'inactive').

	// With the fix, the query becomes WHERE status = "' OR '1'='1".
	// This should return 0 results (empty array "null" or "[]").

	body := rr.Body.String()
	if strings.Contains(body, "Injected") {
		t.Errorf("SQL Injection successful! Found member when shouldn't have.")
	}
}

func TestInvoiceLinking(t *testing.T) {
	// Setup DB
	var err error
	db.DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}

	db.DB.Exec(`CREATE TABLE members (id INTEGER PRIMARY KEY, name TEXT, email TEXT, status TEXT, joined_at DATETIME);`)
	db.DB.Exec(`CREATE TABLE invoices (id INTEGER PRIMARY KEY, member_id INTEGER, member_name TEXT, amount REAL, status TEXT, date DATETIME, FOREIGN KEY(member_id) REFERENCES members(id));`)

	// Create Member
	res, _ := db.DB.Exec("INSERT INTO members (name, email, status) VALUES ('John Doe', 'john@test.com', 'active')")
	memberID, _ := res.LastInsertId()

	// Create Invoice linked to Member
	// We are testing the handlers logic, but handlers rely on HTTP request body.
	// Let's test the database constraint logic or the handler directly?
	// Let's test the Handler CreateInvoice to ensure it accepts member_id.

	payload := `{"member_id": ` + strings.TrimSpace(string(rune(memberID+48))) + `, "member_name": "John Doe", "amount": 100.0}`
	// Wait, rune conversion is messy. formatting...
	// Just hardcode ID 1.
	db.DB.Exec("DELETE FROM members") // Reset
	db.DB.Exec("INSERT INTO members (id, name, email, status) VALUES (1, 'John Doe', 'john@test.com', 'active')")

	payload = `{"member_id": 1, "member_name": "John Doe", "amount": 100.0}`
	req, _ := http.NewRequest("POST", "/api/invoices/create", strings.NewReader(payload))
	rr := httptest.NewRecorder()

	handlers.CreateInvoice(rr, req)

	if rr.Code != 201 {
		t.Errorf("CreateInvoice failed. Got %d, want 201", rr.Code)
	}

	// Verify Data
	var linkedID int
	err = db.DB.QueryRow("SELECT member_id FROM invoices WHERE amount = 100.0").Scan(&linkedID)
	if err != nil {
		t.Fatalf("Failed to query invoice: %v", err)
	}
	if linkedID != 1 {
		t.Errorf("Invoice linked to wrong member ID. Got %d, want 1", linkedID)
	}
}
