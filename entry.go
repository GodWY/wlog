package wlog

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ErrorAdapter = errors.New("please set log adapter")

type Entry struct {
	Data   TYFields
	Logger AdapterLogger
	msg    string
}

// NewEntry 新建日志模块
func NewEntry(logger AdapterLogger, msg string) *Entry {
	return &Entry{Data: make(TYFields, 6), Logger: logger, msg: msg}
}

// WithField 使用结构化日志
func (entry *Entry) withField(key string, value interface{}) *Entry {
	// 重新初始化一个日志
	filed := make(TYFields, 1)
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
	entry.Logger.Info(INFO, msg, entry.Data)
}

// Debug 打印调试日志
func (entry *Entry) Debug(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	entry.Logger.Info(DEBUG, msg, entry.Data)
}

// Warning 打印警告日志
func (entry *Entry) Warning(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	entry.Logger.Info(WARNING, msg, entry.Data)
}

// Error 打印错误日志
func (entry *Entry) Error(msg string) {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	entry.Logger.Info(ERROR, msg, entry.Data)
}

// Flush 输出日志
func (entry *Entry) Flush() *Entry {
	if entry.Logger == nil {
		panic(ErrorAdapter)
	}
	entry.Logger.Info(INFO, entry.msg, entry.Data)
	entry.Data = map[string]interface{}{}
	return entry
}

// MustAppend 预定义字段
func (entry *Entry) MustAppend() *Entry {
	data := make(TYFields, len(entry.Data))
	entry.Data = data
	field := make(TYFields, 6)
	field["trace_id"] = uuid.New()
	return entry.withFields(field)
}

// MustAppendContext  context
func (entry *Entry) MustAppendContext(ctx context.Context) *Entry {
	data := make(TYFields, len(entry.Data))
	entry.Data = data
	field := make(TYFields, 6)

	field["trace_id"] = uuid.New()
	return entry.withFields(field)
}

// MustAppendFromGinContext 使用gin context
func (entry *Entry) MustAppendFromGinContext(ctx *gin.Context) *Entry {
	data := make(TYFields, len(entry.Data))
	entry.Data = data
	field := make(TYFields, 6)
	field["trace_id"] = uuid.New()
	entry.getContext(ctx)
	return entry.withFields(field)
}

// MustContext gin使用预定义参数
func (entry *Entry) getContext(ctx *gin.Context) *Entry {
	data := make(TYFields, 6)
	request := ctx.Request
	data["user-Agent"] = request.Header.Get("User-Agent")
	data["ip"] = ctx.ClientIP()
	data["path"] = request.URL.Path
	data["method"] = request.Method
	data["proto"] = request.Proto
	//data["status"] = request.Response.Status
	return entry.withFields(data)
}

// EventId 事件id
func (entry *Entry) EventId(value string) *Entry {
	return entry.withField("event_id", value)
}

// ErrorMsg 错误信息
func (entry *Entry) ErrorMsg(value string) *Entry {
	return entry.withField("err_msg", value)
}

// ErrorReport 错误上报
func (entry *Entry) ErrorReport(value string) *Entry {
	return entry.withField("err_report", value)
}

// SubEventId subeventId
func (entry *Entry) SubEventId(value string) *Entry {
	return entry.withField("sub_eventId", value)
}

// EventParam 时间参数
func (entry *Entry) EventParam(params interface{}) *Entry {
	return entry.withField("event_Params", params)
}

// Rparam 时间参数
func (entry *Entry) Rparam(params interface{}) *Entry {
	return entry.withField("r_param", params)
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
