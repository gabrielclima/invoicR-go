package domain

import (
	"time"
	"github.com/gabrielclima/go_rest_api/utils"
)

// Invoice struct
type Invoice struct {
	Document        int       `json:"document,string"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	ReferenceMounth int       `json:"reference_mounth"`
	ReferenceYear   int       `json:"reference_year"`
	CreatedAt       time.Time `json:"created_at"`
	IsActice        byte      `json:"is_active"`
	DesactiveAt     utils.NullTime  `json:"desactive_at"`
}

// Invoices collection struct
type Invoices []Invoice
