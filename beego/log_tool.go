package beego

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo"

	"github.com/liangdas/mqant/log"
)

// Logger 日志对象
type Logger struct {
	trace      log.TraceSpan
	userAgent  string
	ip         string
	path       string
	method     string
	proto      string
	status     int
	accessId   string
	signType   string
	secretType string
	random     string
	timestamp  string
	signature  string
	//session
	server_id  string
	topic      string
	user_id    string
	session_id string
	client_id  string
	device_id  string
	msg        string
	hostName   string
}
type ContextParam struct {
	IP         string `json:"ip"`
	UserAgent  string `json:"user_agent"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	Proto      string `json:"proto"`
	Status     int    `json:"status"`
	AccessId   string `json:"access_id"`
	SignType   string `json:"sign_type"`
	SecretType string `json:"secret_type"`
	Random     string `json:"random"`
	Timestamp  string `json:"timestamp"`
	Signature  string `json:"signature"`

	//session
	ServerId  string `json:"server_id"`
	Topic     string `json:"topic"`
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
	ClientId  string `json:"client_id" valid:"required"` //服务项目的clientid
	DeviceId  string `json:"device_id" valid:"optional"` //用户的设备original_deviceid
}

// Newlogger创建一个logger
func Newlogger(msg string) *Logger {
	return &Logger{
		msg:       msg,
		trace:     log.CreateRootTrace(),
		timestamp: time.Now().Local().Format("2006-01-02 15:04:05"),
		hostName:  "",
	}
}

// Debugf 调试格式化打印
func (logger *Logger) Debugf(format string, a ...interface{}) {
	log.TDebug(logger.trace, format, a...)
}

// Errorf 错误格式化打印
func (logger *Logger) Errorf(format string, a ...interface{}) {
	log.TError(logger.trace, format, a...)
}

// Warnf 警告格式化打印
func (logger *Logger) Warnf(format string, a ...interface{}) {
	log.TWarning(logger.trace, format, a...)
}

// Infof 重要信息格式化打印
func (logger *Logger) Infof(format string, a ...interface{}) {
	log.TInfo(logger.trace, format, a...)
}

// GetLogger 获取 Logger 对象
func GetLogger(ctx echo.Context) *Logger {
	return &Logger{
		trace:      ctx.Get("trace").(log.TraceSpan),
		userAgent:  ctx.Request().Header.Get("User-Agent"),
		ip:         ctx.RealIP(),
		path:       ctx.Request().URL.Path,
		method:     ctx.Request().Method,
		proto:      ctx.Request().Proto,
		status:     ctx.Response().Status,
		accessId:   ctx.Request().Header.Get("Access-Id"),
		signType:   ctx.Request().Header.Get("Sign-Type"),
		secretType: ctx.Request().Header.Get("Secret-Type"),
		random:     ctx.Request().Header.Get("Random"),
		timestamp:  ctx.Request().Header.Get("Timestamp"),
		signature:  ctx.Request().Header.Get("Signature"),
	}
}

func GetRootLogger() *Logger {
	return &Logger{
		trace: log.CreateRootTrace(),
	}
}

func CreateLogger(trace, span string) *Logger {
	return &Logger{
		trace: log.CreateTrace(trace, span),
	}
}

type LogMsg struct {
	EventId     string
	SubEventId  string
	ErrorMsg    string
	ErrorReport string
	EventParams interface{}
	Rparam      interface{}
}

type BIMsg struct {
	EventId     string //事件ID
	ProjectId   string `json:"project_id" valid:"required"` //bi分配的项目组项目id
	ClientId    string `json:"client_id" valid:"required"`  //服务项目的clientid
	DeviceId    string `json:"device_id" valid:"optional"`  //用户的设备original_deviceid
	UserId      string `json:"user_id" valid:"optional"`    //用户id
	Type        string `json:"type" valid:"required"`       //ga事件类型。有game、coin、login、pay、push、track、profile
	ErrorCode   string `json:"err_code" valid:"optional"`   //错误码 如果正常事件 ErrorCode=0
	ErrorMsg    string `json:"err_msg" valid:"optional"`    //错误信息
	ErrorReport string //错误描述，由程序员书写，概述错误可能导
	// 致的问题以及解决方案
	Properties map[string]string `json:"properties"  valid:"optional"` //事件的额外补充，必须是string:string格式的map
	Lib        map[string]string `json:"lib" valid:"optional"`         //服务器版本信息,也必须是string:string格式的map
}

// Debugf 调试格式化打印
func (logger *Logger) JDebug(msg LogMsg) {
	log.TDebug(logger.trace, fmt.Sprintf("@%v", msg.EventId), msg.SubEventId, msg.ErrorMsg, msg.ErrorReport, msg.EventParams, msg.Rparam, logger.getContextParam())
}

// Errorf 错误格式化打印
func (logger *Logger) JError(msg LogMsg) {
	log.TError(logger.trace, fmt.Sprintf("@%v", msg.EventId), msg.SubEventId, msg.ErrorMsg, msg.ErrorReport, msg.EventParams, msg.Rparam, logger.getContextParam())
}

// Warnf 警告格式化打印
func (logger *Logger) JWarn(msg LogMsg) {
	log.TWarning(logger.trace, fmt.Sprintf("@%v", msg.EventId), msg.SubEventId, msg.ErrorMsg, msg.ErrorReport, msg.EventParams, msg.Rparam, logger.getContextParam())
}

// Infof 重要信息格式化打印
func (logger *Logger) JInfo(msg LogMsg) {
	log.TInfo(logger.trace, fmt.Sprintf("@%v", msg.EventId), msg.SubEventId, msg.ErrorMsg, msg.ErrorReport, msg.EventParams, msg.Rparam, logger.getContextParam())
}

func (logger *Logger) BiReport(msg BIMsg) {
	defer func() {
		if r := recover(); r != nil {
			logger.JError(LogMsg{
				EventId:    "bi_error",
				SubEventId: "catch",
				ErrorMsg:   fmt.Sprintf("%v", r),
			})
		}
	}()
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	baseMp := map[string]interface{}{
		"project_id": msg.ProjectId,
		"event_id":   msg.EventId,
		"type":       msg.Type,
		"event_time": time.Now().UnixNano() / 1000000, //毫秒
		"err_code":   msg.ErrorCode,
		"err_msg":    msg.ErrorMsg,
		"err_report": msg.ErrorReport,
		// "user_id":    IfString(msg.UserId != "", msg.UserId, logger.user_id),
		// "device_id":  IfString(msg.DeviceId != "", msg.DeviceId, logger.device_id),
		// "client_id":  IfString(msg.ClientId != "", msg.ClientId, logger.client_id),

		"ctx_ip":         logger.ip,
		"ctx_userAgent":  logger.userAgent,
		"ctx_path":       logger.path,
		"ctx_method":     logger.method,
		"ctx_proto":      logger.proto,
		"ctx_status":     logger.status,
		"ctx_accessId":   logger.accessId,
		"ctx_signType":   logger.signType,
		"ctx_secretType": logger.secretType,
		"ctx_random":     logger.random,
		"ctx_timestamp":  logger.timestamp,
		"ctx_signature":  logger.signature,
		"ctx_server_id":  logger.server_id,
		"ctx_session_id": logger.session_id,
		"ctx_topic":      logger.topic,
		"ctx_pid":        Int2String(pid),
		"ctx_hostname":   hostname,
	}
	if logger.trace != nil {
		baseMp["trace_id"] = logger.trace.TraceId()
		baseMp["trace_span"] = logger.trace.ExtractSpan()
	}
	if msg.Properties != nil {
		for k, v := range msg.Properties {
			if !strings.HasPrefix(k, "proj_") {
				baseMp[fmt.Sprintf("proj_%v", k)] = v
			}
		}
	}
	if msg.Lib != nil {
		for k, v := range msg.Lib {
			if !strings.HasPrefix(k, "lib_") {
				baseMp[fmt.Sprintf("lib_%v", k)] = v
			}
		}
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	msgbys, err := json.Marshal(baseMp)
	if err != nil {
		logger.JError(LogMsg{
			EventId:    "bi_error",
			SubEventId: "marshal",
			ErrorMsg:   err.Error(),
		})
		return
	}
	log.BiReport(string(msgbys))
}

func (logger *Logger) getContextParam() ContextParam {
	cp := &ContextParam{}
	cp.IP = logger.ip
	cp.UserAgent = logger.userAgent
	cp.Path = logger.path
	cp.Method = logger.method
	cp.Proto = logger.proto
	cp.Status = logger.status
	cp.AccessId = logger.accessId
	cp.SignType = logger.signType
	cp.SecretType = logger.secretType
	cp.Random = logger.random
	cp.Timestamp = logger.timestamp
	cp.Signature = logger.signature

	cp.UserId = logger.user_id
	cp.ServerId = logger.server_id
	cp.SessionId = logger.session_id
	cp.Topic = logger.topic
	cp.ClientId = logger.client_id
	cp.DeviceId = logger.device_id
	return *cp
}

func Int2String(a int) string {
	return strconv.Itoa(a)
}

func IfString() {}
