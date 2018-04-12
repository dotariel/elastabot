package elastabot

// TriageHandler is anything that can initiate a triage event.
type TriageHandler interface {
	StartTriage(string) error
}

// Triage abstracts the process of incident resolution.
type Triage struct {
	Topic   string
	Handler TriageHandler // This is an interface
}

// NewTriage creates an instance of a Triage.
func NewTriage(topic string, handler TriageHandler) Triage {
	return Triage{Topic: topic, Handler: handler}
}

func (t Triage) Equals(other Command) bool {
	if b, ok := other.(Triage); ok {
		return t.Topic == b.Topic &&
			t.Handler == b.Handler
	}

	return false
}

// Execute is the Command interface entrypoint that begins the triage process.
func (t Triage) Execute() (string, error) {
	t.Handler.StartTriage(t.Topic)
	return "ok", nil
}
