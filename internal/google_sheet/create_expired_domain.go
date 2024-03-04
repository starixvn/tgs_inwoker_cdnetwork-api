package google_sheet

import (
	"bytes"
	"cdnetwork/internal/postgresql"
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"context"
	"fmt"
	"html/template"
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

	msg := createMessage(domains)
	writeMessage(msg, len(domains), nextMonthYear, sheetsService, spreadsheetId)

	fmt.Println("Data successfully written to Google Sheet.")
}

func writeMessage(composeMessage string, len int, sheet string, sheetsService *sheets.Service, spreadsheetId string) {
	rangeValue := fmt.Sprintf("%s!A%d:E%d", sheet, len+3, len+3)

	// Prepare the value range with the message.
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{
			{composeMessage}, // Call the function to compose the message
		},
	}

	// Call the Sheets API to update the value in the specified range.
	_, err := sheetsService.Spreadsheets.Values.Update(spreadsheetId, rangeValue, valueRange).
		ValueInputOption("RAW").Do()
	if err != nil {
		log.LogFatal(fmt.Sprintf("Unable to update cell value: %v", err))
	}
}
func createMessage(domains []postgresql.DomainForExcel) string {
	expirationDate := getResponseDate() + "(五)"

	domainNames := mapStrings(domains, func(domain postgresql.DomainForExcel) string {
		return domain.Domain
	})

	// Prepare data to be passed to the template
	data := struct {
		ExpirationDate string
		Domains        []string
	}{
		ExpirationDate: expirationDate,
		Domains:        domainNames,
	}

	const msgTemplate = `
貴司您好，以下{{ len .Domains }}個域名即將到期，綁定代理、到期日、續約價格請參考上圖；
請於{{ .ExpirationDate }}之前回覆是否需要續約，我司將優先為您處理，謝謝您!
{{- range .Domains }}
{{ . }}
{{- end }}
	`

	tmpl, err := template.New("message").Parse(msgTemplate)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Error parsing message template: %v", err))
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.LogFatal(fmt.Sprintf("Error executing template: %v", err))
	}

	return buf.String()
}

func mapStrings(domains []postgresql.DomainForExcel, f func(postgresql.DomainForExcel) string) []string {
	result := make([]string, len(domains))
	for i, v := range domains {
		result[i] = f(v)
	}
	return result
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

	newSheetProperties := &sheets.SheetProperties{
		Title: nextMonthYear,
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: newSheetProperties,
				},
			},
		},
	}

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

	for _, sheet := range resp.Sheets {
		if sheet.Properties.Title == sheetName {
			return true
		}
	}
	return false
}

func createValueRange(domains []postgresql.DomainForExcel) (*sheets.ValueRange, int) {
	values := [][]interface{}{
		{"Brand", "Agent", "Domain", "Created", "Expired"},
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

func getLastFridayOfMonth(now time.Time) time.Time {
	// Get the first day of the next month
	firstDayOfNextMonth := now.AddDate(0, 1, 0)
	firstDayOfNextMonth = time.Date(firstDayOfNextMonth.Year(), firstDayOfNextMonth.Month(), 1, 0, 0, 0, 0, time.Local)

	// Subtract one day to get the last day of the current month
	lastDayOfCurrentMonth := firstDayOfNextMonth.Add(-24 * time.Hour)

	// Determine the day of the week for the last day of the month
	dayOfWeek := lastDayOfCurrentMonth.Weekday()

	// Calculate the number of days to subtract to get to the last Friday
	daysToSubtract := int(dayOfWeek - time.Friday)
	if daysToSubtract < 0 {
		daysToSubtract += 7 // Ensure positive result
	}

	// Subtract the appropriate number of days to find the last Friday
	return lastDayOfCurrentMonth.Add(-time.Duration(daysToSubtract) * 24 * time.Hour)
}

func getResponseDate() string {
	lastFridayOfMonth := getLastFridayOfMonth(time.Now())

	formattedLastFriday := lastFridayOfMonth.Format("01/02")

	fmt.Println("Last Friday of the month (formatted as {month/day}):", formattedLastFriday)
	return formattedLastFriday
}
