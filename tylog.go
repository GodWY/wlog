package wlog

import (
	"sync"

	"github.com/google/uuid"
)

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

// log tylog日志
type Log struct {
	adapter   AdapterLogger
	msg       string
	entryPool sync.Pool
}

// NewLogger 创建新的日志组件
func NewLogger(adapterLogger AdapterLogger, msg string) *Log {
	return &Log{adapter: adapterLogger, msg: msg}
}

// NewEntry creates a new entry
func (log *Log) WithEntry(msg string) *Entry {
	entry := log.newEntry(msg)
	// defer log.releaseEntry(entry)
	entry.Data["trace_id"] = uuid.New()
	entry.entryMsg = msg
	return entry
}

// NewEntry creates a elk entry
func (log *Log) WithElkEntry(msg string) *Entry {
	entry := log.newEntry(msg)
	entry.Data["trace_id"] = uuid.New()
	entry.entryMsg = msg
	return entry
}

func (log *Log) newEntry(msg string) *Entry {
	entry, ok := log.entryPool.Get().(*Entry)
	if ok && entry != nil {
		return entry
	}
	entry = &Entry{
		Logger:   log,
		Data:     make(TYFields, 6),
		entryMsg: msg,
		msg:      log.msg,
	}
	return entry
}

func (log *Log) releaseEntry(entry *Entry) {
	entry.Data = map[string]interface{}{}
	log.entryPool.Put(entry)
}

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARNING
	ERROR
)
