package swagger

import (
	"encoding/json"
	"time"
)

type NullInt struct {
	Value int
	Valid bool
}

type NullInt64 struct {
	Value int64
	Valid bool
}

type NullFloat32 struct {
	Value float32
	Valid bool
}

type NullFloat64 struct {
	Value float64
	Valid bool
}

type NullTime struct {
	Value time.Time
	Valid bool
}

type NullBool struct {
	Value bool
	Valid bool
}

type NullString struct {
	Value string
	Valid bool
}

func (i *NullInt) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	// The key isn't set to null
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *NullFloat32) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
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

func (i *NullFloat64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
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

func (i *NullTime) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
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

func (i *NullBool) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
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

func (i *NullString) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
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
