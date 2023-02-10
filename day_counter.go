package calendar

import "fmt"

type DayCountConvention string

const (
	US_30_360  = "US_30_360"
	ACT_ACT    = "ACT_ACT"
	ACT_360    = "ACT_360"
	ACT_365    = "ACT_365"
	EUR_30_360 = "EUR_30_360"
)

type DayCounter struct {
	DayCountConvention DayCountConvention
}

func GetDayCounter(dayCountConvention DayCountConvention) (dc DayCounter) {
	return DayCounter{
		DayCountConvention: dayCountConvention,
	}
}

func (dc *DayCounter) YearFrac(startDate Date, endDate Date) (yearFrac float64) {
	if startDate.After(endDate) {
		return 0
	}

	numerator := float64(dc.diffDates(startDate, endDate))
	denom := dc.calcAnnualBasis(startDate, endDate)

	return numerator / denom
}

func (dc *DayCounter) flsLeapYear(year int) bool {
	if year%4 > 0 {
		return false
	} else if year%100 > 0 {
		return true
	} else if year%400 == 0 {
		return true
	} else {
		return false
	}
}

func (dc *DayCounter) flsEndOfMonth(day, month, year int) bool {
	if day < 1 || day > 31 || month < 1 || month > 12 {
		panic(fmt.Errorf("wrong parameters day(%d), month(%d)", day, month))
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return day == 31
	case 4, 6, 9, 11:
		return day == 30
	}

	// When month = 2
	if dc.flsLeapYear(year) {
		return day == 29
	}
	return day == 28
}

func (dc *DayCounter) days360(startYear, endYear, startMonth, endMonth, startDay, endDay int) float64 {
	var yearGap, monthGap, dayGap float64

	yearGap = float64(endYear) - float64(startYear)
	monthGap = float64(endMonth) - float64(startMonth)
	dayGap = float64(endDay) - float64(startDay)
	return yearGap*360 + monthGap*30 + dayGap
}

func (dc *DayCounter) days360Nasd(startDate, endDate Date, useEndOfMonth bool) float64 {
	startDay := startDate.Day()
	startMonth := int(startDate.Month())
	startYear := startDate.Year()
	endDay := endDate.Day()
	endMonth := int(endDate.Month())
	endYear := endDate.Year()

	if (endMonth == 2 && dc.flsEndOfMonth(endDay, endMonth, endYear)) && ((startMonth == 2 && dc.flsEndOfMonth(endDay, endMonth, endYear)) || dc.DayCountConvention == ACT_365) {
		endDay = 30
	}
	if endDay == 31 && (startDay >= 30 || dc.DayCountConvention == ACT_365) {
		endDay = 30
	}
	if startDay == 31 {
		startDay = 30
	}
	if useEndOfMonth && startMonth == 2 && dc.flsEndOfMonth(startDay, startMonth, startYear) {
		startDay = 30
	}
	return dc.days360(startYear, endYear, startMonth, endMonth, startDay, endDay)
}

func (dc *DayCounter) days360Euro(startDate, endDate Date) float64 {
	startDay := startDate.Day()
	startMonth := int(startDate.Month())
	startYear := startDate.Year()
	endDay := endDate.Day()
	endMonth := int(endDate.Month())
	endYear := endDate.Year()

	if startDay == 31 {
		startDay = 30
	}

	if endDay == 31 {
		endDay = 30
	}

	return dc.days360(startYear, endYear, startMonth, endMonth, startDay, endDay)
}

func (dc *DayCounter) diffDates(startDate, endDate Date) float64 {
	if dc.DayCountConvention == US_30_360 {
		return dc.days360Nasd(startDate, endDate, true)
	} else if dc.DayCountConvention == ACT_ACT || dc.DayCountConvention == ACT_360 || dc.DayCountConvention == ACT_365 {
		return float64(endDate.Sub(startDate))
	} else {
		return dc.days360Euro(startDate, endDate)
	}
}

func (dc *DayCounter) calcAnnualBasis(startDate, endDate Date) float64 {
	if dc.DayCountConvention == US_30_360 || dc.DayCountConvention == ACT_360 || dc.DayCountConvention == EUR_30_360 {
		return 360
	} else if dc.DayCountConvention == ACT_365 {
		return 365
	} else {
		startDay := startDate.Day()
		startMonth := int(startDate.Month())
		startYear := startDate.Year()
		endDay := endDate.Day()
		endMonth := int(endDate.Month())
		endYear := endDate.Year()

		if startYear == endYear {
			if dc.flsLeapYear(startYear) {
				return 366
			} else {
				return 365
			}
		} else if endYear-1 == startYear && (startMonth > endMonth || (startMonth == endMonth && (startDay >= endDay))) {
			if dc.flsLeapYear(startYear) {
				if startMonth <= 2 {
					return 366
				} else {
					return 365
				}
			} else if dc.flsLeapYear(endYear) {
				if endMonth > 2 || (endMonth == 2 && endDay == 29) {
					return 366
				} else {
					return 365
				}
			} else {
				return 365
			}
		} else {
			daysSum := 0.0
			for iYear := startYear; iYear <= endYear; iYear++ {
				if dc.flsLeapYear(iYear) {
					daysSum += 366.0
				} else {
					daysSum += 365.0
				}
			}
			return daysSum / (float64(endYear) - float64(startYear) + 1)
		}
	}
}
