package internal

import (
	"fmt"
	"time"
)

// Format the cluster age using weeks, days, and hours
func FormatAge (clusterAge time.Duration) (int, int, int) {
    const hoursPerWeek = 168
    weeks := int(clusterAge.Hours() / hoursPerWeek)
    remainingHours := int(clusterAge.Hours() - float64((weeks * hoursPerWeek)))
    days := int(remainingHours / 24)
    hours := int(remainingHours % 24)
    return  weeks, days, hours
}

// Return the cluster age as a formatted string
func PrintFormattedAge (clusterAge time.Duration) (string) {
    weeks, days, hours := FormatAge(clusterAge)
    formattedString := ""
    if weeks > 0 {
        formattedString += fmt.Sprint(weeks) + " weeks "
    }
    if days > 0 {
        formattedString += fmt.Sprint(days) + " days "
    }
    if hours > 0 {
        formattedString += fmt.Sprint(hours) + " hours"
    }
    return formattedString
}