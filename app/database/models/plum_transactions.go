package models

import (
	"github.com/google/uuid"
	"time"
)

type PlumTransaction struct {
	TransactionID              int `gorm:"primaryKey"`
	Amount                     uint64
	CedingScheme               string
	InvestorReference          uuid.UUID
	FundName                   string
	Product                    string
	Sedol                      string
	ISIN                       string
	TransactionDate            time.Time
	TransactionType            string
	TransactionTypeID          int
	Units                      string
	ClientTransactionReference uuid.UUID
	IDate                      time.Time
}
