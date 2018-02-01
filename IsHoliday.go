package holidayJp

import (
	"io/ioutil"
	"net/http"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var holiday map[string]string
var Client = http.DefaultClient

//var Loc = time.FixedZone("Asia/Tokyo", 9*60*60)
var Loc = time.FixedZone("Asia/Tokyo-3", 6*60*60) // Change date at 3:00am.

// IsHoliday is a function wheather Japanese holiday or not.
func IsHoliday(t time.Time) bool {
	if holiday == nil {
		holiday = make(map[string]string)
		for {
			resp, err := Client.Get("https://raw.githubusercontent.com/holiday-jp/holiday_jp/master/holidays.yml")
			if err == nil {
				defer resp.Body.Close()
				if resp.StatusCode == 200 {
					b, _ := ioutil.ReadAll(resp.Body)
					_ = yaml.Unmarshal(b, &holiday)
					break
				}
			}
			panic("Cannot get holidays.yml")
		}
	}

	_, ok := holiday[t.In(Loc).Format("2006-01-02")]
	return ok
}

// IsSunday is a function wheather Sunday or not.
func IsSunday(t time.Time) bool {
	return t.In(Loc).Weekday() == time.Sunday
}

// IsSaturday is a function wheather Saturday or not.
func IsSaturday(t time.Time) bool {
	return t.In(Loc).Weekday() == time.Saturday
}
