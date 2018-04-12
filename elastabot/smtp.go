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

func (smtp Smtp) StartTriage(string) error {
	return nil
}
