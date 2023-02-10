package calendar

import (
	"math"
	"testing"
	"time"
)

func TestDayCounter(t *testing.T) {
	type Input struct {
		StartDate          Date
		EndDate            Date
		DayCountConvention DayCountConvention
	}

	tolerance := 0.000001
	testCases := make(map[Input]float64)
	testCases[Input{
		Date{time.Date(2012, time.January, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		US_30_360,
	}] = 0.080555556
	testCases[Input{
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		US_30_360,
	}] = 0.086111111
	testCases[Input{
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		US_30_360,
	}] = 0.083333333
	testCases[Input{
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		US_30_360,
	}] = 0.083333333
	testCases[Input{
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2013, time.May, 30, 0, 0, 0, 0, time.Now().Location())},
		US_30_360,
	}] = 1.0

	testCases[Input{
		Date{time.Date(2012, time.January, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		ACT_ACT,
	}] = 0.079234973
	testCases[Input{
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		ACT_ACT,
	}] = 0.084699454
	testCases[Input{
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		ACT_ACT,
	}] = 0.081967213
	testCases[Input{
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		ACT_ACT,
	}] = 0.084699454
	testCases[Input{
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2013, time.May, 30, 0, 0, 0, 0, time.Now().Location())},
		ACT_ACT,
	}] = 0.997260274

	testCases[Input{
		Date{time.Date(2012, time.January, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		ACT_360,
	}] = 0.080555556
	testCases[Input{
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		ACT_360,
	}] = 0.086111111
	testCases[Input{
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		ACT_360,
	}] = 0.083333333
	testCases[Input{
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		ACT_360,
	}] = 0.086111111
	testCases[Input{
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2013, time.May, 30, 0, 0, 0, 0, time.Now().Location())},
		ACT_360,
	}] = 1.011111111

	testCases[Input{
		Date{time.Date(2012, time.January, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		ACT_365,
	}] = 0.079452055
	testCases[Input{
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		ACT_365,
	}] = 0.084931507
	testCases[Input{
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		ACT_365,
	}] = 0.082191781
	testCases[Input{
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		ACT_365,
	}] = 0.084931507
	testCases[Input{
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2013, time.May, 30, 0, 0, 0, 0, time.Now().Location())},
		ACT_365,
	}] = 0.997260274

	testCases[Input{
		Date{time.Date(2012, time.January, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		EUR_30_360,
	}] = 0.080555556
	testCases[Input{
		Date{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		EUR_30_360,
	}] = 0.086111111
	testCases[Input{
		Date{time.Date(2012, time.March, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		EUR_30_360,
	}] = 0.083333333
	testCases[Input{
		Date{time.Date(2012, time.April, 30, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		EUR_30_360,
	}] = 0.083333333
	testCases[Input{
		Date{time.Date(2012, time.May, 31, 0, 0, 0, 0, time.Now().Location())},
		Date{time.Date(2013, time.May, 30, 0, 0, 0, 0, time.Now().Location())},
		EUR_30_360,
	}] = 1.0

	for input, expected := range testCases {
		dc := DayCounter{DayCountConvention: input.DayCountConvention}
		result := dc.YearFrac(input.StartDate, input.EndDate)
		if math.Abs(result-expected) > tolerance {
			t.Errorf(`DayCounter{DayCounterConvention(%v)}.YearFrac(%v, %v) is %f, not %f`, input.DayCountConvention, input.StartDate, input.EndDate, result, expected)
		}
	}
}
