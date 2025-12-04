package main

import (
	"fmt"
	"os"
	"plugin"

	"git.mattmohan.com/matt/advent2025/days"
)

type DayPlugin struct {
	Filename string
	Day      days.Day
}

func getDays() []DayPlugin {
	d, err := os.ReadDir("plugins")
	if err != nil {
		panic(fmt.Errorf("Error reading plugins directory: %w", err))
	}

	days := make([]DayPlugin, 0, len(d))
	for _, entry := range d {
		if entry.IsDir() {
			continue
		}
		day := getDay("plugins/" + entry.Name())
		days = append(days, DayPlugin{
			Filename: entry.Name(),
			Day:      day,
		})
	}
	return days
}

func getDay(filename string) days.Day {
	p, err := plugin.Open(filename)
	if err != nil {
		panic(fmt.Errorf("Error loading plugin %s: %w", filename, err))
	}
	gdFunc, err := p.Lookup("GetDay")
	if err != nil {
		panic(fmt.Errorf("Error looking up GetDay in plugin %s: %w", filename, err))
	}
	getDayFunc, ok := gdFunc.(func() days.Day)
	if !ok {
		panic(fmt.Errorf("Error asserting GetDay function signature in plugin %s", filename))
	}
	return getDayFunc()
}
