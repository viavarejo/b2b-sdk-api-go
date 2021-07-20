package response

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (m *CustomTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	//to remove double quotes from string
	var dateStr string = strings.Trim(string(data), "\"")
	tt, err := time.Parse("2006-01-02T15:04:05", dateStr)
	//tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = CustomTime{tt}
	return err
}
