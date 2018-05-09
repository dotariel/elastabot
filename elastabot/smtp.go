package elastabot

type Smtp struct {
	Host           string
	Port           int
	Secure         bool
	StartTLS       bool
	TimeoutSeconds int
	To             string
	From           string
	SubjectPrefix  string
	Debug          bool
}

// StartTriage sends an email to the configured endpoint
func (smtp Smtp) StartTriage(string) error {
	return nil
}
