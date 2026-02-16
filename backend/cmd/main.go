package main

import (
	"log"
	"net/http"
	"time"

	"github.com/pedrokunz/gym-manager/backend/internal/db"
	"github.com/pedrokunz/gym-manager/backend/internal/graphql"
	"github.com/pedrokunz/gym-manager/backend/internal/handlers"
	"github.com/pedrokunz/gym-manager/backend/internal/middleware"
)

func main() {
	db.InitDB()
	seedData()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.ListMembers(w, r)
		} else if r.Method == "POST" {
			handlers.CreateMember(w, r)
		}
	})

	mux.HandleFunc("/api/members/", handlers.DeleteMember)

	mux.HandleFunc("/api/plans/getall", handlers.GetPlans)
	mux.HandleFunc("/api/plans_create", handlers.CreatePlan)

	mux.HandleFunc("/api/invoices", handlers.GetInvoices)
	mux.HandleFunc("/api/invoices/create", handlers.CreateInvoice)
	mux.HandleFunc("/api/invoices/pay/", handlers.PayInvoice)

	mux.HandleFunc("/graphql", graphql.Handler)

	log.Println("Gym Manager Backend running on :8080")
	if err := http.ListenAndServe(":8080", middleware.AppCORS(mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func seedData() {
	// Check if data already exists
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM plans").Scan(&count)
	if err == nil && count > 0 {
		log.Println("Database already seeded, skipping...")
		return
	}

	log.Println("Seeding database...")

	// db.DB.Exec("DELETE FROM plans")
	// db.DB.Exec("DELETE FROM classes")
	// db.DB.Exec("DELETE FROM members")
	// db.DB.Exec("DELETE FROM invoices")
	// db.DB.Exec("DELETE FROM sqlite_sequence")

	db.DB.Exec("INSERT INTO plans (name, price, duration_months) VALUES ('Pro Monthly', 50.0, 1)")
	db.DB.Exec("INSERT INTO plans (name, price, duration_months) VALUES ('Annual Saver', 480.0, 12)")

	// Seed Member first to get ID
	res, _ := db.DB.Exec("INSERT INTO members (name, email, status, joined_at) VALUES ('Alice Smith', 'alice@example.com', 'active', '2024-01-01')")
	memberID, _ := res.LastInsertId()

	db.DB.Exec("INSERT INTO invoices (member_id, member_name, amount, status, date) VALUES (?, ?, ?, ?, ?)",
		memberID, "Alice Smith", 50.0, "paid", time.Now().AddDate(0, -1, 0).Format(time.RFC3339))
	db.DB.Exec("INSERT INTO invoices (member_id, member_name, amount, status, date) VALUES (?, ?, ?, ?, ?)",
		memberID, "Alice Smith", 50.0, "pending", time.Now().Format(time.RFC3339))

	db.DB.Exec("INSERT INTO classes (name, trainer, schedule) VALUES ('Crossfit 101', 'John Doe', 'Mon/Wed 10:00')")
	db.DB.Exec("INSERT INTO classes (name, trainer, schedule) VALUES ('Yoga Flow', 'Jane Smith', 'Tue/Thu 18:00')")

	// member inserted above
	// db.DB.Exec("INSERT INTO members (name, email, status, joined_at) VALUES ('Alice Smith', 'alice@example.com', 'active', '2024-01-01')")
}
