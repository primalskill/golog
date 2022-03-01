package golog

// NoopLog implements a no operation logger used in tests.
type NoopLog struct {}

func NewNoopLog() *NoopLog {
	return &NoopLog{}
}

func (p *NoopLog) Info(msg string, meta ...Meta) {
	return
}

func (p *NoopLog) Warn(msg string, meta ...Meta) {
	return
}

func (p *NoopLog) Error(err error, meta ...Meta) {
	return
}