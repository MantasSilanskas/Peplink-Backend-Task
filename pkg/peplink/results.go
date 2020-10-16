package peplink

type Results struct {
	RuleType     string
	cryptoID     string
	cryptoName   string
	CryptoPrice  float64
	NeedPrinting bool
	Printed      bool
}

type RawResults struct {
	RuleType     string
	cryptoID     string
	cryptoName   string
	CryptoPrice  float64
	NeedPrinting bool
	Printed      bool
}
