package config

type RunMode string

const (
	RunModeLocal        RunMode = "local"
	RunModeGoogleSheets RunMode = "google_sheets"
)

func (m RunMode) String() string {
	return string(m)
}
