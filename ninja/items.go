package ninja

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Items struct {
	Items []Item `json:"lines"`
}

type Item struct {
	Id                int                 `json:"id"`
	Name              string              `json:"name"`
	Icon              string              `json:"icon"`
	FlavourText       string              `json:"flavourText"`
	StackSize         int                 `json:"stackSize"`
	ChaosValue        float64             `json:"chaosValue"`
	ExaltedValue      float64             `json:"exaltedValue"`
	ExplicitModifiers []ExplicitModifiers `json:"explicitModifiers"`
	Corrupted         bool                `json:"corrupted"`
	ProphecyText      string              `json:"prophecyText"`
	GemLevel          float64             `json:"gemLevel"`
	GemQuality        float64             `json:"gemQuality"`
	MapTier           float64             `json:"mapTier"`
	Links             float64             `json:"links"`
}

func (item *Item) GetExplicitModifiersAsString() string {
	var r string
	for _, modifier := range item.ExplicitModifiers {
		if len(modifier.Text) > 0 {
			r = r + "\n" + ParseSpecialText(modifier.Text)
		}
	}
	return r
}

func getItems(uri string) Items {
	var items Items
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

func getDivinationCards(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetDivinationCardsOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getEssences(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetEssenceOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getGems(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetSkillGemOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getProphecies(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetProphecyOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getMaps(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetMapOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getUniqueMaps(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetUniqueMapOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getJewels(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetUniqueJewelOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getFlasks(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetUniqueFlaskOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getWeapons(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetUniqueWeaponOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getArmours(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetUniqueArmourOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}

func getAccessories(league string) Items {
	league = url.PathEscape(league)
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	uri := ninjaUrl + "GetUniqueAccessoryOverview?League=" + league + "&date=" + currentDate
	return getItems(uri)
}
