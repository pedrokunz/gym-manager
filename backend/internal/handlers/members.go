package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pedrokunz/gym-manager/backend/internal/db"
	"github.com/pedrokunz/gym-manager/backend/internal/models"
)

func ListMembers(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")

	query := "SELECT id, name, email FROM members"
	if statusFilter != "" {
		query += " WHERE status = '" + statusFilter + "'"
	}

	rows, err := db.DB.Query(query)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var m models.Member
		rows.Scan(&m.ID, &m.Name, &m.Email)
		members = append(members, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func CreateMember(w http.ResponseWriter, r *http.Request) {
	var m models.Member
	json.NewDecoder(r.Body).Decode(&m)

	if m.Name == "" || !strings.Contains(m.Email, "@") {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"err": "missing data or bad email"}`)
		return
	}

	res, err := db.DB.Exec("INSERT INTO members (name, email, status, joined_at) VALUES (?, ?, ?, ?)",
		m.Name, m.Email, "active", time.Now())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	m.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/members/")
	id, _ := strconv.Atoi(idStr)

	_, err := db.DB.Exec("DELETE FROM members WHERE id = " + strconv.Itoa(id))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
