package time

import (
	"time"
)
type JsonTime time.Time
const (
	timeFormat = "2006-01-02 15:04:05"
)
func (j JsonTime) MarshalJSON() ([]byte, error) {
	if time.Time(j).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(j).Local().Format(timeFormat) + `"`), nil
}
func (j JsonTime) IsZero() bool {
	return time.Time(j).Local().IsZero()
}
func (j *JsonTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	tf, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*j = JsonTime(tf)
	return err
}

func (j JsonTime) String() string {
	return time.Time(j).Format(timeFormat)
}
