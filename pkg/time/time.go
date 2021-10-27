package time

import (
	"context"
	xtime "time"
)
type JsonTime xtime.Time
const (
	timeFormat = "2006-01-02 15:04:05"
)
func (j JsonTime) MarshalJSON() ([]byte, error) {
	if xtime.Time(j).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + xtime.Time(j).Local().Format(timeFormat) + `"`), nil
}
func (j JsonTime) IsZero() bool {
	return xtime.Time(j).Local().IsZero()
}
func (j *JsonTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	tf, err := xtime.ParseInLocation(`"`+timeFormat+`"`, string(data), xtime.Local)
	*j = JsonTime(tf)
	return err
}

func (j JsonTime) String() string {
	return xtime.Time(j).Format(timeFormat)
}
type Duration xtime.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := xtime.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// Shrink will decrease the duration by comparing with context's timeout duration
// and return new timeout\context\CancelFunc.
func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if ctimeout := xtime.Until(deadline); ctimeout < xtime.Duration(d) {
			// deliver small timeout
			return Duration(ctimeout), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, xtime.Duration(d))
	return d, ctx, cancel
}
