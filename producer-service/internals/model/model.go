package model

import "encoding/json"

type Error struct {
	Message string `json:"message"`
}

func (e *Error) String() []byte {
	parsed, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return parsed
}
