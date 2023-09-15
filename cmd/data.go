package cmd

import (
	"fmt"

	"vue-register/api/providers/sheets_provider"
	"vue-register/api/services/sheets_service"
	cfg "vue-register/pkg/config"
	"vue-register/pkg/driver"
	"vue-register/pkg/handler"
	"vue-register/pkg/models"

	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Reads various spreadsheet pages and inerts into DB",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		data()
	},
}

/**/
func init() {
	config, _ = cfg.ReadConfig(ConfigFile)
	rootCmd.AddCommand(dataCmd)
}

func data() {
	var (
		err error
	)

	conn, err := driver.ConnectSQL(&driver.ConnectParams{
		DBType: driver.DBType(config.DBType),
		Host:   config.DBHost,
		Port:   config.DBPort,
		DBName: config.DBName,
		User:   config.DBUsername,
		Pass:   config.DBPassword,
	})
	checkError(err)
	qHandler := handler.NewQueryHandler(conn)

	sheetsProvider, err := sheets_provider.New(options.SpreadsheetID, config)
	checkError(err)
	sheetsService := sheets_service.New(sheetsProvider)
	checkError(err)
	err = sheetsService.NewRegisterSheet(config)
	checkError(err)

	categories := qHandler.GetCategories()
	name2category := map[string]models.Category{}
	for _, c := range categories {
		name2category[c.Name] = c
	}

	fmt.Println("Reading Budget...")
	err = sheetsService.NewBudgetSheet(config)
	checkError(err)
	budgetSheet, err := sheetsService.ReadBudgetSheet()
	checkError(err)

	for _, be := range budgetSheet.BudgetEntries {
		if be.Category == "" {
			continue
		}
		category := name2category[be.Category]

		qHandler.UpdatesCategory(&category, models.Category{
			WeeklyAmount:        be.WeeklyAmount,
			MonthlyAmount:       be.MonthlyAmount,
			EveryTwoWeeksAmount: be.EveryTwoWeeksAmount,
			TwiceMonthlyAmount:  be.TwiceMonthlyAmount,
			YearlyAmount:        be.YearlyAmount,
		})
	}
}
