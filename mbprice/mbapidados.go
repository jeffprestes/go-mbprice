package mbprice

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const tr = &http.Transport{
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 20,
}

const httpClient = http.Client{
	Timeout:   time.Duration(30 * time.Second),
	Transport: tr,
}

const mbAPIDadosURL = "https://www.mercadobitcoin.net/api/"

//GetMBCryptoPrice gets cryptoasset price from MB
func GetMBCryptoPrice(coinTicker string) (coinPrice Price, err error) {
	err = nil
	//log.Println("GetMBCryptoPrice - Starting to query MB Service for ", coinTicker, " price...")
	resp, err := httpClient.Get(mbAPIDadosURL + coinTicker + "/ticker")
	if err != nil {
		log.Println("GetMBCryptoPrice - httpClient.Get - Erro: ", err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 250 {
		err = errors.New("MercadoBitcoin HTTP response was invalid. " + strconv.Itoa(resp.StatusCode) + " - " + resp.Status)
		log.Println("GetMBCryptoPrice - resp.StatusCode - Erro: ", err.Error())
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("GetMBCryptoPrice - ioutil.ReadAll - Erro: ", err.Error())
		return
	}
	var mbticker MBTicker
	err = json.Unmarshal(content, &mbticker)
	if err != nil {
		log.Println("GetMBCryptoPrice - json.Unmarshal - Erro: ", err.Error())
		return
	}
	coinPrice, err = ConvertMBTickerPriceFormat(mbticker)
	if err != nil {
		log.Println("GetMBCryptoPrice - ConvertMBTickerPriceFormat - Erro: ", err.Error())
		return
	}
	return
}
