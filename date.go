package calendar

import (
	"math"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func GetDate(t time.Time) Date {
	return Date{time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, 0, 0, 0, time.Now().Location(),
	)}
}

const LayoutYYYYMMDD = "20060102"
const LayoutYYYY_MM_DD = "2006-01-02"

func GetDateByString(layout, s string) (Date, error) {
	t, parseError := time.Parse(layout, s)
	if parseError != nil {
		return Date{}, parseError
	}
	return GetDate(t), nil
}

func (d Date) String() string {
	return strings.Split(d.Format(time.RFC3339), "T")[0]
}

func (d Date) YYYYMMDD() string {
	return strings.ReplaceAll(d.String(), "-", "")
}

func (d Date) Equal(c Date) bool {
	return d.Year() == c.Year() && d.Month() == c.Month() && d.Day() == c.Day()
}

func (d Date) Before(c Date) bool {
	return d.Year() < c.Year() || (d.Year() == c.Year() && d.Month() < c.Month()) || (d.Year() == c.Year() && d.Month() == c.Month() && d.Day() < c.Day())
}

func (d Date) After(c Date) bool {
	return c.Before(d)
}

func (d Date) AfterOrEqual(c Date) bool {
	return d.After(c) || d.Equal(c)
}

func (d Date) BeforeOrEqual(c Date) bool {
	return d.Before(c) || d.Equal(c)
}

func (d Date) AddDate(years int, months int, days int) Date {
	return Date{d.Time.AddDate(years, months, days)}
}

func (d Date) Sub(c Date) (days int) {
	return int(math.Round(d.Time.Sub(c.Time).Hours() / 24))
}

func (d Date) SetLocalTimezone() Date {
	return Date{time.Date(
		d.Year(),
		d.Month(),
		d.Day(),
		0, 0, 0, 0, time.Now().Location(),
	)}
}
