package main

import (
	"AIRS/analysis"
	"AIRS/event"
	"AIRS/logging"
	"AIRS/report"
	"AIRS/response"
	"AIRS/scanner"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	events, err := scanner.ReadCSV("incident_data.csv")
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		os.Exit(1)
	}

	scannedEvents := scanner.ScanThreats(events)

	lowSeverityCount, mediumSeverityCount, highSeverityCount := 0, 0, 0
	resultChannel := make(chan string, len(scannedEvents))

	var wg sync.WaitGroup

	for _, e := range scannedEvents {
		wg.Add(1)
		go func(e event.SecurityEvent) {
			defer wg.Done()
			event.ProcessEvent(e, resultChannel)
		}(e)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		fmt.Println(result)
	}

	for _, e := range scannedEvents {
		analyzedIncident := analysis.AnalyzeEvent(e)

		switch analyzedIncident.Severity {
		case "Low":
			lowSeverityCount++
		case "Medium":
			mediumSeverityCount++
		case "High":
			highSeverityCount++
		}

		if err := logging.LogIncident(analyzedIncident); err != nil {
			fmt.Println("Error logging incident:", err)
		}
		logging.SendAlert(analyzedIncident)

		response.HandleResponse(analyzedIncident.Severity, analyzedIncident.ID, analyzedIncident.Type)
	}

	report.GenerateReport(len(scannedEvents), lowSeverityCount, mediumSeverityCount, highSeverityCount, startTime)
}
