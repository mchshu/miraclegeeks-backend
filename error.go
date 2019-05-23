package main

import (
	"encoding/json"
)

type AppErr struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func NewAppErr(err error) *AppErr {
	return &AppErr{
		Errcode: 1,
		Errmsg:  err.Error(),
	}
}

func (app AppErr) String() string {
	data, _ := json.Marshal(app)

	return string(data)
}
