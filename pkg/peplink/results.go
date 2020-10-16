package peplink

type Result struct {
	ID           int
	RuleType     string
	cryptoID     string
	cryptoName   string
	CryptoPrice  float64
	NeedPrinting bool
	Printed      bool
}
