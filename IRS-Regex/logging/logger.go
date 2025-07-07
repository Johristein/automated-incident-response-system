package logging

import (
	"AIRS/incident"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func LogIncident(i incident.Incident) error {
	// Ensure log directory exists
	logDir := "logs"
	os.MkdirAll(logDir, os.ModePerm)

	// Get today's date
	dateStr := time.Now().Format("2006-01-02")

	// Text log
	textLogPath := filepath.Join(logDir, fmt.Sprintf("incident_log_%s.txt", dateStr))
	logFile, err := os.OpenFile(textLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	logLine := fmt.Sprintf("[%s] ID: %d, Type: %s, Severity: %s, Status: %s, IP: %s, Ports: %v, Software: %s, OS: %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		i.ID, i.Type, i.Severity, i.Status, i.IP, i.OpenPorts, i.Software, i.OperatingSystem)

	_, err = logFile.WriteString(logLine)
	if err != nil {
		return err
	}

	// JSON log
	jsonLogPath := filepath.Join(logDir, fmt.Sprintf("incident_log_%s.json", dateStr))
	jsonFile, err := os.OpenFile(jsonLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := json.Marshal(i)
	if err != nil {
		return err
	}

	_, err = jsonFile.WriteString(string(jsonData) + "\n")
	return err
}

func SendAlert(i incident.Incident) {
	if i.Severity == "High" || i.Type == "Unknown Attack" {
		fmt.Printf("[ALERT] 🚨 Urgent incident detected!\n")
		fmt.Printf("ID: %d | Type: %s | Severity: %s | IP: %s\n", i.ID, i.Type, i.Severity, i.IP)
		// Simulate email alert or webhook
	}
}
