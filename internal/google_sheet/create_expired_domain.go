package google_sheet

import (
	"cdnetwork/internal/postgresql"
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

func CreateExpiredDomainExcel(domains []postgresql.DomainForExcel) {
	config := util.GetConfig()
	googleApiKey := config.GoogleSheet.GoogleApiKey
	spreadsheetId := config.GoogleSheet.SheetId

	sheetsService := createSheetsService(googleApiKey)
	nextMonthYear := createSheet(sheetsService, spreadsheetId)
	valueRange, valuesCount := createValueRange(domains)
	writeRange := fmt.Sprintf("%s!A1:E%d", nextMonthYear, valuesCount)

	_, err := sheetsService.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).
		ValueInputOption("RAW").Do()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to write data to sheet: %v", err))
	}

	fmt.Println("Data successfully written to Google Sheet.")
}

func createSheetsService(key string) *sheets.Service {
	creds, err := ioutil.ReadFile(key)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to read client secret file: %v", err))
	}

	conf, err := google.JWTConfigFromJSON(creds, sheets.SpreadsheetsScope)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
	}

	jwtConf := &jwt.Config{
		Email:      conf.Email,
		PrivateKey: conf.PrivateKey,
		Scopes:     []string{sheets.SpreadsheetsScope},
		TokenURL:   conf.TokenURL,
	}

	client := jwtConf.Client(context.Background())

	sheetsService, err := sheets.New(client)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to retrieve Sheets client: %v", err))
	}
	return sheetsService
}

func createSheet(sheetsService *sheets.Service, spreadsheetId string) string {
	nextMonth := time.Now().AddDate(0, 1, 0)
	nextMonthYear := nextMonth.Format("01/2006") // Format: MM/YYYY

	if checkIfSheetExists(sheetsService, spreadsheetId, nextMonthYear) {
		return nextMonthYear
	}
	// Specify the properties of the new sheet.
	newSheetProperties := &sheets.SheetProperties{
		Title: nextMonthYear, // Use the next month/year as the sheet title
	}

	// Create a batch update request to add a new sheet.
	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: newSheetProperties,
				},
			},
		},
	}

	// Execute the batch update request to create a new sheet.
	_, err := sheetsService.Spreadsheets.BatchUpdate(spreadsheetId, batchUpdateRequest).Do()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to create new sheet: %v", err))
	}

	fmt.Printf("New sheet '%s' created successfully.\n", nextMonthYear)
	return nextMonthYear
}

func checkIfSheetExists(sheetsService *sheets.Service, spreadsheetId string, sheetName string) bool {
	resp, err := sheetsService.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to retrieve spreadsheet: %v", err))
	}

	// Iterate through the sheets to check if the specified sheet exists.
	sheetExists := false
	for _, sheet := range resp.Sheets {
		if sheet.Properties.Title == sheetName {
			sheetExists = true
			break
		}
	}
	return sheetExists
}

func createValueRange(domains []postgresql.DomainForExcel) (*sheets.ValueRange, int) {
	values := [][]interface{}{
		{"Brand", "Agent", "Domain", "Created", "Expired"}, // Example data
	}

	for _, domain := range domains {
		row := []interface{}{
			domain.Brand,
			domain.Agent,
			domain.Domain,
			domain.Created,
			domain.Expires,
		}
		values = append(values, row)
	}

	return &sheets.ValueRange{
		Values: values,
	}, len(values)
}
