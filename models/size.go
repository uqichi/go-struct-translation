package models

import (
	"encoding/json"
	"strconv"
)

type Style int

const (
	StyleUndefined Style = iota
	StyleAmericano
	StyleEspresso
	StyleRigoretto
)

var StyleMap = map[Style]string{
	StyleAmericano: "americano",
	StyleEspresso:  "espresso",
	StyleRigoretto: "rigoretto",
}

func (e Style) String() string {
	return StyleMap[e]
}

func (e *Style) Parse(s string) error {
	for k, v := range StyleMap {
		if v == s {
			*e = k
			return nil
		}
	}
	return nil
}

func (e Style) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *Style) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	var v Style
	err = v.Parse(s)
	*e = v
	return err
}
