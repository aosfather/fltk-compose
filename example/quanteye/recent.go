package main

import (
	"encoding/json"
	"os"
)

const _File = "./last.json"

type Recent struct {
	Product string `json:"product"`
	Val     int    `json:"val"`
	Project string `json:"project"`
}

func (r *Recent) Load() {
	data, e := os.ReadFile(_File)
	if e == nil {
		json.Unmarshal(data, r)
	}
}

func (r *Recent) Save() {
	data, e := json.Marshal(r)
	if e == nil {
		os.WriteFile(_File, data, 0644)
	}
}
