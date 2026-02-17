package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pedrokunz/gym-manager/backend/internal/db"
	"github.com/pedrokunz/gym-manager/backend/internal/graphql"
	"github.com/pedrokunz/gym-manager/backend/internal/handlers"
	"github.com/pedrokunz/gym-manager/backend/internal/middleware"
	"github.com/pedrokunz/gym-manager/backend/internal/repository"
	"github.com/pedrokunz/gym-manager/backend/internal/services"
)

func main() {
	db.InitDB()

	// Initialize Repository
	repo := repository.NewSQLiteRepository(db.DB)

	// Seed Data
	if err := repo.SeedData(); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	// Initialize Handlers
	memberHandler := handlers.NewMemberHandler(repo)
	planHandler := handlers.NewPlanHandler(repo)
	billingHandler := handlers.NewBillingHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			memberHandler.ListMembers(w, r)
		} else if r.Method == "POST" {
			memberHandler.CreateMember(w, r)
		}
	})

	mux.HandleFunc("/api/members/", memberHandler.HandleMemberRequest)

	mux.HandleFunc("/api/plans/getall", planHandler.GetPlans)
	mux.HandleFunc("/api/plans_create", planHandler.CreatePlan)

	mux.HandleFunc("/api/invoices", billingHandler.GetInvoices)
	mux.HandleFunc("/api/invoices/create", billingHandler.CreateInvoice)
	mux.HandleFunc("/api/invoices/pay/", billingHandler.PayInvoice)

	// Analytics
	analyticsService := services.NewAnalyticsService(repo)
	mux.HandleFunc("/api/dashboard", func(w http.ResponseWriter, r *http.Request) {
		data, err := analyticsService.GetDashboardData()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	mux.HandleFunc("/graphql", graphql.Handler)

	log.Println("Gym Manager Backend running on :8080")
	if err := http.ListenAndServe(":8080", middleware.AppCORS(mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
