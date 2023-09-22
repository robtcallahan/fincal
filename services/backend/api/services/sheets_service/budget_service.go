package sheets_service

import (
	"errors"
	"fmt"

	"fincal/pkg/config"
	"google.golang.org/api/sheets/v4"
)

type BudgetEntry struct {
	Category            string
	WeeklyAmount        float64
	MonthlyAmount       float64
	TwiceMonthlyAmount  float64
	EveryTwoWeeksAmount float64
	YearlyAmount        float64
}

type BudgetSheet struct {
	ID            int64
	SpreadsheetID string
	TabName       string
	x             sheets.Spreadsheet
	SheetCoords   SheetCoords
	BudgetEntries []*BudgetEntry
	CategoriesMap map[string]*BudgetEntry
}

func (ss *SheetsService) NewBudgetSheet(cfg *config.Config) error {
	ss.BudgetSheet = &BudgetSheet{
		TabName: "Budget",
		SheetCoords: SheetCoords{
			StartRow:       cfg.BudgetStartRow,
			EndRow:         cfg.BudgetEndRow,
			EndColumnName:  "J",
			EndColumnIndex: 9,
		},
	}
	id, err := ss.getSheetID("Budget")
	if err != nil {
		return fmt.Errorf("unable to retrieve spreadsheet: %v", err)
	}
	ss.BudgetSheet.ID = id
	return nil
}

func (ss *SheetsService) ReadBudgetSheet() (*BudgetSheet, error) {
	range_ := fmt.Sprintf("%s!B%d:%s%d", ss.BudgetSheet.TabName, ss.BudgetSheet.SheetCoords.StartRow, ss.BudgetSheet.SheetCoords.EndColumnName, ss.BudgetSheet.SheetCoords.EndRow)
	resp, err := ss.Provider.GetValues(range_)
	if err != nil {
		return nil, fmt.Errorf("could not get sheet values: %s\n", err.Error())
	}
	if len(resp.Values) == 0 {
		return nil, errors.New("no values found")
	}

	var entries []*BudgetEntry
	var categoriesMap = make(map[string]*BudgetEntry)
	for _, values := range resp.Values {
		if ss.isEmptyBudgetRow(values) {
			continue
		}
		entry := ss.populateBudgetSheetEntry(values)
		entries = append(entries, entry)
		categoriesMap[ss.getBudgetCategory(values)] = entry
	}
	ss.BudgetSheet.BudgetEntries = entries
	ss.BudgetSheet.CategoriesMap = categoriesMap
	return ss.BudgetSheet, nil
}

func (ss *SheetsService) isEmptyBudgetRow(values []interface{}) bool {
	if len(values) < 6 || ss.getBudgetCategory(values) == "" {
		return true
	}
	return false
}

func (ss *SheetsService) getBudgetCategory(values []interface{}) string {
	return fmt.Sprintf("%s", values[0])
}

func (ss *SheetsService) populateBudgetSheetEntry(values []interface{}) *BudgetEntry {
	return &BudgetEntry{
		Category:            ss.getBudgetCategory(values),
		WeeklyAmount:        readDollarsValue(values[2]),
		MonthlyAmount:       readDollarsValue(values[3]),
		EveryTwoWeeksAmount: readDollarsValue(values[4]),
		TwiceMonthlyAmount:  readDollarsValue(values[5]),
		YearlyAmount:        readDollarsValue(values[6]),
	}
}
