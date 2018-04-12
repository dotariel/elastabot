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

// Execute is the Command interface entrypoint that begins the triage process.
func (t Triage) Execute() (string, error) {
	t.Handler.StartTriage(t.Topic)
	return "ok", nil
}
