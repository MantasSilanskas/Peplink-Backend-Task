package pkg

import "fmt"

func DownloadFile()  {


	testRuleSet, err := LoadTestRuleSets("TestRulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		panic(err)
	}

	for _, v := range testRuleSet.Rules {
		fmt.Println(v.Price)
	}

	baseUrl := "https://api.coinlore.com/api/ticker/?id="
	fileUrl := baseUrl +


}