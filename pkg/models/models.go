package models

import (
	"gorm.io/gorm"
	"time"
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
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	BankName      string         `json:"bank_name"`
	ColumnID      int            `json:"column_id"`
	TaxDeductible bool           `json:"tax_deductible"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
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

type MerchantsAndCategories struct {
	ID            int
	BankName      string
	Name          string
	ColumnID      int
	Column        Category
	TaxDeductible bool
}

type Category struct {
	ID                  int            `json:"id"`
	Name                string         `json:"name"`
	Color               string         `json:"color"`
	BudgetAmount        float64        `json:"budget_amount"`
	BudgetPeriod        string         `json:"budget_period"`
	BudgetGroup         string         `json:"budget_group"`
	WeeklyAmount        float64        `json:"weekly_amount"`
	MonthlyAmount       float64        `json:"monthly_amount"`
	TwiceMonthlyAmount  float64        `json:"twice_monthly_amount"`
	EveryTwoWeeksAmount float64        `json:"every_two_weeks_amount"`
	YearlyAmount        float64        `json:"yearly_amount"`
	ColumnIndex         int            `json:"column_index"`
	Letter              string         `json:"letter"`
	IsCategory          bool           `json:"is_category"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
}

type CategoryForSelect struct {
	ID    int    `json:"id"`
	Value int    `json:"value"`
	Text  string `json:"text"`
}

type BudgetSheet struct {
	Category            string  `json:"category"`
	WeeklyAmount        float64 `json:"weekly_amount"`
	MonthlyAmount       float64 `json:"monthly_amount"`
	TwiceMonthlyAmount  float64 `json:"twice_monthly_amount"`
	EveryTwoWeeksAmount float64 `json:"every_two_weeks_amount"`
	YearlyAmount        float64 `json:"yearly_amount"`
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
