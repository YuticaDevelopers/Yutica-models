package common

import (
	"time"
)

type JSONTime time.Time

const timeLayout = time.RFC3339 // ISO 8601: "2023-04-24T15:04:05Z07:00"

// MarshalJSON converts JSONTime to JSON string
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := "\"" + time.Time(t).Format(timeLayout) + "\""
	return []byte(stamp), nil
}

// UnmarshalJSON parses JSON string to JSONTime
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"`+timeLayout+`"`, string(data))
	if err != nil {
		return err
	}
	*t = JSONTime(parsed)
	return nil
}

// String returns the ISO time format
func (t JSONTime) String() string {
	return time.Time(t).Format(timeLayout)
}

// ToTime converts to standard time.Time
func (t JSONTime) ToTime() time.Time {
	return time.Time(t)
}
