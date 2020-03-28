package go-mbprice

//Price represents cryptoasset's price
type Price struct {
	High float64 `json:"high" bson:"high"`
	Low  float64 `json:"low" bson:"low"`
	Vol  float64 `json:"vol" bson:"vol"`
	Last float64 `json:"last" bson:"last"`
	Buy  float64 `json:"buy" bson:"buy"`
	Sell float64 `json:"sell" bson:"sell"`
	Open float64 `json:"open" bson:"open"`
	Date int     `json:"date" bson:"date"`
}

//MBTicker represents cryptoasset's price in MB format
type MBTicker struct {
	Ticker struct {
		High string `json:"high"`
		Low  string `json:"low"`
		Vol  string `json:"vol"`
		Last string `json:"last"`
		Buy  string `json:"buy"`
		Sell string `json:"sell"`
		Open string `json:"open"`
		Date int    `json:"date"`
	} `json:"ticker"`
}
