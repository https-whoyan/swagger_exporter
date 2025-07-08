package executor

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func detectSwaggerMajorVersion(file *os.File) (int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return 0, fmt.Errorf("ошибка возврата указателя файла: %w", err)
	}
	var raw map[string]interface{}
	if err = json.Unmarshal(data, &raw); err != nil {
		return 0, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}
	if val, ok := raw["swagger"].(string); ok && strings.HasPrefix(val, "2.") {
		return 2, nil
	}
	if val, ok := raw["openapi"].(string); ok && strings.HasPrefix(val, "3.") {
		return 3, nil
	}
	return 0, fmt.Errorf("не удалось определить версию Swagger/OpenAPI")
}
