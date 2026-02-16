package repository

import (
	"database/sql"
	"time"
)

type Member struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Status   string    `json:"status"`
	JoinedAt time.Time `json:"joined_at"`
}

type Plan struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"total_price"` // Mapped for frontend compatibility
	DurationMonths int     `json:"months"`
	PricePerMonth  float64 `json:"price_per_month,omitempty"`
}

type Invoice struct {
	ID         int       `json:"id"`
	MemberID   int       `json:"member_id"`
	MemberName string    `json:"member_name"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	Date       time.Time `json:"date"`
}

type Repository interface {
	// Members
	ListMembers(status string, limit, offset int) ([]Member, error)
	CreateMember(m Member) (int, error)
	DeleteMember(id int) error

	// Plans
	ListPlans() ([]Plan, error)
	CreatePlan(name string, price float64, months int) error

	// Invoices
	ListInvoices(limit, offset int) ([]Invoice, error)
	CreateInvoice(inv Invoice) error
	PayInvoice(id int) error

	// Data Management
	SeedData() error
}

type SQLiteRepository struct {
	DB *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{DB: db}
}

// Ensure interface implementation
var _ Repository = &SQLiteRepository{}
