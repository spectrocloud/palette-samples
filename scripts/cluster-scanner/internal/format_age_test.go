package internal_test

import (
	"errors"
	"testing"
	"time"

	"github.com/spectrocloud/palette-samples/cluster-scanner/internal"
)

type TestCase struct {
	input                                      string
	expectedHours, expectedDays, expectedWeeks int
	expectedError                              error
	output                                     string
}

func TestFormatAge (t *testing.T) {
	tc := map[string]TestCase{
		"positive duration":{
			input:"2h",
			expectedHours:2,
		},
		"duration with weeks": {
			input: "336h",
			expectedWeeks:2,
		},
		"duration with days": {
			input: "48h",
			expectedDays:2,
		},
		"duration with hours and days": {
			input: "74h",
			expectedDays:3,
			expectedHours:2,
		},
		"duration with hours, days, and weeks": {
			input: "914h",
			expectedDays:3,
			expectedHours:2,
			expectedWeeks:5,
		},
		"zero duration":{
			input:  "0h",
			output: "",
		},
		"negative duration":{
			input:"-2h",
			expectedError: errors.New("-2 is less than zero hours"),
		},
	}

	for key, value := range tc {
		t.Run(key,  func(t *testing.T){
			duration, err := time.ParseDuration(value.input)
			if err != nil {
				t.Errorf("Error parsing duration: %v", err)
			}
			 
			fa, err := internal.FormatAge(duration)

			if value.expectedError != nil && err == nil {
				t.Errorf("Expected an error, but got none")
			}
			if value.expectedError == nil && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if value.expectedError != nil && err != nil {
				if value.expectedError.Error() != err.Error() {
					t.Errorf("Errors do not match: got %v, want %v", err, value.expectedError.Error())
				}
			}
			if fa != nil {
				if fa.Weeks != value.expectedWeeks {
					t.Errorf("got %d weeks, want %d weeks", fa.Weeks, value.expectedWeeks)
				}
				if fa.Days != value.expectedDays {
					t.Errorf("got %d days, want %d days", fa.Days, value.expectedDays)
				}
				if fa.Hours != value.expectedHours {
					t.Errorf("got %d hours, want %d hours", fa.Hours, value.expectedHours)
				}
			}
		})
	}
}

func TestGetFormattedAge(t *testing.T) {
	tc := map[string]TestCase{
		"duration with hours": {
			input: "2h",
			output: "2h",
		},
		"duration with weeks": {
			input: "336h",
			output: "2w",
		},
		"duration with days": {
			input: "48h",
			output: "2d",
		},
		"duration with hours and days": {
			input: "74h",
			output: "3d 2h",
		},
		"duration with hours, days, and weeks": {
			input: "914h",
			output: "5w 3d 2h",
		},
		"negative duration": {
			input: "-2h",
			expectedError: errors.New("-2 is less than zero hours"),
		},
		"zero duration": {
			input:  "0h",
			output: "",
		},
	}

	for key, value := range tc {
		t.Run(key, func(t *testing.T) {
			duration, err := time.ParseDuration(value.input)
			if err != nil {
				t.Errorf("Error parsing duration: %v", err)
			}

			faString, err := internal.GetFormattedAge(duration)

			if value.expectedError != nil && err == nil {
				t.Errorf("Expected an error, but got none")
			}
			if value.expectedError == nil && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if value.expectedError != nil && err != nil {
				if value.expectedError.Error() != err.Error() {
					t.Errorf("Errors do not match: got %v, want %v", err, value.expectedError.Error())
				}
			}
			if faString != nil {
				if *faString != value.output {
					t.Errorf("got '%s', want '%s'", *faString, value.output)
				}
			}
		})
	}

}
