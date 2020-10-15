package main

import (
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg"
)

const (
	baseUrl = "https://api.coinlore.com/api/ticker/?id="
)

func main() {

	testRuleSet, err := pkg.LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		panic(err)
	}

	for _, v := range testRuleSet.Rules {
		fileUrl := baseUrl + v.CryptoID
		pkg.DownloadFile("rawData.txt", fileUrl)
	}

}
