package ninja

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Currencies struct {
	Lines           []Currency        `json:"lines"`
	CurrencyDetails []CurrencyDetails `json:"currencyDetails"`
}
type Currency struct {
	CurrencyName string  `json:"currencyTypeName"`
	ChaosPrice   float64 `json:"chaosEquivalent"`
}

type CurrencyDetails struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
}

func (c *Currency) Price() string {
	return strconv.FormatFloat(c.ChaosPrice, 'f', -1, 64)
}

func (c *Currency) Name() string {
	return c.CurrencyName
}

func (c *Currency) Icon(details []CurrencyDetails) string {
	var icon string
	for _, detail := range details {
		if detail.Name == c.CurrencyName {
			icon = detail.Icon
		}
	}
	return icon
}

func getCurrency(league string) Currencies {
	league = url.PathEscape(league)
	var currencies Currencies
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetCurrencyOverview?League=" + league + "&date=" + currentDate
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return currencies
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &currencies)
	return currencies
}

func getFragments(league string) Currencies {
	league = url.PathEscape(league)
	var fragments Currencies
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetFragmentOverview?League=" + league + "&date=" + currentDate
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return fragments
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &fragments)
	return fragments
}
