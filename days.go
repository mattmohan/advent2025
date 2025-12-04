package main

import (
	"git.mattmohan.com/matt/advent2025/day4"
	"git.mattmohan.com/matt/advent2025/days"

	"git.mattmohan.com/matt/advent2025/day1"
	"git.mattmohan.com/matt/advent2025/day2"
	"git.mattmohan.com/matt/advent2025/day3"
)

func getDays() []days.Day {
	return []days.Day{
		day1.GetDay(),
		day2.GetDay(),
		day3.GetDay(),
		day4.GetDay(),
	}
}
