package alpha_vantage

import (
	"time"

	"github.com/JonathanNasc/toroburro/pkg/assets"
)

const SourceName string = "alpha_vantage"

type AlphaVantage struct {}

func (av *AlphaVantage) GetAsset(code string) assets.AssetPrice {
	return assets.AssetPrice{
		Code: "12",
		Price: 12.12,
		Date: time.Now(),
	}
}