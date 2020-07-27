package websocket

import (
	"encoding/json"
	"time"
)

type JSONInt64 struct {
	Value int64
	Valid bool
	Set   bool
}

type JSONFloat32 struct {
	Value float32
	Valid bool
	Set   bool
}

type JSONFloat64 struct {
	Value float64
	Valid bool
	Set   bool
}

type JSONTime struct {
	Value time.Time
	Valid bool
	Set   bool
}

type JSONBool struct {
	Value bool
	Valid bool
	Set   bool
}

type JSONString struct {
	Value string
	Valid bool
	Set   bool
}

func (i *JSONInt64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *JSONFloat32) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp float32
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *JSONFloat64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp float64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *JSONString) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *JSONBool) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *JSONTime) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp time.Time
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}
