package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Car struct {
	ID          int
	LastUpdated time.Time
	CreatedTime time.Time
	Type        string `json:"type"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	SpeedRange  int    `json:"speed_range"`
	Features    []string
}

func (s *Car) Check() error {
	if s.Name == "" {
		return errors.New("vehicle Name Cannot Be Empty")
	}

	if s.Type == "" {
		return errors.New("vehicle Type Cannot Be Empty")
	}

	if s.Color == "" {
		return errors.New("vehicle Color Cannot Be Empty")
	}

	if match, _ := regexp.MatchString("^[0-9]+$", strconv.Itoa(s.SpeedRange)); !match {
		return errors.New("invalid Vehicle SpeedRange Specified")
	}

	if len(s.Features) == 0 {
		return errors.New("vehicle Features Cannot Be Empty")
	}

	s.convertToLower()

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

func (s *Car) convertToLower() {
	s.Type = strings.ToLower(s.Type)
	s.Color = strings.ToLower(s.Color)

	for index, elem := range s.Features {
		s.Features[index] = strings.ToLower(elem)
	}
}

func contains(value string, slice []string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}
