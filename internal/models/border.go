package models

import "google.golang.org/api/sheets/v4"

type Border = sheets.Border
type Borders struct {
	Border *Border `json:"border,omitempty"`
}
