package domain

import (
	"errors"
	"time"
)

type Schema struct {
	ID          int       `json:"id,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	SpeedRange  int       `json:"speed_range"`
	Features    []string  `json:"features"`
}

func (s *Schema) Check() error {
	if !contains(s.Type, []string{"sedan", "van", "suv", "motor-bike"}) {
		return errors.New("vehicle type can only be one of sedan, van, suv, motor-bike")
	}

	if !contains(s.Color, []string{"red", "green", "blue"}) {
		return errors.New("vehicle color can only be one of red, green or blue")
	}

	if s.SpeedRange > 240 || s.SpeedRange < 0 {
		return errors.New("vehicle speedrange should be between 0 and 240")
	}

	for _, elem := range s.Features {
		if !contains(elem, []string{"sunroof", "panorama", "auto-parking", "surround-system"}) {
			return errors.New("vehicle does not offer option " + elem)
		}
	}

	return nil
}

func contains(value string, slice []string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}
