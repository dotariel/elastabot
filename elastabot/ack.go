package elastabot

type Ack struct {
	Alert    string
	Duration int
	Triage   *Triage
}

func NewAck(alert string, duration string, triage *Triage) *Ack {
	return &Ack{
		Alert:    alert,
		Duration: convertToMinutes(duration),
		Triage:   triage,
	}
}

func (a Ack) Equals(other Command) bool {
	if b, ok := other.(Ack); ok {
		return a.Alert == b.Alert &&
			a.Duration == b.Duration
		// a.Triage.Equals(b.Triage)
	}

	return false
}

func (a Ack) Execute() (string, error) {
	// Talk to ES...
	// Also handle or delegate to Triage.execute()
	return "", nil
}

func convertToMinutes(d string) (minutes int) {
	if len(d) == 0 {
		return
	}
	// if not str:
	//   mins = defaultMins
	// elif str.endswith('m'):
	//   mins = int(str[:-1])
	// elif str.endswith('h'):
	//   mins = int(str[:-1]) * 60
	// elif str.endswith('d'):
	//   mins = int(str[:-1]) * 1440
	// elif str.endswith('h'):
	//   mins = int(str[:-1]) * 10080
	// else:
	//   mins = int(str)
	return 0
}
