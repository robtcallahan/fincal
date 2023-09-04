package models

import (
	"gorm.io/gorm"
)

// Transaction ...
type Transaction struct {
	gorm.Model
	Key            string
	Source         string
	Date           string
	Name           string
	BankName       string
	Note           string
	Amount         float64 // The actual transaction value
	Withdrawal     float64 // the amount that goes in the Withdrawal column (positive)
	Deposit        float64 // the amount that goes in the Deposit column (positive)
	CreditPurchase float64 // the amount that goes in the Credit Purchases column (positive)
	Budget         float64 // the amount that goes in the Budget category column (negative)
	CreditCard     float64 // the amount that goes in the Credit Card column (positive)
	ColumnIndex    int
	Color          string
	IsCategory     bool
	TaxDeductible  bool
	IsCheck        bool
}

// Merchant ...
type Merchant struct {
	gorm.Model
	ID            int    `json:"id"`
	Name          string `json:"name"`
	BankName      string `json:"bank_name"`
	ColumnID      int    `json:"column_id"`
	TaxDeductible bool   `json:"tax_deductible"`
}

type MerchantColumn struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	BankName         string `json:"bank_name"`
	ColumnID         int    `json:"column_id"`
	ColumnName       string `json:"column_name"`
	ColumnColor      string `json:"column_color"`
	ColumnIsCategory bool   `json:"column_is_category"`
	TaxDeductible    bool   `json:"tax_deductible"`
}

type MerchantsAndColumns struct {
	ID            int
	BankName      string
	Name          string
	ColumnID      int
	Column        Column
	TaxDeductible bool
}

// Column ...
type Column struct {
	gorm.Model
	ID          int
	Name        string
	Color       string
	ColumnIndex int
	Letter      string
	IsCategory  bool
}

// DataRow ...
type DataRow struct {
	ID            int
	Name          string
	BankName      string
	ColumnName    string
	ColumnIndex   int
	Color         string
	IsCategory    bool
	TaxDeductible bool
}
