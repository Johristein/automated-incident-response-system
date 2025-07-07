# incident-response-system-regex-basedm
A modular incident response system written in Go to automatically detect and classify security threats using log message analysis and pattern matching.

# 🛡️ IRS-RB — Incident Response System Regex Based in Go

## 📌 Project Overview  
IRS-RB is a lightweight and modular security tool developed in Go for detecting and responding to threats such as **SQL injections**, **DDoS attacks**, **phishing attempts**, and **malware infections** using **regular expressions** and structured event classification. It simulates real-time detection, alerting, classification, response, and report generation — all in a clean Go architecture.

---

## 📂 Input Data (Device Logs)

- **Input Format**: CSV  
- **Structure**:
  | IP Address | Open Ports | Software | OS | Details | Message |
  |------------|------------|----------|----|---------|---------|

- **Example Entry**:
- 192.168.1.5,80,443,Apache,Linux,Trojan found in binary


---

## 🔍 Insights & Findings

- 🧠 Uses **regex patterns** to detect security incidents:
- `SQL Injection` → `(?i)select\s+.*\s+from|union\s+select|--|;`
- `Phishing` → `(?i)click here|reset password|account suspended`
- `DDoS` → `(?i)flood|overloaded|ddos|service unavailable`
- `Malware` → `(?i)trojan|virus|worm|malware|ransomware`

- 🗂️ Every incident is classified by:
- Type (`SQL Injection`, `Phishing`, etc.)
- Severity (`Low`, `Medium`, `High`)
- Status (`Active`, `Investigating`, `Resolved`)

- 🔔 Alerts are triggered for high-severity or unknown attacks  
- 📊 Generates summary reports of all detected events

---

## ✅ Key Features

- 🔍 **Pattern-based analysis** using Go’s `regexp` engine  
- 🧱 **Modular structure**: analysis, logging, report, response, scanner  
- 📤 **Dual-format logging**: `.txt` and `.json` output  
- 🔔 **Live alerts** based on severity and type  
- 🧾 **Summary reporting** at the end of execution  

---

## 📦 Project Structure
AIRS/
├── analysis/ → Threat classification logic

├── event/ → Event object structure & simulation

├── incident/ → Incident result schema

├── logging/ → Logger & alert system
├── logs/ → file to store logs

├── report/ → Summary report generation

├── response/ → Response simulation based on severity

├── scanner/ → CSV parser for events

└── main.go → Program entry point


---

## 🚀 How to Run

```bash
git clone https://github.com/yourname/AIRS.git
cd AIRS
go run main.go
```

---

## 🧠 System Flow
Read CSV Logs → Detect & Classify → Alert & Log → Respond → Summarize

## 🔍 Detection  
Pattern matching using Go `regexp`

## 📤 Logging  
Writes structured `.txt` and `.json` incident logs in `/logs`

## 🛠️ Response  
Handles incidents based on severity:  
- **Low** → log  
- **Medium** → notify  
- **High** → trigger response (simulated)

## 📊 Report  
Prints summary with:  
- Total incidents  
- Severity counts  
- Processing time

---
## 🧪 Sample Output
```
Analyzing event ID 2: Trojan found in binary
Event ID: 2, Type: Malware Infection, Severity: High, Status: Active
[ALERT] 🚨 Urgent incident detected!
ID: 2 | Type: Malware Infection | Severity: High | IP: 192.168.1.5

--- Incident Response Summary ---
Total Events Processed: 5
Low Severity Incidents: 1
Medium Severity Incidents: 2
High Severity Incidents: 2
Processing Duration: 3.42s
```

---

## 🔮 Future Enhancements

- 📊 Web-based dashboard for real-time display
- 🧠 Integrate LLM/AI model for anomaly classification

---

## 🛡️ License

This project is licensed under the [MIT License](./LICENSE).

---

## 👤 Author

**Johristein**  
GitHub: Johristein

---

> 🛡️ This project showcases how Go can be used to build an automated incident response system — combining real-time event processing, pattern-based threat detection, and structured logging for security operations and educational use.
