package datetime

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Interval int

const (
	INVALID Interval = iota
	DAILY
	WEEKLY
	TWO_WEEKLY
	THREE_WEEKLY
	FOUR_WEEKLY
	MONTHLY
	FIRST_OF_MONTH
	SECOND_OF_MONTH
	THIRD_OF_MONTH
	FOURTH_OF_MONTH
	QUARTERLY
	YEARLY
)

func (interval Interval) String() string {
	switch interval {
	case INVALID:
		return "invalid"
	case DAILY:
		return "every day"
	case WEEKLY:
		return "every week"
	case TWO_WEEKLY:
		return "every two weeks"
	case THREE_WEEKLY:
		return "every three weeks"
	case FOUR_WEEKLY:
		return "every four weeks"
	case MONTHLY:
		return "every month"
	case FIRST_OF_MONTH:
		return "every first mon, tue, ... of the month"
	case SECOND_OF_MONTH:
		return "every second mon, tue, ... of the month"
	case THIRD_OF_MONTH:
		return "every third mon, tue, ... of the month"
	case FOURTH_OF_MONTH:
		return "every fourth mon, tue, ... of the month"
	case QUARTERLY:
		return "every quarter"
	case YEARLY:
		return "every year"
	default:
		return "unknown"
	}
}

type PeriodType int

const (
	DAY PeriodType = iota + 1
	WEEK
	MONTH
	QUARTER
	YEAR
)

var (
	reYyyymmdd = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	reDdmmyyyy = regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
)

type DateErrorType int

const (
	ERR_PARAM_MISMATCH DateErrorType = iota + 1
	ERR_BAD_FORMAT
)

type DateError struct {
	errType DateErrorType
	msg     string
}

func NewDateError(errType DateErrorType, msg string) *DateError {
	err := new(DateError)
	err.errType = errType
	err.msg = msg
	return err
}

func (e *DateError) Error() string {
	return e.msg
}

type Date struct {
	Year  int
	Month int
	Day   int
	Time  time.Time
}

func NewDate(args ...interface{}) (date *Date, err error) {
	location, _ := time.LoadLocation("Local")
	date = new(Date)
	switch len(args) {
	case 0:
		date.Time = time.Now()
		date.Year = date.Time.Year()
		date.Month = int(date.Time.Month())
		date.Day = date.Time.Day()
	case 1:
		dateString := args[0].(string)
		if reYyyymmdd.MatchString(dateString) {
			e := strings.Split(dateString, "-")
			date.Year, err = strconv.Atoi(e[0])
			date.Month, err = strconv.Atoi(e[1])
			date.Day, err = strconv.Atoi(e[2])
		} else if reDdmmyyyy.MatchString(dateString) {
			e := strings.Split(dateString, "-")
			date.Day, err = strconv.Atoi(e[0])
			date.Month, err = strconv.Atoi(e[1])
			date.Year, err = strconv.Atoi(e[2])
		} else {
			err = NewDateError(ERR_BAD_FORMAT,
				fmt.Sprintf("Expected yyyy-mm-dd or dd-mm-yyyy, got %s", dateString))
		}
	case 3:
		date.Year, _ = args[0].(int)
		date.Month, _ = args[1].(int)
		date.Day, _ = args[2].(int)
	default:
		err = NewDateError(ERR_PARAM_MISMATCH,
			fmt.Sprintf("0, 1 or 3 parameters expected, got %d", len(args)))
	}
	date.Time = time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, location)
	return
}

func Today() (today *Date) {
	today, _ = NewDate()
	return
}

func (d *Date) Copy() (date *Date) {
	date, _ = NewDate(d.Year, d.Month, d.Day)
	return
}

func (d *Date) Add(what PeriodType, count int) {
	switch what {
	case DAY:
		d.Time = d.Time.AddDate(0, 0, count)
	case WEEK:
		d.Time = d.Time.AddDate(0, 0, count*7)
	case MONTH:
		d.Time = d.Time.AddDate(0, count, 0)
	case QUARTER:
		d.Time = d.Time.AddDate(0, count*3, 0)
	case YEAR:
		d.Time = d.Time.AddDate(count, 0, 0)
	}
	year, month, day := d.Time.Date()
	d.Year = year
	d.Month = int(month)
	d.Day = day
}

func (d *Date) AddDays(count int) {
	d.Add(DAY, count)
}

func (d *Date) AddWeeks(count int) {
	d.Add(WEEK, count)
}

func (d *Date) AddMonths(count int) {
	d.Add(MONTH, count)
}

func (d *Date) AddQuarters(count int) {
	d.Add(QUARTER, count)
}

func (d *Date) AddYears(count int) {
	d.Add(YEAR, count)
}

func (d *Date) Sub(date *Date) int {
	duration := d.Time.Sub(date.Time)
	return int(duration.Hours() / 24)
}

func (d *Date) Age(args ...*Date) (years int, months int) {
	var referenceDate *Date
	if len(args) == 0 {
		referenceDate, _ = NewDate()
	} else {
		referenceDate = args[0]
	}
	m := (referenceDate.Year-d.Year)*12 + (referenceDate.Month - d.Month)
	if referenceDate.Day < d.Day {
		m--
	}
	years = m / 12
	months = m % 12
	return
}

func (d *Date) Eq(date *Date) bool {
	return (d.Time == date.Time)
}

func (d *Date) Ne(date *Date) bool {
	return (d.Time != date.Time)
}

func (d *Date) Gt(date *Date) bool {
	return d.Time.After(date.Time)
}

func (d *Date) Ge(date *Date) bool {
	return d.Eq(date) || d.Gt(date)
}

func (d *Date) Lt(date *Date) bool {
	return d.Time.Before(date.Time)
}

func (d *Date) Le(date *Date) bool {
	return d.Eq(date) || d.Lt(date)
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d *Date) Ddmmyyyy() string {
	return fmt.Sprintf("%02d-%02d-%d", d.Day, d.Month, d.Year)
}

func (d *Date) Yyyymmdd() string {
	return d.String()
}

type DatePattern struct {
	Interval Interval
	Begin    *Date
	End      *Date
}

func NewDatePattern(interval Interval, args ...interface{}) (repeater *DatePattern, err error) {
	if len(args) == 0 || len(args) > 3 {
		err = NewDateError(ERR_PARAM_MISMATCH,
			fmt.Sprintf("1 or 2 args expected, got %d", len(args)))
		return
	}
	repeater = new(DatePattern)
	repeater.Interval = interval

	for i, arg := range args {
		var date *Date

		switch arg := arg.(type) {
		case string:
			date, err = NewDate(arg)
		case *Date:
			date = arg
		}
		if err != nil {
			return
		}
		if date != nil && i == 0 {
			repeater.Begin = date
		} else if date != nil && i == 1 {
			repeater.End = date
		}
	}
	return
}

func (dp *DatePattern) Contains(arg interface{}) bool {
	var date *Date
	var err error

	switch arg := arg.(type) {
	case string:
		date, err = NewDate(arg)
	case *Date:
		date = arg
	default:
		return false
	}

	if err != nil {
		return false
	}

	switch dp.Interval {
	case DAILY:
		return dp.everyDayContains(date)
	case WEEKLY:
		return dp.everyXWeeksContains(date, 1)
	case TWO_WEEKLY:
		return dp.everyXWeeksContains(date, 2)
	case THREE_WEEKLY:
		return dp.everyXWeeksContains(date, 3)
	case FOUR_WEEKLY:
		return dp.everyXWeeksContains(date, 4)
	case MONTHLY:
		return dp.everyMonthContains(date)
	}
	return false
}

func (dp *DatePattern) everyDayContains(date *Date) bool {
	return date.Ge(dp.Begin) && (dp.End == nil || date.Le(dp.End))
}

func (dp *DatePattern) everyXWeeksContains(date *Date, weeks int) bool {
	if dp.End != nil && dp.End.Eq(date) {
		return true
	}
	if dp.Begin.Gt(date) || (dp.End != nil && dp.End.Lt(date)) {
		return false
	}
	compare := dp.Begin.Copy()

	for compare.Le(date) {
		if compare.Eq(date) {
			return true
		}
		compare.Add(WEEK, weeks)
	}
	return false
}

func (dp *DatePattern) everyMonthContains(date *Date) bool {
	return false
}
