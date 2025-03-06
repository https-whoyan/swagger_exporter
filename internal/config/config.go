package config

type CommonConfig struct {
	JsonFileName string `toml:"json_file_name" json:"json_file_name"`
	Microservice string `toml:"microservice" json:"microservice"`
}

func (c CommonConfig) Common() CommonConfig {
	return c
}

type LocalConfig struct {
	CommonConfig
	OutputFileName string `toml:"output_file_name" json:"output_file_name"`
}

func (l LocalConfig) RunMode() RunMode {
	return RunModeLocal
}

func (l LocalConfig) Config() interface{} {
	return l
}

type GoogleSheetsConfig struct {
	CommonConfig
	GoogleSheetsCredsFile string `toml:"google_sheets_creds_file" json:"google_sheets_creds_file"`
	SheetID               string `toml:"sheet_id" json:"sheet_id"`
}

func (g GoogleSheetsConfig) RunMode() RunMode {
	return RunModeGoogleSheets
}

func (g GoogleSheetsConfig) Config() interface{} {
	return g
}

type Config interface {
	RunMode() RunMode
	Common() CommonConfig
	Config() interface{}
}
