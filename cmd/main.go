package main

import (
	"encoding/json"
	"fmt"
	"os"
	""
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

