package main

import (
	"fmt"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg"
)

func main() {

	testRuleSet, err := LoadRuleSets("TestRulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		panic(err)
	}

	for _, v := range testRuleSet.Rules {
		fmt.Println(v.Price)
	}

}
