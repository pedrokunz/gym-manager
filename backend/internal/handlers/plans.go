package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrokunz/gym-manager/backend/internal/db"
)

func GetPlans(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.DB.Query("SELECT * FROM plans")
	defer rows.Close()

	plans := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name string
		var price float64
		var months int
		rows.Scan(&id, &name, &price, &months)

		monthly := price / float64(months)

		plans = append(plans, map[string]interface{}{
			"id":              id,
			"name":            name,
			"total_price":     price,
			"months":          months,
			"price_per_month": monthly,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	db.DB.Exec("INSERT INTO plans (name, price, duration_months) VALUES (?, ?, ?)",
		data["name"], data["price"], data["duration_months"])

	w.WriteHeader(201)
}
