package ninja

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const ninjaUrl = "http://poe.ninja/api/Data/"
const League = "HC Flashback Event (BRE002)"

var c *cache.Cache

func init() {
	c = cache.New(30*time.Minute, 32*time.Minute)
}

const (
	CategoryDivinationCards = iota
	CategoryCurrency
	CategoryFragments
	CategoryEssences
	CategoryGems
	CategoryProphecies
	CategoryMaps
	CategoryUniques
)

var Categories = map[int][]string{
	CategoryDivinationCards: {"DivinationCards"},
	CategoryCurrency:        {"Currency"},
	CategoryFragments:       {"Fragment"},
	CategoryEssences:        {"Essence"},
	CategoryGems:            {"SkillGem"},
	CategoryProphecies:      {"Prophecy"},
	CategoryMaps:            {"UniqueMap", "Map"},
	CategoryUniques:         {"UniqueJewel", "UniqueFlask", "UniqueWeapon", "UniqueArmour", "UniqueAccessory"},
}

func SearchItems(category int, itemName string) []Item {
	var items Items
	x, found := c.Get(fmt.Sprintf("%d", category))
	if found {
		items = x.(Items)
	} else {
		categories := Categories[category]
		for _, categoryItem := range categories {
			items.Items = append(items.Items, getItems(categoryItem, League).Items...)
		}
		c.Set(fmt.Sprintf("%d", category), items, cache.DefaultExpiration)
	}
	return items.Filter(itemName)
}

func getItems(category, league string) Items {
	var items Items
	league = url.PathEscape(league)
	currentDate := time.Now().Format("2006-01-02")
	uri := fmt.Sprintf("%sGet%sOverview?League=%s&date=%s", ninjaUrl, category, league, currentDate)
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return items
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &items)
	return items
}
