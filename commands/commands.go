package commands

import (
	"fmt"
	"github.com/Mi7teR/shavronne/ninja"
	"strings"
)

func Process(message string) (response []string) {
	if strings.HasPrefix(message, "!h") {
		response = help()
		return
	}
	if len(message) < 7 {
		var message string
		message = fmt.Sprintln("**Error**")
		message = message + fmt.Sprintln("Request length should be at least 4 symbols")
		response = append(response, message)
		return
	}

	commands := map[string]func(message string) (response []string){
		"!d ": divinationCards,
		"!c ": currency,
		"!с ": currency,
		"!f ": fragments,
		"!e ": essences,
		"!g ": gems,
		"!p ": prophecies,
		"!m ": maps,
		"!u ": uniques,
	}

	for command, f := range commands {
		if strings.HasPrefix(message, command) {
			response = createSearchRequest(message, command, f)
		}
	}
	return
}

func help() (response []string) {
	var message string
	message = fmt.Sprintln("__**Example commands:**__")
	message += fmt.Sprintln("**Divination cards prices search:**")
	message += fmt.Sprintln("```!d humi```")
	message += fmt.Sprintln("**Currency prices search:**")
	message += fmt.Sprintln("```!c exalted```")
	message += fmt.Sprintln("**Fragments prices search:**")
	message += fmt.Sprintln("```!f hope```")
	message += fmt.Sprintln("**Essences prices search:**")
	message += fmt.Sprintln("```!e horror```")
	message += fmt.Sprintln("**Gems prices search:**")
	message += fmt.Sprintln("```!g spark```")
	message += fmt.Sprintln("**Prophecies prices search:**")
	message += fmt.Sprintln("```!p wind```")
	message += fmt.Sprintln("**Maps prices search:**")
	message += fmt.Sprintln("```!m manor```")
	message += fmt.Sprintln("**Uniques prices search:**")
	message += fmt.Sprintln("```!u voltaxic```")
	message += fmt.Sprintln("_Bot using poe.ninja API, thx to poe.ninja owner_")
	response = append(response, message)
	return
}

func createSearchRequest(message, command string, fn func(string) []string) []string {
	message = strings.TrimSpace(strings.TrimLeft(message, command))
	return fn(message)
}

func notFoundError() string {
	return "Search results not found"
}

func divinationCards(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryDivinationCards, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.Name)
			itemMessage += fmt.Sprintf("_%s_\n", ninja.ParseSpecialText(item.FlavourText))
			itemMessage += fmt.Sprintln("**Result:**")
			itemMessage += fmt.Sprintf("_%s_\n", item.GetExplicitModifiersAsString())
			itemMessage += fmt.Sprintf("**Stack size:** ```%d```\n", item.StackSize)
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosValue)
			itemMessage += fmt.Sprintf("**Price in Exalted Orbs:** ```%.2f```\n", item.ExaltedValue)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func currency(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryCurrency, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.CurrencyName)
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosPrice)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func fragments(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryFragments, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.CurrencyName)
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosPrice)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func essences(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryEssences, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.Name)
			itemMessage += fmt.Sprintf("_%s_\n", item.GetExplicitModifiersAsString())
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosValue)
			itemMessage += fmt.Sprintf("**Price in Exalted Orbs:** ```%.2f```\n", item.ExaltedValue)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func gems(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryGems, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.Name)
			itemMessage += fmt.Sprintf("_%s_\n", item.GetExplicitModifiersAsString())
			itemMessage += fmt.Sprintf("**Level:** ```%.2f```\n", item.GemLevel)
			itemMessage += fmt.Sprintf("**Quality:** ```%.2f```\n", item.GemQuality)
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosValue)
			itemMessage += fmt.Sprintf("**Price in Exalted Orbs:** ```%.2f```\n", item.ExaltedValue)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func prophecies(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryProphecies, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.Name)
			itemMessage += fmt.Sprintf("_%s_\n", ninja.ParseSpecialText(item.FlavourText))
			itemMessage += fmt.Sprintf("**Prophecy:**\n ```%s```\n", ninja.ParseSpecialText(item.ProphecyText))
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosValue)
			itemMessage += fmt.Sprintf("**Price in Exalted Orbs:** ```%.2f```\n", item.ExaltedValue)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func maps(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryMaps, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.Name)
			if len(item.FlavourText) > 0 {
				itemMessage += fmt.Sprintf("_%s_\n", ninja.ParseSpecialText(item.FlavourText))
			}
			if len(item.ExplicitModifiers) > 0 {
				itemMessage += fmt.Sprintf("**Modifiers:** \n```%s```\n", item.GetExplicitModifiersAsString())
			}
			itemMessage += fmt.Sprintf("**Tier:** ```%.2f```\n", item.MapTier)
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosValue)
			itemMessage += fmt.Sprintf("**Price in Exalted Orbs:** ```%.2f```\n", item.ExaltedValue)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}

func uniques(message string) (response []string) {
	items := ninja.SearchItems(ninja.CategoryUniques, message)
	if len(items) > 0 {
		for _, item := range items {
			var itemMessage string
			itemMessage += fmt.Sprintf("__**%s:**__\n", item.Name)
			if len(item.FlavourText) > 0 {
				itemMessage += fmt.Sprintf("_%s_\n", ninja.ParseSpecialText(item.FlavourText))
			}
			if len(item.ExplicitModifiers) > 0 {
				itemMessage += fmt.Sprintf("**Modifiers:** \n```%s```\n", item.GetExplicitModifiersAsString())
			}
			if item.Links > 0 {
				itemMessage += fmt.Sprintf("**Links:** ```%.2f```\n", item.Links)
			}
			itemMessage += fmt.Sprintf("**Price in Chaos Orbs:** ```%.2f```\n", item.ChaosValue)
			itemMessage += fmt.Sprintf("**Price in Exalted Orbs:** ```%.2f```\n", item.ExaltedValue)
			response = append(response, itemMessage)
		}
	} else {
		response = append(response, notFoundError())
	}
	return
}
