package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Secrets struct {
	ID          int            `json:"id"`
	UsersID     int            `json:"users_id"`
	Password    []byte         `json:"password"`
	Token       string         `json:"token"`
	TokenExpiry sql.NullTime   `json:"token_expiry"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

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
	CategoryID     int
	Color          string
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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BankName    string `json:"bank_name"`
	ColumnID    int    `json:"column_id"`
	ColumnName  string `json:"column_name"`
	ColumnColor string `json:"column_color"`
}

type MerchantsAndCategories struct {
	ID       int
	BankName string
	Name     string
	Category Category
}

type Category struct {
	ID                    int            `json:"id"`
	ColumnIndex           float64        `json:"column_index"`
	Name                  string         `json:"name"`
	Color                 string         `json:"color"`
	BudgetGroup           string         `json:"budget_group"`
	BudgetAmount          float64        `json:"budget_amount"`
	BudgetPeriod          string         `json:"budget_period"`
	PayPeriodContribution float64        `json:"pay_period_contribution"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
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
	ID         int
	Name       string
	BankName   string
	ColumnName string
	CategoryID int
	Color      string
}
