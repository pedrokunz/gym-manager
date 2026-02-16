package repository

import (
	"errors"
	"time"
)

type MemoryRepository struct {
	Members  []Member
	Plans    []Plan
	Invoices []Invoice
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		Members:  []Member{},
		Plans:    []Plan{},
		Invoices: []Invoice{},
	}
}

func (r *MemoryRepository) ListMembers(status string, limit, offset int) ([]Member, error) {
	var filtered []Member
	if status == "" {
		filtered = r.Members
	} else {
		for _, m := range r.Members {
			if m.Status == status {
				filtered = append(filtered, m)
			}
		}
	}

	total := len(filtered)
	if offset >= total {
		return []Member{}, nil
	}

	end := total
	if limit > 0 {
		end = offset + limit
		if end > total {
			end = total
		}
	}

	return filtered[offset:end], nil
}

func (r *MemoryRepository) CreateMember(m Member) (int, error) {
	m.ID = len(r.Members) + 1
	m.JoinedAt = time.Now()
	r.Members = append(r.Members, m)
	return m.ID, nil
}

func (r *MemoryRepository) DeleteMember(id int) error {
	for i, m := range r.Members {
		if m.ID == id {
			r.Members = append(r.Members[:i], r.Members[i+1:]...)
			return nil
		}
	}
	return errors.New("member not found")
}

func (r *MemoryRepository) ListPlans() ([]Plan, error) {
	return r.Plans, nil
}

func (r *MemoryRepository) CreatePlan(name string, price float64, months int) error {
	p := Plan{
		ID:             len(r.Plans) + 1,
		Name:           name,
		Price:          price,
		DurationMonths: months,
	}
	r.Plans = append(r.Plans, p)
	return nil
}

func (r *MemoryRepository) ListInvoices(limit, offset int) ([]Invoice, error) {
	total := len(r.Invoices)
	if offset >= total {
		return []Invoice{}, nil
	}

	end := total
	if limit > 0 {
		end = offset + limit
		if end > total {
			end = total
		}
	}
	return r.Invoices[offset:end], nil
}

func (r *MemoryRepository) CreateInvoice(inv Invoice) error {
	inv.ID = len(r.Invoices) + 1
	inv.Date = time.Now()
	r.Invoices = append(r.Invoices, inv)
	return nil
}

func (r *MemoryRepository) PayInvoice(id int) error {
	for i, inv := range r.Invoices {
		if inv.ID == id {
			r.Invoices[i].Status = "paid"
			return nil
		}
	}
	return errors.New("invoice not found")
}

func (r *MemoryRepository) SeedData() error {
	// No-op for memory repo, or pre-fill
	return nil
}
