package calendar

import (
	"time"

	"github.com/willbelucky/calendar/city"
)

type Calendar struct {
	holidays              map[string][]string
	BusinessDays          []BusinessDay
	BusinessDayConvention BusinessDayConvention
}

func GetCalendar(businessDays []BusinessDay, businessDayConvention BusinessDayConvention) (c Calendar) {
	c = Calendar{
		holidays:              map[string][]string{},
		BusinessDays:          businessDays,
		BusinessDayConvention: businessDayConvention,
	}
	for _, businessDay := range businessDays {
		switch businessDay {
		case SOUTH_KOREA_SEOUL:
			for dateString, name := range city.SouthKoreaSeoulHolidays {
				c.addHoliday(dateString, name)
			}
		default:
			panic("Invalid BusinessDay")
		}
	}
	return c
}

type BusinessDayConvention string

const (
	UNADJUSTED         BusinessDayConvention = "UNADJUSTED"
	FOLLOWING          BusinessDayConvention = "FOLLOWING"
	MODIFIED_FOLLOWING BusinessDayConvention = "MODIFIED_FOLLOWING"
	PRECEDING          BusinessDayConvention = "PRECEDING"
	MODIFIED_PRECEDING BusinessDayConvention = "MODIFIED_PRECEDING"
)

type BusinessDay string

const (
	SOUTH_KOREA_SEOUL BusinessDay = "SOUTH_KOREA_SEOUL"
)

func (c *Calendar) IsBusinessDay(date Date) bool {
	// If the date is a Saturday or Sunday, return false
	if isWeekend(date) {
		return false
	}

	// If the date exists in the holidays map, return false
	if _, ok := c.holidays[date.String()]; ok {
		return false
	}

	// Otherwise, return true
	return true
}

func isWeekend(date Date) bool {
	switch date.Weekday() {
	case time.Saturday:
		return true
	case time.Sunday:
		return true
	}
	return false
}

func (c *Calendar) Adjust(date Date) Date {
	switch c.BusinessDayConvention {
	case UNADJUSTED:
		return date
	case FOLLOWING:
		return c.following(date)
	case MODIFIED_FOLLOWING:
		return c.modifiedFollowing(date)
	case PRECEDING:
		return c.preceding(date)
	case MODIFIED_PRECEDING:
		return c.modifiedPreceding(date)
	default:
		panic("Invalid BusinessDayConvention")
	}
}

func (c *Calendar) following(date Date) Date {
	if c.IsBusinessDay(date) {
		return date
	}

	return c.following(date.AddDate(0, 0, 1))
}

func (c *Calendar) preceding(date Date) Date {
	if c.IsBusinessDay(date) {
		return date
	}

	return c.preceding(date.AddDate(0, 0, -1))
}

func (c *Calendar) modifiedFollowing(date Date) Date {
	if c.IsBusinessDay(date) {
		return date
	}

	following := c.following(date)
	if following.Month() == date.Month() {
		return following
	} else {
		return c.preceding(date)
	}
}

func (c *Calendar) modifiedPreceding(date Date) Date {
	if c.IsBusinessDay(date) {
		return date
	}

	preceding := c.preceding(date)
	if preceding.Month() == date.Month() {
		return preceding
	} else {
		return c.following(date)
	}
}

func (c *Calendar) addHoliday(dateString string, name string) {
	// If the date exists in the holidays map, append the name to the slice
	// Else, create a new slice with the name
	c.holidays[dateString] = append(c.holidays[dateString], name)
}
