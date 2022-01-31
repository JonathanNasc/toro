package main

import (
	"fmt"

	"github.com/JonathanNasc/toroburro/pkg/aggregator"
	"github.com/JonathanNasc/toroburro/pkg/assets"
	"github.com/JonathanNasc/toroburro/pkg/settings"
)

func main() {
	s := settings.LoadSettings()
	assetsResultList := aggregator.GetAssetsResultList(s.Assets)
	printCSV(assetsResultList)
}

func printCSV(assetResultList []assets.AssetResult) {
	header := "asset,price,date"
	fmt.Println(header)

	for _, assetResult := range assetResultList {
		fmt.Println(assetResult.ToCSV())
	}
}
