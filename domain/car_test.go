package domain

import (
	"fmt"
	"testing"
)

var car Car

func TestCarSuccess(t *testing.T) {
	var (
		AllowedColors   = []string{"red", "blue", "green"}
		AllowedType     = []string{"sedan", "van", "suv", "motor-bike"}
		AllowedFeatures = []string{"panorama", "surround-system", "auto-parking", "sunroof"}
	)

	car = Car{
		Name:     "Test Car",
		Type:     "sedan",
		Features: AllowedFeatures,
	}

	for _, color := range AllowedColors {
		car.Color = color
		assertCarErr(t, false)
	}

	car.Type = ""

	for _, carType := range AllowedType {
		car.Type = carType
		assertCarErr(t, false)
	}
}

func TestCarCheckFailure(t *testing.T) {
	var (
		InvalidColors      = []string{"white", "brown", "purple", "yellow"}
		InvalidType        = []string{"jetski", "airplane", "boat", "ship", "meow", "bicycle"}
		InvalidSpeedRanges = []int{-1, 250, 280, 300, -100}
	)

	/* Empty Fields Test
	######################
	*/
	car = Car{} //Test Empty Car Object

	assertCarErr(t, true)

	car.Name = "Test Car" // Set Name To Allow For Test Of Type

	assertCarErr(t, true)

	car.Type = "InvalidType" // Set Car Type To Allow For Color Test

	assertCarErr(t, true)

	car.Color = "InvalidColor" // Set Car Color To Allow For Features Test

	assertCarErr(t, true)
	car.Features = []string{"Invalid", "Set", "Of", "Features"}

	//	assertCarErr(t, true)

	/* Test For Invalid Values
	#############################
	*/

	// Test For Invalid Car Types
	for _, types := range InvalidType {
		car.Type = types
		assertCarErr(t, true)
	}

	car.Type = "sedan" // Set Valid Car Type Value To Allow For Invalid Car Color Test

	// Test For Invalid Car Colors
	for _, color := range InvalidColors {
		car.Color = color
		assertCarErr(t, true)
	}

	car.Color = "red" // Set Valid Car Color Value To Allow For Invalid Speed Range Test

	// Test For Invalid SpeedRange Values

	for _, speedRange := range InvalidSpeedRanges {
		car.SpeedRange = speedRange
		assertCarErr(t, true)
	}

	car.SpeedRange = 200 // Set Valid Speed Range To Allow For Features Test
	// Test For Invalid Car Features Values

	assertCarErr(t, true)

}

func assertCarErr(t *testing.T, isError bool) {
	err := car.Check()

	if isError {
		if err == nil {
			t.Errorf("Expected An Error, Got: %v\n", err)
		}

	} else {

		if err != nil {
			fmt.Println(err)
			t.Errorf("Expected : nil, Got: %v\n", err)
		}
	}
}
