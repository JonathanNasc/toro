package main

import (
	"github.com/JonathanNasc/toroburro/pkg/aggregator"
	"github.com/JonathanNasc/toroburro/pkg/assets"
	"github.com/JonathanNasc/toroburro/pkg/settings"
)

func main() {
	s := settings.LoadSettings()
	assetsPriceList := aggregator.GetAssetsPriceList(s.Assets)
	print(assetsPriceList)
}

func print (assetPriceList []assets.AssetPrice) {
	//print head
	for _, assetPrice := range assetPriceList {
		assetPrice.ToCSV()
	}
}
