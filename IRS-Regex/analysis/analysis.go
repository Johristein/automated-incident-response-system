package analysis

import (
	"AIRS/event"
	"AIRS/incident"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// AnalyzeEvent dynamically analyzes the event and classifies it into specific attack types
func AnalyzeEvent(e event.SecurityEvent) incident.Incident {
	analyzedEvent := incident.Incident{
		ID:              e.ID,
		Message:         e.Message,
		IP:              e.IP,
		OpenPorts:       e.OpenPorts,
		Software:        e.Software,
		OperatingSystem: e.OperatingSystem,
		Details:         e.Details,
	}

	fmt.Printf("Analyzing event ID %d: %s\n", analyzedEvent.ID, e.Message)

	switch {
	case isSQLInjection(e.Message):
		analyzedEvent.Type = "SQL Injection"
		analyzedEvent.Severity = "High"
		analyzedEvent.Status = "Active"
	case isPhishingAttempt(e.Message):
		analyzedEvent.Type = "Phishing"
		analyzedEvent.Severity = "Medium"
		analyzedEvent.Status = "Investigating"
	case isDDosAttack(e.Message):
		analyzedEvent.Type = "DDoS Attack"
		analyzedEvent.Severity = "High"
		analyzedEvent.Status = "Active"
	case isMalwareDetected(e.Message):
		analyzedEvent.Type = "Malware Infection"
		analyzedEvent.Severity = randomSeverity([]string{"Low", "Medium", "High"})
		analyzedEvent.Status = determineStatus(analyzedEvent.Severity)
	default:
		analyzedEvent.Type = "Security Anomaly"
		analyzedEvent.Severity = "Medium"
		analyzedEvent.Status = "Investigating"
	}

	fmt.Printf("Event ID: %d, Type: %s, Severity: %s, Status: %s\n",
		analyzedEvent.ID, analyzedEvent.Type, analyzedEvent.Severity, analyzedEvent.Status)

	return analyzedEvent
}

func isSQLInjection(message string) bool {
	regex := `(?i)select\s+.*\s+from|union\s+select|--|;`
	return regexp.MustCompile(regex).MatchString(message)
}

func isPhishingAttempt(message string) bool {
	regex := `(?i)click\s+here|reset\s+password|account\s+suspended`
	return regexp.MustCompile(regex).MatchString(message)
}

func isDDosAttack(message string) bool {
	regex := `(?i)flood|overloaded|ddos|service\s+unavailable`
	return regexp.MustCompile(regex).MatchString(message)
}

func isMalwareDetected(message string) bool {
	regex := `(?i)trojan|virus|worm|malware|ransomware`
	return regexp.MustCompile(regex).MatchString(message)
}

func randomSeverity(choices []string) string {
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(len(choices))]
}

func determineStatus(severity string) string {
	switch severity {
	case "High":
		return "Active"
	case "Medium":
		return "Investigating"
	case "Low":
		return "Resolved"
	default:
		return "Unknown"
	}
}
