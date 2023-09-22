package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"

	cfg "fincal/pkg/config"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doImport()
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}

// Column ...
type Column struct {
	gorm.Model
	Name       string
	Color      string
	Index      int
	IsCategory bool
}

// Merchant ...
type Merchant struct {
	gorm.Model
	BankName string
	Name     string
	ColumnID uint
}

// ColumnCSVRow ...
type ColumnCSVRow struct {
	Name  string `csv:"Name"`
	Color string `csv:"Color"`
}

// MerchToColCSVRow ...
type MerchToColCSVRow struct {
	Merchant   string `csv:"Merchant"`
	ColumnName string `csv:"Column Name"`
}

// MerchantCSVRow ...
type MerchantCSVRow struct {
	BankName string `csv:"Bank Name"`
	Name     string `csv:"Name"`
}

// ColLookup ...
type ColLookup map[string]string

// Config ...
type Config struct {
	AppConfig  *cfg.Config
	DBName     string
	DBUsername string
	DBPassword string
	DB         *gorm.DB
}

func doImport() {
	var err error

	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dbHost := fmt.Sprintf("%s", viper.Get("DB_HOST"))

	config := &Config{
		AppConfig:  config,
		DBName:     config.DBName,
		DBUsername: config.DBUsername,
		DBPassword: config.DBPassword,
	}

	dsn := config.DBUsername + ":" + config.DBPassword + "@tcp(" + dbHost + ":3306)/" +
		config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Creating the database...")
	config.runSQL("../../sql/register.sql")

	fmt.Println("Importing 'columns' table data...")
	config.importColumns()

	fmt.Println("Importing 'merchants to categories' table data...")
	colLookup := config.importMerchToCats()

	fmt.Println("Importing 'merchants' table data...")
	config.importMerchants(colLookup)
}

func (c *Config) importMerchants(cl ColLookup) {
	c.DB.AutoMigrate(&Merchant{})

	f, err := os.Open("merchants.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	rows := []*MerchantCSVRow{}
	if err := gocsv.UnmarshalFile(f, &rows); err != nil {
		panic(err)
	}

	for _, r := range rows {
		colName := cl[r.Name]
		if colName == "" {
			fmt.Printf("    column not found for %s\n", r.Name)
			continue
		}

		col := Column{}
		c.DB.Where("name = ?", colName).First(&col)

		c.DB.Create(&Merchant{
			Name:     r.Name,
			BankName: r.BankName,
			ColumnID: col.ID,
		})
	}
}

func (c *Config) importMerchToCats() ColLookup {
	c.DB.AutoMigrate(&Column{})

	f, err := os.Open("merch_to_cats.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	rows := []*MerchToColCSVRow{}
	if err := gocsv.UnmarshalFile(f, &rows); err != nil {
		panic(err)
	}

	var cl ColLookup
	cl = make(ColLookup)
	for _, r := range rows {
		cl[r.Merchant] = r.ColumnName
		// fmt.Printf("c[%s] = %s\n", r.Merchant, r.ColumnName)
	}

	return cl
}

func (c *Config) importColumns() {
	c.DB.AutoMigrate(&Column{})

	f, err := os.Open("columns.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	rows := []*ColumnCSVRow{}
	if err := gocsv.UnmarshalFile(f, &rows); err != nil {
		panic(err)
	}

	var name string
	var isCat bool
	for i, r := range rows {
		name = r.Name
		isCat = true
		if i < 10 {
			isCat = false
		}
		if name == "" {
			name = fmt.Sprintf("old-%d", i)
		}
		c.DB.Create(&Column{
			Index:      i,
			Name:       name,
			Color:      r.Color,
			IsCategory: isCat,
		})
	}
}

func (c *Config) runSQL(file string) {
	sql, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	cmd := &exec.Cmd{
		Path: "/usr/local/mysql/bin/mysql",
		Args: []string{"/usr/local/mysql/bin/mysql", "--silent",
			"-u", c.DBUsername, "--password=" + c.DBPassword, "-D", c.DBName},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  sql,
	}
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
