package ninja

import "strings"

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
	CurrencyName      string              `json:"currencyTypeName"`
	ChaosPrice        float64             `json:"chaosEquivalent"`
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

func (items *Items) Filter(itemName string) []Item {
	var result []Item
	for _, item := range items.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(itemName)) ||
			strings.Contains(strings.ToLower(item.CurrencyName), strings.ToLower(itemName)) {
			result = append(result, item)
		}
	}
	return result
}
