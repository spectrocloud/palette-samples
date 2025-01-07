package internal

import (
	"fmt"
	"time"
)

type FormattedAge struct {
	Days, Weeks, Hours int
}

// FormatAge returns the age of a cluster as the number of weeks, days, and hours.
func FormatAge(clusterAge time.Duration) (*FormattedAge, error) {
	const hoursPerWeek int = 168
	hoursAge := int(clusterAge.Hours())
	if hoursAge < 0 {
		return nil, fmt.Errorf("%v is less than zero hours", hoursAge)
	}
	weeks := hoursAge / hoursPerWeek
	remainingHours := hoursAge - weeks*hoursPerWeek
	days := remainingHours / 24
	hours := remainingHours % 24
	return &FormattedAge{
		Days:  days,
		Weeks: weeks,
		Hours: hours,
	}, nil
}

// GetFormattedAge returns a formatted string representation of the cluster's age.
// It uses FormatAge to calculate the age in weeks, days, and hours and formats
// the output accordingly.
func GetFormattedAge(clusterAge time.Duration) (*string, error) {
	fa, err := FormatAge(clusterAge)
	if err != nil {
		return nil, err
	}
	var formattedString string
	if fa.Weeks > 0 {
		formattedString = fmt.Sprintf("%dw", fa.Weeks)
	}
	if fa.Days > 0 {
		if len(formattedString) > 0 {
			formattedString += " "
		}
		formattedString = fmt.Sprintf("%s%dd", formattedString, fa.Days)
	}
	if fa.Hours > 0 {
		if len(formattedString) > 0 {
			formattedString += " "
		}
		formattedString = fmt.Sprintf("%s%dh", formattedString, fa.Hours)
	}
	return &formattedString, nil
}