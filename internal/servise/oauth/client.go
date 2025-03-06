package oauth

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io"
	"log"
	"net/http"
	"os"
)

func GetHTTPCli(ctx context.Context, creedsFile *os.File) (*http.Client, error) {
	jsonData, err := io.ReadAll(creedsFile)
	if err != nil {
		return nil, err
	}
	config, err := google.JWTConfigFromJSON(jsonData, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatal("Ошибка аутентификации:", err)
	}
	client := config.Client(context.Background())
	return client, nil
}
