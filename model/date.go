package model

import (
	"encoding/json"
	"time"
)

const format = "2006-01-02"

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {

	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t, _ := time.Parse(format, v)
	*d = Date(t)
	return nil
}

func (d *Date) MarshallJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Date) String() string {
	return time.Time(d).Format(format)
}
