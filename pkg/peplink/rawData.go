package peplink

type CryptoCurrencyData struct {
	ID               string `json:"id"`
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Nameid           string `json:"nameid"`
	Rank             int    `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange7D  string `json:"percent_change_7d"`
	MarketCapUsd     string `json:"market_cap_usd"`
	Volume24         string `json:"volume24"`
	Volume24Native   string `json:"volume24_native"`
	Csupply          string `json:"csupply"`
	PriceBtc         string `json:"price_btc"`
	Tsupply          string `json:"tsupply"`
	Msupply          string `json:"msupply"`
}

type CryptoCurrencyDataSlice []CryptoCurrencyData
