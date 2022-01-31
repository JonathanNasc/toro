package assets

import (
	"fmt"
	"strconv"
	"time"
)

type AssetResult struct {
	Code   string
	Price  float32
	Date   time.Time
}

func (a AssetResult) ToCSV() string {
	result := strconv.FormatFloat(float64(a.Price), 'e', -1, 64)
	return fmt.Sprintf("%s,%s,%s", a.Code, result, a.Date.Format(time.RFC3339))
}
