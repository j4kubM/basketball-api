package model

import (
	"encoding/json"
	"fmt"
	"time"
)

const format = "dd-mm-yyyy"

type Date struct{ time.Time }

func (d *Date) UnmarshallJSON(data []byte) error {

	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t, _ := time.Parse(format, v)
	d = &Date{t}
	return nil
}

func (d *Date) MarshallJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) String() string {
	return fmt.Sprintf("%q", d.Format(format))
}
