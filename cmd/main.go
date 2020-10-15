package main

import (
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg"
)

const (
	baseUrl = "https://api.coinlore.com/api/ticker/?id="
)

func main() {

	ruleSet, err := pkg.LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		panic(err)
	}

	var currentCryptoID string
	for _, v := range ruleSet.Rules {
		if currentCryptoID == v.CryptoID {
			continue
		}
		currentCryptoID = v.CryptoID
		fileUrl := baseUrl + v.CryptoID
		pkg.DownloadFile("rawData"+v.CryptoID+".json", fileUrl)
	}

	pkg.ReadDataFile("rawData90.json")

}
