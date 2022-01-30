package assets

import (
	"fmt"
	"strconv"
	"time"
)

type AssetPrice struct {
	Code  string
	Price float32
	Date  time.Time
}

func (a AssetPrice) ToCSV() string {
	price := strconv.FormatFloat(float64(a.Price), 'e', -1, 64)
	return fmt.Sprintf("%s,%s,%s", a.Code, price, a.Date.Format(time.RFC3339))
}
