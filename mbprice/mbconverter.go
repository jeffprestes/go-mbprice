package mbprice

import "strconv"

//ConvertMBTickerPriceFormat converts JSON with MB Ticker's info to Price model
func ConvertMBTickerPriceFormat(ticker MBTicker) (price Price, err error) {
	price.High, err = strconv.ParseFloat(ticker.Ticker.High, 64)
	if err != nil {
		return
	}
	price.Low, err = strconv.ParseFloat(ticker.Ticker.Low, 64)
	if err != nil {
		return
	}
	price.Vol, err = strconv.ParseFloat(ticker.Ticker.Vol, 64)
	if err != nil {
		return
	}
	price.Last, err = strconv.ParseFloat(ticker.Ticker.Last, 64)
	if err != nil {
		return
	}
	price.Buy, err = strconv.ParseFloat(ticker.Ticker.Buy, 64)
	if err != nil {
		return
	}
	price.Sell, err = strconv.ParseFloat(ticker.Ticker.Sell, 64)
	if err != nil {
		return
	}
	price.Open, err = strconv.ParseFloat(ticker.Ticker.Open, 64)
	if err != nil {
		return
	}
	price.Date = ticker.Ticker.Date
	return
}
