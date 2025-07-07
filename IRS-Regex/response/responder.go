package response

import (
	"fmt"
	"sync"
)

func HandleResponse(severity string, incidentID int, eventType string) {
	var wg sync.WaitGroup

	switch severity {
	case "Low":
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("[WARNING] Incident ID %d: %s event needs investigation.\n", incidentID, eventType)
		}()
	case "Medium":
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("[NOTIFY] Admin notified for Incident ID %d: %s (Medium severity).\n", incidentID, eventType)
		}()
	case "High":
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("[ACTION] Countermeasures deployed for Incident ID %d: %s (High severity).\n", incidentID, eventType)
		}()
	}

	wg.Wait()
}
