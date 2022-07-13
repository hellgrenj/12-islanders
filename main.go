package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultWeight = 100

var numberOfSeesawMeasurements = 0

type islander struct {
	name   string
	weight int
}
type simres struct {
	islander                   islander
	diff                       string
	numberOfSeesawMeasurements int
}

func main() {
	islanders := initIslanders()
	result := findOddIslander(islanders)
	fmt.Printf("found %s islander %v (weight: %v) with %v seesaw measurements\n", result.diff, result.islander.name, result.islander.weight, result.numberOfSeesawMeasurements)
}
func findOddIslander(islanders []islander) simres {
	fmt.Printf("the islanders are: %v\n", islanders)
	left := islanders[:4]       // L1 L2 L3 L4 (named after starting position)
	right := islanders[4:8]     // R1 R2 R3 R4 (named after starting position)
	sideline := islanders[8:12] // S1 S2 S3 S4 (named after starting position)
	result := seesaw(left, right)
	if result == "balanced" {
		return handleBalanced(left, right, sideline)
	} else if result == "left" {
		return handleLeftHeavy(left, right, sideline)
	} else { // right
		// make sure you have the heavy side on the left... i.e pass in right as left and left as right
		return handleLeftHeavy(right, left, sideline)
	}
}
func handleLeftHeavy(left, right, sideline []islander) simres {
	// seesaw: L1 L2 L3 L4 (heavy)      R1 R2 R3 R4
	// sideline: S1 S2 S3 S4
	// now switch R1 R2 R3 and S2 S3 S4 AND switch R1 AND L1
	left2 := []islander{right[0], left[1], left[2], left[3]}
	right2 := []islander{left[0], sideline[1], sideline[2], sideline[3]}
	sideline2 := []islander{sideline[0], right[1], right[2], right[3]}

	result2 := seesaw(left2, right2)
	if result2 == "left" { // left still heavy
		// we now know there is a heavy islander amongst L2 L3 and L4
		// seesaw: R1 *L2 L3 L4*    L1 S2 S3 S4
		// sideline: S1 R2 R3 R4
		left21 := []islander{left2[1]}
		right21 := []islander{left2[2]}
		// now compare L2 and L3
		result21 := seesaw(left21, right21)
		if result21 == "balanced" {
			return simres{left[3], "heavy", numberOfSeesawMeasurements} // L4
		} else if result21 == "left" {
			return simres{left21[0], "heavy", numberOfSeesawMeasurements} // L2
		} else {
			return simres{right21[0], "heavy", numberOfSeesawMeasurements} // L3
		}

	} else if result2 == "balanced" {
		// we now know that the light islander is amongst R2 R3 and R4
		// seesaw: R1 L2 L3 L4    L1 S2 S3 S4
		// sideline: S1 *R2 R3 R4*
		left22 := []islander{sideline2[1]}
		right22 := []islander{sideline2[2]}
		// now compare R2 and R3
		result22 := seesaw(left22, right22)
		if result22 == "balanced" {
			return simres{sideline2[3], "light", numberOfSeesawMeasurements} // R4
		} else if result22 == "left" {
			return simres{right22[0], "light", numberOfSeesawMeasurements} // R3
		} else {
			return simres{left22[0], "light", numberOfSeesawMeasurements} // R2
		}

	} else { // right side now heavy
		// we now know it is either R1 or L1
		// seesaw: *R1* L2 L3 L4    *L1* S2 S3 S4 (heavy)
		// sideline: S1 R2 R3 R4
		left23 := []islander{left2[0]}
		right23 := []islander{sideline2[0]} // known neutral islander ...
		// now compare R1 against a neutral islander (S1 for example)
		result23 := seesaw(left23, right23)
		if result23 == "balanced" {
			return simres{right2[0], "heavy", numberOfSeesawMeasurements} // L1
		} else {
			return simres{left23[0], "light", numberOfSeesawMeasurements} // R1
		}
	}
}
func handleBalanced(left, right, sideline []islander) simres {
	// seesaw: L1 L2 L3 L4   R1 R2 R3 R4 (balanced)
	// sideline: S1 S2 S3 S4
	// now compare S1 S2 S3 with R1 R2 R3
	left2 := sideline[:3]
	right2 := right[:3]
	result2 := seesaw(left2, right2)

	if result2 == "balanced" {
		// seesaw: S1 S2 S3  R1 R2 R3  (balanced)
		// sideline: L1 L2 L3 S4 L4 R4
		left21 := []islander{sideline[3]}
		right21 := []islander{left2[0]}
		// compare S4 with neutral L1
		result21 := seesaw(left21, right21)
		if result21 == "left" {
			return simres{sideline[3], "heavy", numberOfSeesawMeasurements} // heavy S4
		} else {
			return simres{sideline[3], "light", numberOfSeesawMeasurements} // light S4
		}
	} else if result2 == "left" {
		// seesaw: S1 S2 S3 (heavy)  R1 R2 R3
		// sideline: L1 L2 L3 S4 L4 R4
		left22 := []islander{left2[0]}
		right22 := []islander{left2[1]}
		// compare S1 with S2
		result22 := seesaw(left22, right22)
		if result22 == "balanced" {
			return simres{left2[2], "heavy", numberOfSeesawMeasurements} // S3
		} else if result22 == "left" {
			return simres{left22[0], "heavy", numberOfSeesawMeasurements} // S1
		} else {
			return simres{right22[0], "heavy", numberOfSeesawMeasurements} // S2
		}
	} else {
		// seesaw: S1 S2 S3  R1 R2 R3 (heavy)
		// sideline: L1 L2 L3 S4 L4 R4
		left23 := []islander{left2[0]}
		right23 := []islander{left2[1]}
		// compare S1 and S2
		result23 := seesaw(left23, right23)
		if result23 == "balanced" {
			return simres{left2[2], "light", numberOfSeesawMeasurements} // S3
		} else if result23 == "left" {
			return simres{right23[0], "light", numberOfSeesawMeasurements} // S2
		} else {
			return simres{left23[0], "light", numberOfSeesawMeasurements} // S1
		}
	}
}
func initIslanders() []islander {
	var islanders []islander
	for i := 0; i < 12; i++ {
		newIslander := islander{
			name:   fmt.Sprintf(("%d"), i+1),
			weight: defaultWeight,
		}
		islanders = append(islanders, newIslander)
	}
	// make a random islander randomly heavier or lighter
	islanders[getRandomNumber(0, len(islanders))].weight = getRandomNumber(80, 120)

	return islanders
}
func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	rndNumb := rand.Intn(max-min) + min
	if rndNumb == defaultWeight {
		rndNumb = getRandomNumber(min, max) // dont let it be the same as the default weight
	}
	return rndNumb
}

func seesaw(left, right []islander) string {
	numberOfSeesawMeasurements++
	leftTotaltWeight := 0
	for _, islander := range left {
		leftTotaltWeight += islander.weight
	}
	rightTotaltWeight := 0
	for _, islander := range right {
		rightTotaltWeight += islander.weight
	}
	if leftTotaltWeight == rightTotaltWeight {
		return "balanced"
	} else if leftTotaltWeight > rightTotaltWeight {
		return "left"
	} else {
		return "right"
	}
}
