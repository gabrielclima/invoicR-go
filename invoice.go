package main

import (
	"time"
)

type Invoice struct {
	Document        int       `json:"document"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	ReferenceMounth int       `json:"reference_mounth"`
	ReferenceYear   int       `json:"reference_year"`
	CreatedAt       time.Time `json:"created_at"`
	IsActice        byte      `json:"is_active"`
	DesactiveAt     NullTime  `json:"desactive_at"`
}

type Invoices []Invoice
