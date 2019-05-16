package datetimeutil

import (
	"log"
	"time"
)

var (
	_layoutDate     = "2006-01-02"
	_layoutDatetime = "2006-01-02 15:04:05"
)

func GetDateBeforeToday(days int) time.Time {
	return time.Now().AddDate(0, 0, -days)
}

func GetDatetimeBeforeNow(seconds int) time.Time {
	return time.Now().Add(-time.Second * time.Duration(seconds))
}

// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func GetDatetimeDuration(s string) time.Time {
	d, _ := time.ParseDuration(s)
	return time.Now().Add(d)
}

func GetDatetimeBefore(datetime time.Time, seconds int) time.Time {
	return datetime.Add(-time.Second * time.Duration(seconds))
}

func GetDatetimeDiffNano(begin time.Time, end time.Time) int64 {
	return begin.UnixNano() - end.UnixNano()
}

func GetDatetimeDiff(begin time.Time, end time.Time) int64 {
	return begin.Unix() - end.Unix()
}

func GetDateStrBeforeToday(days int) string {
	return GetDateBeforeToday(days).Format(_layoutDate)
}

func GetDatetimeStrBeforeNow(seconds int) string {
	return GetDatetimeBeforeNow(seconds).Format(_layoutDatetime)
}

func GetStrFromDatetime(datetime time.Time) string {
	return datetime.Format(_layoutDatetime)
}

func GetDatetimeFromStr(datetimeStr string) time.Time {
	t, err := time.Parse(_layoutDatetime, datetimeStr)
	if err != nil {
		log.Fatalln(err)
	}
	return t
}

func GetCurrentTimeStr() string {
	return time.Now().Format(_layoutDatetime)
}

func GetCurrentTimeUnix() int64 {
	return time.Now().Unix()
}

func GetCurrentTimeUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetCurrentTime() time.Time {
	return time.Now()
}
