package wlog

// TYFileds type, used to pass to `WithFields`.
type TYFields map[string]interface{}

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Warning(msg string)
	Error(msg string)
}

type AdapterLogger interface {
	Info(level Level, msg string, field TYFields)
}

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARNING
	ERROR
)
