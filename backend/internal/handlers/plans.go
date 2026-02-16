package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrokunz/gym-manager/backend/internal/repository"
)

type PlanHandler struct {
	Repo repository.Repository
}

func NewPlanHandler(repo repository.Repository) *PlanHandler {
	return &PlanHandler{Repo: repo}
}

func (h *PlanHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := h.Repo.ListPlans()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

func (h *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	// Plan creation using map/interface to match previous loose typing,
	// or define a precise struct. Ideally repo.CreatePlan takes specific arguments.
	// Let's decode into a temp struct matching the JSON expected by frontend.
	var input struct {
		Name           string  `json:"name"`
		Price          float64 `json:"price"` // Frontend sends 'price' or 'total_price'?
		DurationMonths int     `json:"duration_months"`
	}
	// Note: previous implementation used map["price"].
	// Let's assume frontend sends 'price'.

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}

	err := h.Repo.CreatePlan(input.Name, input.Price, input.DurationMonths)
	if err != nil {
		http.Error(w, "Failed to create plan", 500)
		return
	}

	w.WriteHeader(201)
}
