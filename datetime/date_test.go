package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	d, err := NewDate()
	tm := time.Now()

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if d.Year != tm.Year() {
		t.Errorf("d.Year = %d, expected %d", d.Year, tm.Year())
	}
	if d.Month != int(tm.Month()) {
		t.Errorf("d.Month = %d, expected %d", d.Month, int(tm.Month()))
	}
	if d.Day != tm.Day() {
		t.Errorf("d.Day = %d, expected %d", d.Day, tm.Day())
	}
}

func TestNewDateStr1(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(fmt.Sprintf("%d-%02d-%02d", year, month, day))

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if d.Year != year {
		t.Errorf("d.Year = %d, expected %d", d.Year, year)
	}
	if d.Month != month {
		t.Errorf("d.Month = %d, expected %d", d.Month, month)
	}
	if d.Day != day {
		t.Errorf("d.Day = %d, expected %d", d.Day, day)
	}

}

func TestNewDateStr2(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(fmt.Sprintf("%02d-%02d-%d", day, month, year))

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if d.Year != year {
		t.Errorf("d.Year = %d, expected %d", d.Year, year)
	}
	if d.Month != month {
		t.Errorf("d.Month = %d, expected %d", d.Month, month)
	}
	if d.Day != day {
		t.Errorf("d.Day = %d, expected %d", d.Day, day)
	}

}

func TestNewDateInt(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if d.Year != year {
		t.Errorf("d.Year = %d, expected %d", d.Year, year)
	}
	if d.Month != month {
		t.Errorf("d.Month = %d, expected %d", d.Month, month)
	}
	if d.Day != day {
		t.Errorf("d.Day = %d, expected %d", d.Day, day)
	}

}

func TestAddDays(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	d.AddDays(3)
	expected := day + 3
	if d.Year != year || d.Month != month || d.Day != expected {
		t.Errorf("d.Year = %d, d.Month = %d, d.Day = %d. Expected d.Year = %d, d.Month = %d, d.Day = %d",
			d.Year, d.Month, d.Day, year, month, expected)
	}
}

func TestAddWeeks(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	d.AddWeeks(2)
	expected := day + 14
	if d.Year != year || d.Month != month || d.Day != expected {
		t.Errorf("d.Year = %d, d.Month = %d, d.Day = %d. Expected d.Year = %d, d.Month = %d, d.Day = %d",
			d.Year, d.Month, d.Day, year, month, expected)
	}
}

func TestAddMonths(t *testing.T) {
	d, err := NewDate("2014-01-31")

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	d.AddMonths(3)
	if d.String() != "2014-05-01" {
		t.Errorf("2014-01-31 AddMonths(3): expected 2014-05-01, got %s", d.String())
	}
}

func TestDdmmyyyy(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	expected := fmt.Sprintf("%02d-%02d-%d", day, month, year)
	if d.Ddmmyyyy() != expected {
		t.Errorf("d.Ddmmyyyy = %s, expected %s", d.Ddmmyyyy(), expected)
	}
}

func TestString(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	expected := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	if d.String() != expected {
		t.Errorf("d.Yyyymmdd = %s, expected %s", d.String(), expected)
	}
}

func TestYyyymmdd(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	expected := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	if d.Yyyymmdd() != expected {
		t.Errorf("d.Yyyymmdd = %s, expected %s", d.Yyyymmdd(), expected)
	}
}

func TestCopy(t *testing.T) {
	year, month, day := 2014, 3, 4
	d, err := NewDate(year, month, day)

	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	dc := d.Copy()
	if dc.Year != year || dc.Month != month || dc.Day != day {
		t.Errorf("dc.Year = %d, dc.Month = %d, dc.Day = %d. Expected dc.Year = %d, dc.Month = %d, dc.Day = %d",
			dc.Year, dc.Month, dc.Day, year, month, day)
	}
}

func TestSub(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate("2014-08-10")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	expected := 20
	diff := d2.Sub(d1)
	if diff != expected {
		t.Errorf("2014-08-10.diff(2014-07-21) = %d, expected %d", diff, expected)
	}
}

func TestAge(t *testing.T) {
	birthdate, err := NewDate("1965-08-22")
	referencedate, err := NewDate("2014-09-05")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	expectedYears := 49
	expectedMonths := 0
	years, months := birthdate.Age(referencedate)

	if expectedYears != years || expectedMonths != months {
		t.Errorf("Birthdate: %s, referencedate: %s, expected years: %d, got %d, expected months: %d, got %d",
			birthdate.String(), referencedate.String, expectedYears, years, expectedMonths, months)
	}
}

func TestEq(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate(2014, 7, 21)
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if !d1.Eq(d2) {
		t.Errorf("d1.Eq(d2) = false, expected true")
	}
}

func TestNe(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate("2014-07-22")
	d3, err := NewDate("2014-07-21")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if !d1.Ne(d2) {
		t.Errorf("Expected %s != %s", d1.String(), d2.String())
	}
	if d1.Ne(d3) {
		t.Errorf("Didn't expect %s != %s", d1.String(), d3.String())
	}
}

func TestGe(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate("2014-07-21")
	d3, err := NewDate("2014-07-20")
	d4, err := NewDate("2014-07-22")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if !d1.Ge(d2) {
		t.Errorf("Expected %s >= %s", d1.String(), d2.String())
	}
	if !d1.Ge(d3) {
		t.Errorf("Expected %s >= %s", d1.String(), d3.String())
	}
	if d1.Ge(d4) {
		t.Errorf("Didn't expect %s >= %s", d1.String(), d4.String())
	}
}

func TestGt(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate("2014-07-20")
	d3, err := NewDate("2014-07-22")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if !d1.Gt(d2) {
		t.Errorf("Expected %s > %s", d1.String(), d2.String())
	}
	if d1.Gt(d3) {
		t.Errorf("Didn't expect %s > %s", d1.String(), d3.String())
	}
}

func TestLe(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate("2014-07-21")
	d3, err := NewDate("2014-07-22")
	d4, err := NewDate("2014-07-20")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if !d1.Le(d2) {
		t.Errorf("Expected %s <= %s", d1.String(), d2.String())
	}
	if !d1.Le(d3) {
		t.Errorf("Expected %s <= %s", d1.String(), d3.String())
	}
	if d1.Le(d4) {
		t.Errorf("Didn't expect %s <= %s", d1.String(), d4.String())
	}
}

func TestLt(t *testing.T) {
	d1, err := NewDate("2014-07-21")
	d2, err := NewDate("2014-07-22")
	d3, err := NewDate("2014-07-20")
	if err != nil {
		t.Fatalf("NewDate(): %s", err.Error())
	}
	if !d1.Lt(d2) {
		t.Errorf("Expected %s < %s", d1.String(), d2.String())
	}
	if d1.Lt(d3) {
		t.Errorf("Didn't expect %s < %s", d1.String(), d3.String())
	}
}

func TestNewDatePattern(t *testing.T) {
	datepattern, err := NewDatePattern(WEEKLY, "2014-07-22", "2014-09-19")
	if err != nil {
		t.Fatalf("NewDatePattern(): %s", err.Error())
	}
	if datepattern.Interval != WEEKLY {
		t.Errorf("Expected Interval to be WEEKLY")
	}
	if datepattern.Begin.String() != "2014-07-22" {
		t.Errorf("Expected Begin = 2014-07-22, got %s", datepattern.Begin.String())
	}
	if datepattern.End.String() != "2014-09-19" {
		t.Errorf("Expected End = 2014-09-19, got %s", datepattern.End.String())
	}
	datepattern, err = NewDatePattern(WEEKLY)
	if err == nil {
		t.Errorf("Didn't expect err to be nil without Begin or End")
	}
	begin, err := NewDate("2014-07-22")
	end, err := NewDate("2014-09-19")
	if err != nil {
		t.Fatalf("%s", err)
	}
	datepattern, err = NewDatePattern(WEEKLY, begin, end)
	if err != nil {
		t.Fatalf("NewDatePattern(): %s", err.Error())
	}
	if datepattern.Interval != WEEKLY {
		t.Errorf("Expected Interval to be WEEKLY")
	}
	if datepattern.Begin.String() != "2014-07-22" {
		t.Errorf("Expected Begin = 2014-07-22, got %s", datepattern.Begin.String())
	}
	if datepattern.End.String() != "2014-09-19" {
		t.Errorf("Expected End = 2014-09-19, got %s", datepattern.End.String())
	}

}

func TestContains(t *testing.T) {
	datepattern, err := NewDatePattern(WEEKLY, "2014-07-22", "2014-09-19")
	if err != nil {
		t.Fatalf("NewDatePattern(): %s", err.Error())
	}
	contains := datepattern.Contains("2014-08-05")
	if !contains {
		t.Errorf("Expected 2014-08-05 to be in weekly interval from 2014-07-22 to 2014-09-19")
	}

}

func TestLookup(t *testing.T) {
	type Persoon struct {
		Key  string
		Name string
	}
	lookup := make(map[string]*Persoon)
	lijst := make([]*Persoon, 0)

	persoon := lookup["henk"]
	if persoon == nil {
		persoon = &Persoon{"henk", "Henk Burgstra"}
		lookup["henk"] = persoon
		lijst = append(lijst, persoon)
	}
	persoon = lookup["henk"]
	persoon.Name = "Een andere naam"
	for _, persoon := range lijst {
		fmt.Println(persoon.Name)
	}
}
