package sheets

import (
	"context"
	"google.golang.org/api/sheets/v4"
)

func (c *sheetsClient) doRequest(
	ctx context.Context, req *sheets.BatchUpdateSpreadsheetRequest,
) error {
	_, err := c.sheetsS.Spreadsheets.BatchUpdate(c.sheetID, req).Context(ctx).Do()
	return err
}

func parsePosition(pos string) (int, int) {
	col := 0
	row := 0

	for _, r := range pos {
		if r >= 'A' && r <= 'Z' {
			col = col*26 + int(r-'A') + 1
		} else if r >= '0' && r <= '9' {
			row = row*10 + int(r-'0')
		} else {
			break
		}
	}

	return row - 1, col - 1
}

func columnLetterToIndex(col string) int64 {
	index := int64(0)
	for i := 0; i < len(col); i++ {
		index = index*26 + int64(col[i]-'A'+1)
	}
	return index - 1
}
