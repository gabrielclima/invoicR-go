package models

import (
	"time"
)

// Invoice struct
type Invoice struct {
	Id              int       `json:"id,string"`
	Document        int       `json:"document,string"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount,string"`
	ReferenceMounth int       `json:"reference_mounth,string"`
	ReferenceYear   int       `json:"reference_year,string"`
	CreatedAt       time.Time `json:"created_at"`
	IsActice        byte      `json:"is_active,string"`
	DesactiveAt     time.Time `json:"-"`
}

// Invoice collection struct
type Invoices []Invoice
