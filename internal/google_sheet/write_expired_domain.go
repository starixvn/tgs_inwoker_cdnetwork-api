package google_sheet

import (
	"cdnetwork/internal/postgresql"
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"context"
	"fmt"
	"io/ioutil"

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
	writeRange := "Sheet1!A1"

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

	// Call the Sheets API to write data.
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to retrieve Sheets client: %v", err))
	}

	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).
		ValueInputOption("RAW").Do()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to write data to sheet: %v", err))
	}

	fmt.Println("Data successfully written to Google Sheet.")
}
