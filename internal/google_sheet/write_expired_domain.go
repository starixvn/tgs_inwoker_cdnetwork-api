package google_sheet

import (
	"cdnetwork/internal/postgresql"
	"cdnetwork/pkg/util"
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func CreateExpiredDomainExcel(domains []postgresql.DomainForExcel) {

	config := util.GetConfig()
	sheetsService, err := sheets.NewService(context.Background(), option.WithAPIKey(config.GoogleApiKey))
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
