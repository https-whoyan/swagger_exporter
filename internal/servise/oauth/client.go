package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io"
	"net/http"
	"os"
)

func GetHTTPCli(_ context.Context, creedsFile *os.File) (*http.Client, error) {
	jsonData, err := io.ReadAll(creedsFile)
	if err != nil {
		return nil, err
	}
	config, err := google.JWTConfigFromJSON(jsonData, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, fmt.Errorf("ошибка аутентификации: %v", err)
	}
	client := config.Client(context.Background())
	return client, nil
}
