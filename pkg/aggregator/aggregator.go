package aggregator

import (
	"github.com/JonathanNasc/toroburro/pkg/assets"
	"github.com/JonathanNasc/toroburro/pkg/settings"
	"github.com/JonathanNasc/toroburro/pkg/sources"
)

func GetAssetsResultList(configs []settings.Asset) []assets.AssetResult {
	var allAssetsResult []assets.AssetResult
	for _, config := range configs {
		assetsResultBySource := getAssetsResultBySource(config.Codes, config.Source)
		allAssetsResult = append(allAssetsResult, assetsResultBySource...)
	}

	return allAssetsResult
}

func getAssetsResultBySource(assetCodes []string, sourceName string) []assets.AssetResult {
	var assetsResultList []assets.AssetResult
	service := sources.Resolve(sourceName)

	for _, code := range assetCodes {
		asset := service.GetAsset(code)
		assetsResultList = append(assetsResultList, asset)
	}

	return assetsResultList
}
