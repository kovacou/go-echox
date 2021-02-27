package echox

import (
	"strings"
	"time"

	"github.com/kovacou/go-convert"
	"github.com/kovacou/go-types"
	"github.com/labstack/echo"
)

// QueryScopes returns parsed scopes.
func QueryScopes(ctx echo.Context) types.Strings {
	scopes := strings.TrimSpace(strings.TrimRight(ctx.QueryParam("scopes"), ","))
	if len(scopes) == 0 {
		return types.Strings{"all"}
	}
	return strings.Split(scopes, ",")
}

// QueryDate returns the given query as date.
func QueryDate(ctx echo.Context, key string) types.Date {
	return ParseDate(ctx.QueryParam(key))
}

// QueryDateOrNow returns the given query as date and if the
// date is not valid, current date will be used.
func QueryDateOrNow(ctx echo.Context, key string) types.Date {
	d := ParseDate(ctx.QueryParam(key))
	if d.IsZero() {
		d = types.NewDate(time.Now().UTC())
	}
	return d
}

// QueryPage returns the page query.
func QueryPage(ctx echo.Context) uint64 {
	p := convert.Uint64(ctx.QueryParam("page"))
	if p < 1 {
		p = 1
	}
	return p
}

// parseDate parse the given string as date.
func ParseDate(v string) types.Date {
	utc := time.Now().UTC()

	switch v {
	case "today", "now":
		return types.NewDate(utc)
	case "yesterday", "previousDay":
		return types.NewDate(utc.AddDate(0, 0, -1))
	case "tomorrow", "nextDay":
		return types.NewDate(utc.AddDate(0, 0, 1))
	case "currentMonth", "currentMonthStart":
		return types.NewDate(time.Date(utc.Year(), utc.Month(), 1, 0, 0, 0, 0, time.UTC))
	case "currentMonthEnd":
		return types.NewDate(time.Date(utc.Year(), utc.Month()+1, 0, 0, 0, 0, 0, time.UTC))
	case "currentYear", "currentYearStart":
		return types.NewDate(time.Date(utc.Year(), time.January, 1, 0, 0, 0, 0, time.UTC))
	case "currentYearEnd":
		return types.NewDate(time.Date(utc.Year()+1, time.January, 0, 0, 0, 0, 0, time.UTC))
	case "lastMonth":
		return types.NewDate(utc.AddDate(0, -1, 0))
	case "lastMonthStart":
		return types.NewDate(time.Date(utc.Year(), utc.Month()-1, 1, 0, 0, 0, 0, time.UTC))
	case "lastMonthEnd":
		return types.NewDate(time.Date(utc.Year(), utc.Month(), 0, 0, 0, 0, 0, time.UTC))
	case "lastYear":
		return types.NewDate(utc.AddDate(-1, 0, 0))
	case "lastYearStart":
		return types.NewDate(time.Date(utc.Year()-1, time.January, 1, 0, 0, 0, 0, time.UTC))
	case "lastYearEnd":
		return types.NewDate(time.Date(utc.Year(), time.January, 0, 0, 0, 0, 0, time.UTC))
	case "nextMonth":
		return types.NewDate(utc.AddDate(0, 1, 0))
	case "nextMonthStart":
		return types.NewDate(time.Date(utc.Year(), utc.Month()+1, 1, 0, 0, 0, 0, time.UTC))
	case "nextMonthEnd":
		return types.NewDate(time.Date(utc.Year(), utc.Month()+2, 0, 0, 0, 0, 0, time.UTC))
	case "nextYear":
		return types.NewDate(utc.AddDate(1, 0, 0))
	case "nextYearStart":
		return types.NewDate(time.Date(utc.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC))
	case "nextYearEnd":
		return types.NewDate(time.Date(utc.Year()+2, time.January, 0, 0, 0, 0, 0, time.UTC))
	}

	d, err := time.Parse(types.DateFormat, v)
	if err != nil {
		d, _ = time.Parse(types.DateYearMonthFormat, v)
	}
	return types.NewDate(d)
}
