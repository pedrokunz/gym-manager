package services

import (
	"github.com/pedrokunz/gym-manager/backend/internal/repository"
)

type AnalyticsService struct {
	Repo repository.Repository
}

func NewAnalyticsService(repo repository.Repository) *AnalyticsService {
	return &AnalyticsService{Repo: repo}
}

type DashboardData struct {
	TotalMembers   int                  `json:"total_members"`
	ActiveMembers  int                  `json:"active_members"`
	TotalRevenue   float64              `json:"total_revenue"`
	MonthlyRevenue float64              `json:"monthly_revenue"`
	RecentInvoices []repository.Invoice `json:"recent_invoices"`
}

func (s *AnalyticsService) GetDashboardData() (*DashboardData, error) {
	// 1. Members Stats
	// Efficiently we probably want repository methods for counts, but listing works for MVP
	allMembers, err := s.Repo.ListMembers("", 10000, 0) // naive limit
	if err != nil {
		return nil, err
	}

	totalMembers := len(allMembers)
	activeMembers := 0
	for _, m := range allMembers {
		if m.Status == "active" {
			activeMembers++
		}
	}

	// 2. Revenue Stats
	allInvoices, err := s.Repo.ListInvoices(10000, 0) // naive limit
	if err != nil {
		return nil, err
	}

	totalRevenue := 0.0
	// recent := make([]repository.Invoice, 0)
	// Just take top 5 from list (since ListInvoices orders by date DESC)
	recentLimit := 5
	if len(allInvoices) < 5 {
		recentLimit = len(allInvoices)
	}
	recent := allInvoices[:recentLimit]

	for _, inv := range allInvoices {
		if inv.Status == "paid" {
			totalRevenue += inv.Amount
		}
	}

	return &DashboardData{
		TotalMembers:   totalMembers,
		ActiveMembers:  activeMembers,
		TotalRevenue:   totalRevenue,
		MonthlyRevenue: totalRevenue, // simplified for now
		RecentInvoices: recent,
	}, nil
}
