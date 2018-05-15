package ninja

import (
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

const ninjaUrl = "http://poe.ninja/api/Data/"
const League = "HC Flashback Event (BRE002)"

var c *cache.Cache

func init() {
	c = cache.New(30*time.Minute, 32*time.Minute)
}

func GetDivinationCard(cardName string) []Item {
	var result []Item
	var items Items
	x, found := c.Get("cards")
	if found {
		items = x.(Items)
	} else {
		items = getDivinationCards(League)
		c.Set("cards", items, cache.DefaultExpiration)
	}
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(cardName)) {
			result = append(result, item)
		}
	}
	return result
}

func GetCurrency(currencyName string) ([]Currency, []CurrencyDetails) {
	var currency []Currency
	var currencies Currencies
	x, found := c.Get("currency")
	if found {
		currencies = x.(Currencies)
	} else {
		currencies = getCurrency(League)
		c.Set("currency", currencies, cache.DefaultExpiration)
	}
	for _, curr := range currencies.Lines {
		if strings.Contains(strings.ToLower(curr.CurrencyName), strings.ToLower(currencyName)) {
			currency = append(currency, curr)
		}
	}
	return currency, currencies.CurrencyDetails
}

func GetFragments(fragmentName string) ([]Currency, []CurrencyDetails) {
	var fragment []Currency
	var fragments Currencies
	x, found := c.Get("fragments")
	if found {
		fragments = x.(Currencies)
	} else {
		fragments = getFragments(League)
		c.Set("fragments", fragments, cache.DefaultExpiration)
	}
	for _, frag := range fragments.Lines {
		if strings.Contains(strings.ToLower(frag.CurrencyName), strings.ToLower(fragmentName)) {
			fragment = append(fragment, frag)
		}
	}
	return fragment, fragments.CurrencyDetails
}

func GetEssences(essenceName string) []Item {
	var result []Item
	var items Items
	x, found := c.Get("essences")
	if found {
		items = x.(Items)
	} else {
		items = getEssences(League)
		c.Set("essences", items, cache.DefaultExpiration)
	}
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(essenceName)) {
			result = append(result, item)
		}
	}
	return result
}

func GetGems(gemName string) []Item {
	var result []Item
	var items Items
	x, found := c.Get("gems")
	if found {
		items = x.(Items)
	} else {
		items = getGems(League)
		c.Set("gems", items, cache.DefaultExpiration)
	}
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(gemName)) {
			result = append(result, item)
		}
	}
	return result
}

func GetProphecies(prophecyName string) []Item {
	var result []Item
	var items Items
	x, found := c.Get("prophecies")
	if found {
		items = x.(Items)
	} else {
		items = getProphecies(League)
		c.Set("prophecies", items, cache.DefaultExpiration)
	}
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(prophecyName)) {
			result = append(result, item)
		}
	}
	return result
}

func GetMaps(itemName string) []Item {
	var result []Item
	var items Items
	x, found := c.Get("maps")
	if found {
		items = x.(Items)
	} else {
		items = getMaps(League)
		t := &items
		t.Items = append(t.Items, getUniqueMaps(League).Items...)
		c.Set("maps", items, cache.DefaultExpiration)
	}
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(itemName)) {
			result = append(result, item)
		}
	}
	return result
}

func GetUniques(itemName string) []Item {
	var result []Item
	var items Items
	x, found := c.Get("uniques")
	if found {
		items = x.(Items)
	} else {
		items = getJewels(League)
		t := &items
		t.Items = append(t.Items, getFlasks(League).Items...)
		t.Items = append(t.Items, getWeapons(League).Items...)
		t.Items = append(t.Items, getArmours(League).Items...)
		t.Items = append(t.Items, getAccessories(League).Items...)
		c.Set("uniques", items, cache.DefaultExpiration)
	}
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(itemName)) {
			result = append(result, item)
		}
	}
	return result
}
