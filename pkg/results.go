package pkg

import (
	"fmt"
	"strconv"
)

const (
	baseUrl               = "https://api.coinlore.com/api/ticker/?id="
	beginningDataFileName = "rawData"
	dataFileExtension     = ".json"
)

type rulesID struct {
	ID []string
}

func Parse() error {

	ruleSet, err := LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		fmt.Println(err)
	}

	var idSlice rulesID

	var currentCryptoID string
	for i, v := range ruleSet.Rules {
		if currentCryptoID == v.CryptoID {
			continue
		}
		currentCryptoID = v.CryptoID
		fileUrl := baseUrl + v.CryptoID
		DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
		idSlice.ID[i] = v.CryptoID
		i++
	}

	var lastRuleID string

	for i, v := range ruleSet.Rules {
		if lastRuleID != v.CryptoID {
			data, err := ReadDataFile(beginningDataFileName + idSlice.ID[i] + dataFileExtension)
			if err != nil {
				return err
			}
			for _, elem := range data {
				price, err := strconv.ParseFloat(elem.PriceUsd, 64)
				if err != nil {
					return err
				}
				fmt.Println(v.Price, price)
				if elem.ID == v.CryptoID {
					for i, r := range ruleSet.Rules {
						if r.Rule == "gt" && price > r.Price {
							fmt.Println("Cryptocurrency", elem.ID, elem.Name, "price is greater than", v.Price)
						}
						if r.Rule == "lt" && price < r.Price {
							fmt.Println("Cryptocurrency", elem.ID, elem.Name, "price is lower than", v.Price)
						}
						i++
						if i == 2 {
							continue
						}
					}
				} else {
					continue
				}
			}
		}
		lastRuleID = v.CryptoID
	}
	return nil
}

//for _ , _ = range ruleSet.Rules {
//	data, err := ReadDataFile("rawData"+ idSlice.ID +".json")
//	if err != nil {
//		fmt.Println(err)
//	}
//	for _, elem := range data {
//
//		price, err := strconv.ParseFloat(elem.PriceUsd, 64)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(price)
//		for i, v := range ruleSet.Rules {
//			if v.Rule == "gt" && price > v.Price {
//				fmt.Println("Cryptocurrency", elem.ID, elem.Name, "price is greater than", v.Price)
//			}
//			if v.Rule == "lt" && price < v.Price {
//				fmt.Println("Cryptocurrency", elem.ID, elem.Name, "price is lower than", v.Price)
//			}
//			i++
//			if i == 2 {
//				break
//
//			}
//		}
//	}
//}
