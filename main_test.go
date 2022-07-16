package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFindOddIslander(t *testing.T) {

	t.Logf("\n cases when one islander is light")
	for i := 0; i < 12; i++ { // test all variants of lighter
		islanders := getTestIslanders(i, 80)
		result := findOddIslander(islanders)
		t.Logf("found %s islander %v (weight: %v) with %v seesaw measurements\n", result.diff, result.islander.name, result.islander.weight, result.numberOfSeesawMeasurements)
		if result.diff != "light" {
			t.Errorf("Expected result.diff to be 'light' got '%v'", result.diff)
		}
		if result.islander.name != strconv.Itoa(i+1) {
			t.Errorf("Expected result.islander.name to be %s got '%v'", strconv.Itoa(i+1), result.islander.name)
		}
		if result.islander.weight != 80 {
			t.Errorf("Expected result.islander.weight to be 80 got '%v'", result.islander.weight)
		}
		if result.numberOfSeesawMeasurements > 3 {
			t.Errorf("Expected result.numberOfSeesawMeasurements to be 3 or less got '%v'", result.numberOfSeesawMeasurements)
		}
	}
	t.Logf("\n\n cases when one islander is heavy")
	for i := 0; i < 12; i++ { // test all variants of heavier
		islanders := getTestIslanders(i, 120)
		result := findOddIslander(islanders)
		t.Logf("found %s islander %v (weight: %v) with %v seesaw measurements\n", result.diff, result.islander.name, result.islander.weight, result.numberOfSeesawMeasurements)
		if result.diff != "heavy" {
			t.Errorf("Expected result.diff to be 'heavy' got '%v'", result.diff)
		}
		if result.islander.name != strconv.Itoa(i+1) {
			t.Errorf("Expected result.islander.name to be '%s' got '%v'", strconv.Itoa(i+1), result.islander.name)
		}
		if result.islander.weight != 120 {
			t.Errorf("Expected result.islander.weight to be 120 got '%v'", result.islander.weight)
		}
		if result.numberOfSeesawMeasurements > 3 {
			t.Errorf("Expected result.numberOfSeesawMeasurements to be 3 or less got '%v'", result.numberOfSeesawMeasurements)
		}
	}
}
func getTestIslanders(indexOddOne int, weight int) []islander {
	var islanders []islander
	for i := 0; i < 12; i++ {
		newIslander := islander{
			name:   fmt.Sprintf(("%d"), i+1),
			weight: defaultWeight,
		}
		islanders = append(islanders, newIslander)
	}
	islanders[indexOddOne].weight = weight

	return islanders
}
