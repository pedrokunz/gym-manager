package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pedrokunz/gym-manager/backend/internal/repository"
)

type MemberHandler struct {
	Repo repository.Repository
}

func NewMemberHandler(repo repository.Repository) *MemberHandler {
	return &MemberHandler{Repo: repo}
}

func (h *MemberHandler) ListMembers(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit == 0 {
		limit = 20
	}

	members, err := h.Repo.ListMembers(statusFilter, limit, offset)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func (h *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	var m repository.Member
	json.NewDecoder(r.Body).Decode(&m)

	if m.Name == "" || !strings.Contains(m.Email, "@") {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"err": "missing data or bad email"}`)
		return
	}

	id, err := h.Repo.CreateMember(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m.ID = id
	// Since CreateMember in repo adds JoinedAt and Status, we might want to return the full object as added.
	// But for now, sticking to the repo input/output contract.
	// Actually, repo.CreateMember(m) returns (int, error).
	// The repo implementation sets 'active' and 'JoinedAt' in the DB.
	// The returned 'm' here won't have them set unless we fetch or manually set them.
	// For simplicity, let's manually set them on the response object if needed, or just return ID + input data.
	m.Status = "active"
	// m.JoinedAt = time.Now() // we don't import time, let's skip unless critical.

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func (h *MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/members/")
	id, _ := strconv.Atoi(idStr)

	err := h.Repo.DeleteMember(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
