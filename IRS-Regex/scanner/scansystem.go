package scanner

import (
	"AIRS/event"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func ReadCSV(filename string) ([]event.SecurityEvent, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	events := []event.SecurityEvent{}
	for i, record := range records[1:] {
		if len(record) < 6 {
			continue
		}

		ip := strings.TrimSpace(record[0])
		openPorts := strings.Split(record[1], ",")
		var ports []int
		for _, portStr := range openPorts {
			port, err := strconv.Atoi(strings.TrimSpace(portStr))
			if err == nil {
				ports = append(ports, port)
			}
		}

		e := event.SecurityEvent{
			ID:              i + 1,
			IP:              ip,
			OpenPorts:       ports,
			Software:        strings.TrimSpace(record[2]),
			OperatingSystem: strings.TrimSpace(record[3]),
			Details:         strings.TrimSpace(record[4]),
			Message:         strings.TrimSpace(record[5]),
		}
		events = append(events, e)
	}
	return events, nil
}

func ScanThreats(devices []event.SecurityEvent) []event.SecurityEvent {
	return devices
}
