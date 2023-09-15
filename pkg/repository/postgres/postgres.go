package postgres

import (
	"fmt"
	"gorm.io/gorm/clause"

	"vue-register/pkg/models"
	repo "vue-register/pkg/repository"

	"gorm.io/gorm"
)

type postgresQueryRepo struct {
	Conn *gorm.DB
}

// NewPostgreSQLQueryRepo returns the implementation of post repository interface
func NewPostgreSQLQueryRepo(conn *gorm.DB) repo.QueryRepo {
	return &postgresQueryRepo{
		Conn: conn,
	}
}

func (r *postgresQueryRepo) GetTransactions() []models.Transaction {
	var trans []models.Transaction
	r.Conn.Order("date").Find(&trans)
	return trans
}

func (r *postgresQueryRepo) SaveTransaction(trans *models.Transaction) {
	r.Conn.Save(trans)
}

func (r *postgresQueryRepo) UpdateTransactionTables(trans []*models.Transaction) {
	_ = r.Conn.AutoMigrate(&models.Transaction{})

	for _, t := range trans {
		result := r.Conn.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(t)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}

func (r *postgresQueryRepo) CreateDB(dbName string) (*gorm.DB, error) {
	db := r.Conn.Exec("CREATE DATABASE " + dbName)
	return db, db.Error
}

func (r *postgresQueryRepo) GetCategories() []models.Category {
	var cols []models.Category
	r.Conn.Order("column_index").Find(&cols)
	return cols
}

func (r *postgresQueryRepo) CreateCategory(category *models.Category) *models.Category {
	tx := r.Conn.Create(category)
	category.ID = tx.Statement.Model.(*models.Category).ID
	return category
}

func (r *postgresQueryRepo) UpdatesCategory(category *models.Category, values interface{}) {
	r.Conn.Model(&category).Where("id = ?", category.ID).Updates(values)
}

func (r *postgresQueryRepo) GetCategoriesForSelect() []models.CategoryForSelect {
	var cols []models.Category
	r.Conn.Where("is_category = 1").Order("name").Find(&cols)

	var catsSelect []models.CategoryForSelect
	for _, c := range cols {
		s := models.CategoryForSelect{
			ID:    c.ID,
			Value: c.ID,
			Text:  c.Name,
		}
		catsSelect = append(catsSelect, s)
	}
	return catsSelect
}

// GetMerchants ...
func (r *postgresQueryRepo) GetMerchants() []models.Merchant {
	var merch []models.Merchant
	r.Conn.Order("name").Find(&merch)
	return merch
}

func (r *postgresQueryRepo) GetMerchantsAndCategories() []models.MerchantsAndCategories {
	var merch []models.MerchantsAndCategories
	r.Conn.Order("name").Find(&merch)
	return merch
}

// CreateMerchant ...
func (r *postgresQueryRepo) CreateMerchant(m *models.Merchant) {
	result := r.Conn.Create(&models.Merchant{
		Name:     m.Name,
		BankName: m.BankName,
		ColumnID: m.ColumnID,
	})
	if result.Error != nil {
		panic(result.Error)
	}
}

func (r *postgresQueryRepo) UpdateMerchant(m *models.Merchant, column, value string) {
	r.Conn.Model(m).Update(column, value)
}

func (r *postgresQueryRepo) DeleteMerchant(m *models.Merchant) {
	r.Conn.Delete(m)
}

// GetLookupData ...
func (r *postgresQueryRepo) GetLookupData() []*models.DataRow {
	var merchants []models.MerchantsAndCategories

	r.Conn.Preload("Category").Find(&merchants)

	var data []*models.DataRow
	for _, m := range merchants {
		data = append(data, &models.DataRow{
			Name:        m.Name,
			BankName:    m.BankName,
			ColumnName:  m.Column.Name,
			ColumnIndex: m.Column.ColumnIndex,
			Color:       m.Column.Color,
			IsCategory:  m.Column.IsCategory,
		})
	}
	return data
}

// GetNameMapToColumn creates a map lookup from trans name to budget category/column names
func (r *postgresQueryRepo) GetNameMapToColumn() map[string]string {
	cols := r.GetLookupData()

	transNameToColName := make(map[string]string)
	for _, c := range cols {
		transNameToColName[c.Name] = c.ColumnName
	}
	return transNameToColName
}

// PrintData ...
func (r *postgresQueryRepo) PrintData() {
	var merchants []models.MerchantsAndCategories
	r.Conn.Preload("Category").Find(&merchants)

	fmt.Printf("[Num] %-35s %-30s %-30s %-s\n", "Bank Name", "Name", "Category Name", "Category Index")
	for i, m := range merchants {
		fmt.Printf("[%3d] %-35s %-30s %-30s %2d\n", i+1, m.BankName, m.Name, m.Column.Name, m.Column.ColumnIndex)
	}
}

// PrintTable ...
func (r *postgresQueryRepo) PrintTable(table string) {
	switch table {
	case "merchants":
		var merchants []models.MerchantsAndCategories
		result := r.Conn.Find(&merchants)
		fmt.Printf("%d rows found\n", result.RowsAffected)
		for _, m := range merchants {
			fmt.Printf("%d %s %s %s\n", m.ID, m.BankName, m.Name, m.Column.Name)
		}
	}
}
