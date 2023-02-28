package convertor

import "time"

func StringToTime(dateTime string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, dateTime+"Z")
	if err != nil {
		return t, err
	}
	return t, nil
}
