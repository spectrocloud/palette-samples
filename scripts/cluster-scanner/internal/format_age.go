package internal

import (
	"fmt"
	"time"
)

// FormatAge returns the age of a cluster as the number of weeks, days, and hours.
func FormatAge (clusterAge time.Duration) (int, int, int) {
    const hoursPerWeek int = 168
    weeks := int(clusterAge.Hours()) / hoursPerWeek
    remainingHours := int(clusterAge.Hours()) - weeks * hoursPerWeek
    days := remainingHours / 24
    hours := remainingHours % 24
    return  weeks, days, hours
}

// PrintFormattedAge returns a formatted string representation of the cluster's age.
// It uses FormatAge to calculate the age in weeks, days, and hours and formats
// the output accordingly.
func PrintFormattedAge (clusterAge time.Duration) (string) {
    weeks, days, hours := FormatAge(clusterAge)
    var formattedString string
    if weeks > 0 {
       formattedString = fmt.Sprintf("%d weeks ", weeks)
    }
    if days > 0 {
        formattedString = fmt.Sprintf("%s %d days ", formattedString, days)
    }
    if hours > 0 {
        formattedString = fmt.Sprintf("%s %d hours", formattedString, hours)
    }
    return formattedString
}