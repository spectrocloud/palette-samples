package internal

import (
	"fmt"
	"time"
)

// Format the cluster age using weeks, days, and hours
func FormatAge (clusterAge time.Duration) (int, int, int) {
    const hoursPerWeek int = 168
    weeks := int(clusterAge.Hours()) / hoursPerWeek
    remainingHours := int(clusterAge.Hours()) - weeks * hoursPerWeek
    days := remainingHours / 24
    hours := remainingHours % 24
    return  weeks, days, hours
}

// Return the cluster age as a formatted string
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