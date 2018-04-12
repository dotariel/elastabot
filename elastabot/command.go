package elastabot

type Command interface {
	Execute() (string, error)
	Equals(Command) bool
}
