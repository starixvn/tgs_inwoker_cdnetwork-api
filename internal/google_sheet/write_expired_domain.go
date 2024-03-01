package google_sheet

import (
	"cdnetwork/internal/postgresql"
	"cdnetwork/pkg/util"
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func CreateExpiredDomainExcel(domains []postgresql.DomainForExcel) {

	config := util.GetConfig()

	ctx := context.Background()

	// Read the contents of credentials.json file
	credentialsJSON, err := ioutil.ReadFile(config.GoogleApiKey)
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	// Obtain credentials from the JSON file.
	creds, err := google.CredentialsFromJSON(ctx, credentialsJSON, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to obtain credentials: %v", err)
	}

	sheetsService, err := sheets.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("Unable to create Google Sheets service: %v", err)
	}

	// ID of the Google Sheets document.
	spreadsheetId := "1NkZr5NbB-eykKAySsHfXv3OLUw2_T6zZXukCWdDUkrU"

	// Range where you want to write data.
	writeRange := "Sheet1!A1"

	// Data to be written to the sheet.
	values := [][]interface{}{
		{"Data1", "Data2", "Data3"},
	}

	// Create the value range object.
	valueRange := sheets.ValueRange{
		Values: values,
	}

	// Write data to the Google Sheets document.
	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetId, writeRange, &valueRange).
		ValueInputOption("RAW").
		Do()
	if err != nil {
		log.Fatalf("Unable to write data to sheet: %v", err)
	}

	fmt.Println("Data added successfully!")
}
