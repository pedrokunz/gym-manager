package repository

import (
	"log"
	"time"
)

// -- Members --

func (r *SQLiteRepository) ListMembers(statusFilter string, limit, offset int) ([]Member, error) {
	query := "SELECT id, name, email FROM members"
	var args []interface{}

	if statusFilter != "" {
		query += " WHERE status = ?"
		args = append(args, statusFilter)
	}

	if limit > 0 {
		query += " LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var m Member
		if err := rows.Scan(&m.ID, &m.Name, &m.Email); err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

func (r *SQLiteRepository) CreateMember(m Member) (int, error) {
	res, err := r.DB.Exec("INSERT INTO members (name, email, status, joined_at) VALUES (?, ?, ?, ?)",
		m.Name, m.Email, "active", time.Now())
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *SQLiteRepository) DeleteMember(id int) error {
	_, err := r.DB.Exec("DELETE FROM members WHERE id = ?", id)
	return err
}

// -- Plans --

func (r *SQLiteRepository) ListPlans() ([]Plan, error) {
	rows, err := r.DB.Query("SELECT id, name, price, duration_months FROM plans")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []Plan
	for rows.Next() {
		var p Plan
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.DurationMonths); err != nil {
			return nil, err
		}
		if p.DurationMonths > 0 {
			p.PricePerMonth = p.Price / float64(p.DurationMonths)
		}
		plans = append(plans, p)
	}
	return plans, nil
}

func (r *SQLiteRepository) CreatePlan(name string, price float64, months int) error {
	_, err := r.DB.Exec("INSERT INTO plans (name, price, duration_months) VALUES (?, ?, ?)",
		name, price, months)
	return err
}

// -- Invoices --

func (r *SQLiteRepository) ListInvoices(limit, offset int) ([]Invoice, error) {
	query := "SELECT id, member_id, member_name, amount, status, date FROM invoices ORDER BY date DESC"
	var args []interface{}

	if limit > 0 {
		query += " LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var i Invoice
		var dateStr string
		if err := rows.Scan(&i.ID, &i.MemberID, &i.MemberName, &i.Amount, &i.Status, &dateStr); err != nil {
			return nil, err
		}
		i.Date, _ = time.Parse(time.RFC3339, dateStr)
		invoices = append(invoices, i)
	}
	return invoices, nil
}

func (r *SQLiteRepository) CreateInvoice(inv Invoice) error {
	_, err := r.DB.Exec("INSERT INTO invoices (member_id, member_name, amount, status, date) VALUES (?, ?, ?, ?, ?)",
		inv.MemberID, inv.MemberName, inv.Amount, "pending", time.Now().Format(time.RFC3339))
	return err
}

func (r *SQLiteRepository) PayInvoice(id int) error {
	_, err := r.DB.Exec("UPDATE invoices SET status = 'paid' WHERE id = ?", id)
	return err
}

// -- Data Management --

func (r *SQLiteRepository) SeedData() error {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM plans").Scan(&count)
	if err == nil && count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	log.Println("Seeding database via Repository...")

	// Ideally execute these in a transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	tx.Exec("INSERT INTO plans (name, price, duration_months) VALUES ('Pro Monthly', 50.0, 1)")
	tx.Exec("INSERT INTO plans (name, price, duration_months) VALUES ('Annual Saver', 480.0, 12)")

	res, _ := tx.Exec("INSERT INTO members (name, email, status, joined_at) VALUES ('Alice Smith', 'alice@example.com', 'active', '2024-01-01')")
	memberID, _ := res.LastInsertId()

	tx.Exec("INSERT INTO invoices (member_id, member_name, amount, status, date) VALUES (?, ?, ?, ?, ?)",
		memberID, "Alice Smith", 50.0, "paid", time.Now().AddDate(0, -1, 0).Format(time.RFC3339))
	tx.Exec("INSERT INTO invoices (member_id, member_name, amount, status, date) VALUES (?, ?, ?, ?, ?)",
		memberID, "Alice Smith", 50.0, "pending", time.Now().Format(time.RFC3339))

	tx.Exec("INSERT INTO classes (name, trainer, schedule) VALUES ('Crossfit 101', 'John Doe', 'Mon/Wed 10:00')")
	tx.Exec("INSERT INTO classes (name, trainer, schedule) VALUES ('Yoga Flow', 'Jane Smith', 'Tue/Thu 18:00')")

	return tx.Commit()
}

// GetMember retrieves a single member by ID
func (r *SQLiteRepository) GetMember(id int) (*Member, error) {
	row := r.DB.QueryRow("SELECT id, name, email, status, joined_at FROM members WHERE id = ?", id)
	var m Member
	var joinedAtStr string
	if err := row.Scan(&m.ID, &m.Name, &m.Email, &m.Status, &joinedAtStr); err != nil {
		return nil, err
	}
	// Try parsing standard formats
	parsedTime, err := time.Parse(time.RFC3339, joinedAtStr)
	if err != nil {
		// Try YYYY-MM-DD
		parsedTime, _ = time.Parse("2006-01-02", joinedAtStr)
	}
	m.JoinedAt = parsedTime
	return &m, nil
}

// ListInvoicesByMember retrieves invoices for a specific member
func (r *SQLiteRepository) ListInvoicesByMember(memberID int, limit, offset int) ([]Invoice, error) {
	query := "SELECT id, member_id, member_name, amount, status, date FROM invoices WHERE member_id = ? ORDER BY date DESC"
	var args []interface{}
	args = append(args, memberID)

	if limit > 0 {
		query += " LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var i Invoice
		var dateStr string
		if err := rows.Scan(&i.ID, &i.MemberID, &i.MemberName, &i.Amount, &i.Status, &dateStr); err != nil {
			return nil, err
		}
		// Try parsing standard formats
		parsedTime, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			parsedTime, _ = time.Parse("2006-01-02", dateStr)
		}
		i.Date = parsedTime
		invoices = append(invoices, i)
	}
	return invoices, nil
}
