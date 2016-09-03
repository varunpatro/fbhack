package data

import (
	"../models"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	//"io"
	"time"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("sheets.googleapis.com-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetData() map[int]models.User {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/drive")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1Q9xcqMUXLHF57-NvCQXSv2Q_NTT7L_rVsBqRNHL9E1c"
	readRange := "Residents!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	fmt.Println(resp.Header)
	users := make(map[int]models.User)
	if len(resp.Values) > 0 {
		log.Print("Retrieved data")
		for _, row := range resp.Values {
			idstr := row[0].(string)
			id, err := strconv.Atoi(idstr)
			if err != nil {
				log.Print("some errors")
			}
			user := models.User{Name: row[1].(string), Id: id, Phone: row[2].(string)}
			users[user.Id] = user

		}
		return users
	} else {
		fmt.Print("No data found.")
		return nil
	}
}

func UpdateStatus(id int, status bool, sheetTitle string) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/drive")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	spreadsheetId := "1Q9xcqMUXLHF57-NvCQXSv2Q_NTT7L_rVsBqRNHL9E1c"
	rowToWrite := sheetTitle + "!"
	if status {
		rowToWrite += "D"
	} else {
		rowToWrite += "E"
	}
	rowToWrite += strconv.Itoa(id + 1)
	rowToWrite = strings.TrimSpace(rowToWrite)

	row := make([]interface{}, 1)
	row[0] = "OK"
	data := make([][]interface{}, 1)
	data[0] = row

	val := sheets.ValueRange{}
	val.Values = data
	resp, err := srv.Spreadsheets.Values.Update(spreadsheetId, rowToWrite, &val).ValueInputOption("RAW").Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func NewSheet() (string, int64) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/drive")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}
	spreadsheetId := "1Q9xcqMUXLHF57-NvCQXSv2Q_NTT7L_rVsBqRNHL9E1c"
	newSheetTitle := strconv.Itoa(time.Now().Day()) + time.Now().Month().String() + strconv.Itoa(time.Now().Year())
	sheetProperties := sheets.SheetProperties{}
	sheetProperties.Title = newSheetTitle
	addSheetReq := sheets.AddSheetRequest{}
	addSheetReq.Properties = &sheetProperties
	requestWrapper := sheets.Request{}
	requestWrapper.AddSheet = &addSheetReq

	req := sheets.BatchUpdateSpreadsheetRequest{}

	reqArr := make([]*sheets.Request, 1)
	reqArr[0] = &requestWrapper
	req.Requests = reqArr
	resp, err := srv.Spreadsheets.BatchUpdate(spreadsheetId, &req).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("new sheet header")
	fmt.Println(resp)
	fmt.Println()
	rowsToWrite := newSheetTitle + "!A1:E1"
	valRange := sheets.ValueRange{}
	row := make([]interface{}, 5)
	row[0] = "Id"
	row[1] = "Name"
	row[2] = "Phone Number"
	row[3] = "Mark Safe"
	row[4] = "Mark Unsafe"
	data := make([][]interface{}, 1)
	data[0] = row
	valRange.Values = data
	srv.Spreadsheets.Values.Update(spreadsheetId, rowsToWrite, &valRange).ValueInputOption("RAW").Do()

	rowsToWrite = newSheetTitle + "!A2:C"
	valRange = sheets.ValueRange{}
	valRange.Values = parseUsers(GetData())
	srv.Spreadsheets.Values.Update(spreadsheetId, rowsToWrite, &valRange).ValueInputOption("RAW").Do()


	return newSheetTitle, GetSheetId(newSheetTitle)
}

func GetSheetId(sheetTitle string) (int64){
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/drive")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}
	spreadsheetId := "1Q9xcqMUXLHF57-NvCQXSv2Q_NTT7L_rVsBqRNHL9E1c"

	resp, err := srv.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil{
		log.Fatal(err)
	}
	for _, s := range resp.Sheets {
		title := s.Properties.Title
		fmt.Println(title)
		if title == sheetTitle {
			return s.Properties.SheetId
		}
	}
	return -1
}

func parseUsers(users map[int]models.User) [][]interface{} {
	rows := make([][]interface{}, 0)
	for _, u := range users {
		row := make([]interface{}, 3)
		row[0] = u.Id
		row[1] = u.Name
		row[2] = u.Phone
		rows = append(rows, row)
	}
	return rows
}