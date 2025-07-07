package report

import (
	"fmt"
	"time"
)

// GenerateReport will generate a summary report of all processed incidents
func GenerateReport(totalEvents int, lowSeverityCount int, mediumSeverityCount int, highSeverityCount int, startTime time.Time) {
	duration := time.Since(startTime)

	fmt.Printf("\n--- Incident Response Summary ---\n")
	fmt.Printf("Total Events Processed: %d\n", totalEvents)
	fmt.Printf("Low Severity Incidents: %d\n", lowSeverityCount)
	fmt.Printf("Medium Severity Incidents: %d\n", mediumSeverityCount)
	fmt.Printf("High Severity Incidents: %d\n", highSeverityCount)
	fmt.Printf("Processing Duration: %v\n", duration)
}
