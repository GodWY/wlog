package wlog

import (
	"errors"
)

var ErrorAdapter = errors.New("please set log adapter")

type Entry struct {
	Data     TYFields
	Logger   *Log
	msg      string
	entryMsg string
}

// WithField 使用结构化日志
func (entry *Entry) withField(key string, value interface{}) *Entry {
	// 重新初始化一个日志
	filed := make(TYFields, 2)
	filed[key] = value
	return entry.withFields(filed)
}

// WithFields 使用结构化日志
func (entry *Entry) withFields(fields TYFields) *Entry {
	// 重新初始化一个日志
	data := make(TYFields, len(entry.Data)+2)
	for key, value := range entry.Data {
		data[key] = value
	}
	// 组合key
	for key, value := range fields {
		data[key] = value
	}
	entry.Data = data
	return entry
}

// Info 打印普通日志
func (entry *Entry) Info(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	field := make(TYFields, 3)
	field["trace_id"] = entry.Data["trace_id"]
	field["func"] = entry.entryMsg
	field["service"] = entry.msg
	entry.Logger.adapter.Info(INFO, msg, field)
}

// Debug 打印调试日志
func (entry *Entry) Debug(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	field := make(TYFields, 3)
	field["trace_id"] = entry.Data["trace_id"]
	field["func"] = entry.entryMsg
	field["service"] = entry.msg
	entry.Logger.adapter.Info(DEBUG, msg, field)

}

// Warning 打印警告日志
func (entry *Entry) Warning(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	field := make(TYFields, 3)
	field["trace_id"] = entry.Data["trace_id"]
	field["func"] = entry.entryMsg
	field["service"] = entry.msg
	entry.Logger.adapter.Info(WARNING, msg, field)
}

// Error 打印错误日志
func (entry *Entry) Error(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	field := make(TYFields, 3)
	field["trace_id"] = entry.Data["trace_id"]
	field["func"] = entry.entryMsg
	field["service"] = entry.msg
	entry.Logger.adapter.Info(ERROR, msg, field)
}

// Flush 输出日志
func (entry *Entry) Flush() *Entry {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	defer entry.Logger.releaseEntry(entry)
	entry.Logger.adapter.Info(INFO, "flush", entry.Data)
	return entry
}

// MustAppend 预定义字段
func (entry *Entry) MustAppend() *Entry {
	return entry
}

// Int64 类型
func (entry *Entry) Int64(key string, value int64) *Entry {
	return entry.withField(key, value)
}

// Int 类型
func (entry *Entry) Int(key string, value int) *Entry {
	return entry.withField(key, value)
}

// String 类型
func (entry *Entry) String(key string, value string) *Entry {
	return entry.withField(key, value)
}

// Int32 值类型的数据
func (entry *Entry) Int32(key string, value int32) *Entry {
	return entry.withField(key, value)
}

// Object 类数据
func (entry *Entry) Object(key string, value interface{}) *Entry {
	return entry.withField(key, value)
}
