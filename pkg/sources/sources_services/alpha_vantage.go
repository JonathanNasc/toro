package alpha_vantage

import (
	"time"

	"github.com/JonathanNasc/toroburro/pkg/assets"
)

const SourceName string = "alpha_vantage"

type AlphaVantage struct {
	Name string
}

func (av *AlphaVantage) GetAsset(code string) assets.AssetResult {
	return assets.AssetResult{
		Code:  "12",
		Price: 12.12,
		Date:  time.Now(),
	}
}

func (av *AlphaVantage) GetName() string {
	return av.Name
}

func New() *AlphaVantage {
	return &AlphaVantage{
		Name: SourceName,
	}
}
