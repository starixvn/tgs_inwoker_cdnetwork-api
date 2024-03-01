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
	creds, err := ioutil.ReadFile(config.GoogleApiKey)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to read client secret file: %v", err))
	}

	conf, err := google.JWTConfigFromJSON(creds, sheets.SpreadsheetsScope)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
	}

	// Create a new JWT config for authentication.
	jwtConf := &jwt.Config{
		Email:      conf.Email,
		PrivateKey: conf.PrivateKey,
		Scopes:     []string{sheets.SpreadsheetsScope},
		TokenURL:   conf.TokenURL,
	}

	// Create a new Google Sheets client.
	client := jwtConf.Client(context.Background())

	// Specify the spreadsheet ID and range.
	spreadsheetId := "1NkZr5NbB-eykKAySsHfXv3OLUw2_T6zZXukCWdDUkrU"

	///////////////////////////////////////////////

	// Initialize the Sheets service.
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.LogFatal(fmt.Sprint("Unable to retrieve Sheets client: %v", err))
	}

	nextMonth := time.Now().AddDate(0, 1, 0)
	nextMonthYear := nextMonth.Format("01/2006") // Format: MM/YYYY

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
	_, err = sheetsService.Spreadsheets.BatchUpdate(spreadsheetId, batchUpdateRequest).Do()
	if err != nil {
		log.LogFatal(fmt.Sprint("Unable to create new sheet: %v", err))
	}

	fmt.Printf("New sheet '%s' created successfully.\n", nextMonthYear)

	////////////////////////////////////////////////

	// Create the value range to write.
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

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	writeRange := fmt.Sprintf("%s!A1:E%d", nextMonthYear, len(values))

	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).
		ValueInputOption("RAW").Do()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to write data to sheet: %v", err))
	}

	fmt.Println("Data successfully written to Google Sheet.")
}
