package peplink

import (
	"fmt"
	"strconv"
)

const (
	baseUrl               = "https://api.coinlore.com/api/ticker/?id="
	beginningDataFileName = "rawData"
	dataFileExtension     = ".json"
)

// funkcijos pavadinimas neatitinka realybes. Funkcija atlieka ne tik parse funkcionaluma, bet
// viduje issitraukia per api informacija, ir ja sulygina su rules.
// reiketu issiskaidyti funkcijas, kad butu galima implementuot kazkaip panasiai:
/*
	rules, err := readRules(path)
	if err != nil { .. }

	data, err := extractFromAPI(rules)
	if err != nil { .. }

	triggers := comparePrices(rules, data)
	for .. range triggers {
		print triggered rule
	}
*/
// isskaidymas i tokais funkcijas leistu lengviau implementuot testus. Siuo atveju mus ypac
// domintu comparePrices funkcijos testai.
func Parse(resultMap map[int]bool, rulesPrices map[int]float64) (*map[int]bool, *map[int]float64, error) {

	ruleSet, err := LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		return &resultMap, &rulesPrices, err
	}

	dataMap := make(map[string]CryptoCurrencyData)

	for i, v := range ruleSet.Rules {

		fileUrl := baseUrl + v.CryptoID

		// FIXME: visa sita if { .. } reiketu ismest i atskira funkcija
		if mapData, ok := dataMap[v.CryptoID]; !ok {

			// XXX: asmeniskai man nelabai patinka sitas sprendimas, jog rezultatai
			// pirma yra irasomi i faila, o po to, nuskaitomas failo turinys ir taip
			// atliekami palyginimai. Debuginimui gal ir smagu tureti failus, bet
			// rezultatus debuginimo tikslais galima printinti i stdout. O siaip,
			// siuo atveju reviewinant tiesiog prideda daugiau nereikalingo complexity
			// kuri reikia suprasti.
			err := DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
			if err != nil {
				return &resultMap, &rulesPrices, err
			}
			mapData, err = ReadDataFile(beginningDataFileName + v.CryptoID + dataFileExtension)
			if err != nil {
				return &resultMap, &rulesPrices, err
			}
			dataMap[v.CryptoID] = mapData
		}

		// nelabai suprantu kam reikalinga sita vieta
		if rulesPrices[i] != v.Price {
			resultMap[i] = false
		}

		price, err := strconv.ParseFloat(dataMap[v.CryptoID].PriceUsd, 64)
		if err != nil {
			return &resultMap, &rulesPrices, err
		}
		if ok := resultMap[i]; !ok {
			if price > v.Price && v.Rule == "gt" && resultMap[i] == false {
				fmt.Println("Cryptocurrency", " id:", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is greater than", v.Price)
				resultMap[i] = true
			}
			if price < v.Price && v.Rule == "lt" && resultMap[i] == false {
				fmt.Println("Cryptocurrency", " id:", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is lower than", v.Price)
				resultMap[i] = true
			}
		}
		rulesPrices[i] = v.Price

		// tu jau iteruoji su `for i, v := range`, todel i++ nebereikia
		// cia atlikti paciam, range padaro tai uz tave. Atlieki i++ bet jo reiksme
		// bus numusta kiekvienoj ciklo iteracijoj.
		i++
	}

	return &resultMap, &rulesPrices, nil
}
