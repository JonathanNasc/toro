package aggregator

import (
	"github.com/JonathanNasc/toroburro/pkg/sources"
	"github.com/JonathanNasc/toroburro/pkg/assets"
	"github.com/JonathanNasc/toroburro/pkg/settings"
)

func GetAssetsPriceList(configs []settings.Asset) []assets.AssetPrice {
	var allAssetsPrice []assets.AssetPrice
	for _, config := range configs {
		assetsPriceBySource := getAssetsPriceBySource(config.Codes, config.Source)
		allAssetsPrice = append(allAssetsPrice, assetsPriceBySource...)
	}

	return allAssetsPrice
}

func getAssetsPriceBySource(assetCodes []string, sourceName string) []assets.AssetPrice {
	var assetsPriceList []assets.AssetPrice
	service := sources.Resolve(sourceName)
	
	for _, code := range assetCodes {
		asset := service.GetAsset(code)
		assetsPriceList = append(assetsPriceList, asset)
	}

	return assetsPriceList
}