package logger

type logger struct{}

func New() *logger {
	return &logger{}
}