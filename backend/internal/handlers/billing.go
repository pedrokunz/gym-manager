package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/pedrokunz/gym-manager/backend/internal/repository"
)

type BillingHandler struct {
	Repo repository.Repository
}

func NewBillingHandler(repo repository.Repository) *BillingHandler {
	return &BillingHandler{Repo: repo}
}

func (h *BillingHandler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit == 0 {
		limit = 20
	}

	invoices, err := h.Repo.ListInvoices(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func (h *BillingHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var i repository.Invoice
	json.NewDecoder(r.Body).Decode(&i)

	// Basic validation
	if i.MemberID == 0 {
		http.Error(w, "member_id is required", 400)
		return
	}

	// Wait, repository.CreateInvoice takes 'Invoice' struct.
	// We need to set 'pending' and 'time.Now()' somewhere if the repository doesn't.
	// Checking sqlite_repository.go:
	// func (r *SQLiteRepository) CreateInvoice(inv Invoice) error {
	//    _, err := r.DB.Exec("... VALUES (?, ?, ?, ?, ?)", inv.MemberID, inv.MemberName, inv.Amount, "pending", time.Now()...
	// }
	// So repo handles Status and Date! We just pass the partial invoice.

	err := h.Repo.CreateInvoice(i)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}

func (h *BillingHandler) PayInvoice(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/invoices/pay/")
	id, _ := strconv.Atoi(idStr)

	err := h.Repo.PayInvoice(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
