package wlog

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Elk struct {
	*Entry
	// 主机名
	// Host string `json:"host"`
	// // 调用方id
	// ServerID string `json:"server_id"`
	// // 路由
	// Route string `json:"route"`
	// // 时间戳
	// Timestamp int64 `json:"time_stamp"`
	// // 日期
	// Date string `json:"date"`
	// // 请求id
	// RequestID string `json:"request_id"`
	// // ua
	// UserAgent string `json:"user_agent"`
	// // 项目ID
	// Namespace string `json:"name_space"`
	// // appId
	// AppID string `json:"app_id"`
	// // package
	// Package   string `json:"package"`
	// UserID    string `json:"user_id"`
	// EventID   string `json:"event_id"`
	// ErrorCode string `json:"error_code"`
	// ErrorMsg  string `json:"error_msg"`
	// // 元数据
	// Metadata map[string]string `json:"metadata"`
	// // 对象
	// Object interface{} `json:"object"`
}

func Host(string)                {}
func ServerID(string)            {}
func Route(string)               {}
func Timestamp(int64)            {}
func Date(string)                {}
func RequestID(string)           {}
func UserAgent(string)           {}
func NameSpace(string)           {}
func AppID(string)               {}
func Package(string)             {}
func UserID(string)              {}
func EventID(string)             {}
func ErrorCode(string)           {}
func ErrorMsg(string)            {}
func Metadata(map[string]string) {}
func Object(interface{})         {}

// MustAppendContext  context
func (elk *Entry) MustAppendContext(ctx context.Context) *Entry {
	return elk
}

// MustAppendFromGinContext 使用gin context
func (elk *Entry) MustAppendFromGinContext(ctx *gin.Context) *Entry {
	elk.withFields(elk.getGinContext(ctx))
	return elk
}

// getGinContext gin使用预定义参数
func (elk *Entry) getGinContext(ctx *gin.Context) TYFields {
	data := make(TYFields, 6)
	request := ctx.Request
	data["user-Agent"] = request.Header.Get("User-Agent")
	data["ip"] = ctx.ClientIP()
	data["path"] = request.URL.Path
	data["method"] = request.Method
	data["proto"] = request.Proto
	return data
}

// EventId 事件id
func (elk *Entry) EventId(value string) *Entry {
	elk.withField("event_id", value)
	return elk
}

// ErrorMsg 错误信息
func (elk *Entry) ErrorMsg(value string) *Entry {
	elk.withField("err_msg", value)
	return elk
}

// ErrorReport 错误上报
func (elk *Entry) ErrorReport(value string) *Entry {
	elk.withField("err_report", value)
	return elk
}

// SubEventId subeventId
func (elk *Entry) SubEventId(value string) *Entry {
	elk.withField("sub_eventId", value)
	return elk
}

// EventParam 时间参数
func (elk *Entry) EventParam(params interface{}) *Entry {
	elk.withField("event_Params", params)
	return elk
}

// Rparam 时间参数
func (elk *Entry) Rparam(params interface{}) *Entry {
	elk.withField("r_param", params)
	return elk
}

// NameSpace 项目ID
func (elk *Entry) NameSpace(value string) *Entry {
	elk.withField("name_space", value)
	return elk
}

// NameSpace 项目ID
func (elk *Entry) UseID(value string) *Entry {
	elk.withField("user_id", value)
	return elk
}

// Metadata
func (elk *Entry) Metadata(data map[string]string) *Entry {
	elk.withField("metadata", data)
	return elk
}
