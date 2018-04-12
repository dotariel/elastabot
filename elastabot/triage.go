package elastabot

// Triage abstracts the process of incident resolution.
type Triage struct {
	Topic   string
	Handler TriageHandler
}

type TriageHandler interface {
	StartTriage(string) error
}

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

func (t Triage) Execute() (string, error) {
	t.Handler.StartTriage(t.Topic)
	return "ok", nil
}
