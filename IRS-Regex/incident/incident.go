package incident

type Incident struct {
	ID              int
	Message         string
	IP              string
	OpenPorts       []int
	Software        string
	OperatingSystem string
	Severity        string
	Type            string
	Status          string
	Details         string
}
