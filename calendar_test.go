package calendar

import "testing"

func TestGetCalendarOfSouthKoreaSeoulWithFollowingConvention(t *testing.T) {
	calendar := GetCalendar([]BusinessDay{SOUTH_KOREA_SEOUL}, FOLLOWING)
	var testDate Date
	// 2023-01-24 should be a holiday
	testDate, _ = GetDateByString(LayoutYYYY_MM_DD, "2023-01-24")
	if isBusinessDay := calendar.IsBusinessDay(testDate); isBusinessDay {
		t.Errorf("Expected %s to be a holiday", testDate.String())
	}

	// The adjusted date of 2023-01-24 should be 2023-01-25
	testDate, _ = GetDateByString(LayoutYYYY_MM_DD, "2023-01-24")
	if adjustedDate := calendar.Adjust(testDate); adjustedDate.String() != "2023-01-25" {
		t.Errorf("Expected %s to be adjusted to 2023-01-25", testDate.String())
	}

	// 2023-01-25 should be a business day
	testDate, _ = GetDateByString(LayoutYYYY_MM_DD, "2023-01-25")
	if isBusinessDay := calendar.IsBusinessDay(testDate); !isBusinessDay {
		t.Errorf("Expected %s to be a business day", testDate.String())
	}

	// The adjusted date of 2023-01-25 should be 2023-01-25
	testDate, _ = GetDateByString(LayoutYYYY_MM_DD, "2023-01-25")
	if adjustedDate := calendar.Adjust(testDate); adjustedDate.String() != "2023-01-25" {
		t.Errorf("Expected %s to be adjusted to 2023-01-25", testDate.String())
	}
}
