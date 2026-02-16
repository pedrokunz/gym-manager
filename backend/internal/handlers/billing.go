package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pedrokunz/gym-manager/backend/internal/db"
)

type Invoice struct {
	ID         int       `json:"id"`
	MemberID   int       `json:"member_id"`
	MemberName string    `json:"member_name"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	Date       time.Time `json:"date"`
}

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, member_id, member_name, amount, status, date FROM invoices ORDER BY date DESC")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var i Invoice
		var dateStr string
		err := rows.Scan(&i.ID, &i.MemberID, &i.MemberName, &i.Amount, &i.Status, &dateStr)
		if err == nil {
			i.Date, _ = time.Parse(time.RFC3339, dateStr)
			invoices = append(invoices, i)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var i Invoice
	json.NewDecoder(r.Body).Decode(&i)

	// Basic validation
	if i.MemberID == 0 {
		http.Error(w, "member_id is required", 400)
		return
	}

	_, err := db.DB.Exec("INSERT INTO invoices (member_id, member_name, amount, status, date) VALUES (?, ?, ?, ?, ?)",
		i.MemberID, i.MemberName, i.Amount, "pending", time.Now().Format(time.RFC3339))

	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}

func PayInvoice(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/invoices/pay/")
	id, _ := strconv.Atoi(idStr)

	_, err := db.DB.Exec("UPDATE invoices SET status = 'paid' WHERE id = ?", id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
