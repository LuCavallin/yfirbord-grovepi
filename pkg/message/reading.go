package message

import "time"

type Reading struct {
	name string `json:"name"`
	value interface{} `json:"value"`
	time time.Time `json:"time"`
}

func NewReading(name string, value interface{}) *Reading {
	return &Reading{name, value, time.Now()}
}