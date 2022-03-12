package beego

import (
	"encoding/json"
	"github.com/GodWY/wlog"
	logs "github.com/GodWY/wlog/lib/beego"
	"github.com/google/uuid"
)

type BeegoLogger struct {
	Option  *Options
	log     *logs.BeeLogger
	Trace   *Trace
	Msg     string
	Level   wlog.Level
	gaField map[string]string
}

type Trace struct {
	*logs.BeegoTraceSpan
}

// New creates a new
func New(opts *Options) *BeegoLogger {
	return &BeegoLogger{
		Option: opts,
		log:    CreateBeegoLogger(opts),
		Trace:  &Trace{},
	}
}

// NewGa 初始化Ga逻辑
func NewGa(opts *Options) *BeegoLogger {
	return &BeegoLogger{
		Option:  opts,
		log:     CreateBeegoLogger(opts),
		gaField: map[string]string{},
		//Trace:  &Trace{},
	}
}

// MustInstallGa Ga日志必选参数
func (bee *BeegoLogger) MustInstallGa(projectID, clientID string) {
	bee.gaField["project_id"] = projectID
	bee.gaField["client_id"] = clientID
}

// WithFilter 过滤字段
func (bee *BeegoLogger) WithFilter() {
}

// Info 输出日志格式
func (bee *BeegoLogger) Info(level wlog.Level, msg string, data wlog.TYFields) {
	bee.Trace.BeegoTraceSpan = &logs.BeegoTraceSpan{
		Trace: uuid.New().String(),
	}
	// 如果设置的level大于传入的level则不进行打印
	if bee.Level > level {
		return
	}
	switch level {
	case wlog.DEBUG:
		bee.log.Debugf(bee.Trace.BeegoTraceSpan, msg, data)
	case wlog.INFO:
		bee.log.Infof(bee.Trace.BeegoTraceSpan, msg, data)
	case wlog.WARNING:
		bee.log.Warnf(bee.Trace.BeegoTraceSpan, msg, data)
	case wlog.ERROR:
		bee.log.Errorf(bee.Trace.BeegoTraceSpan, msg, data)
	}
}

// SetLevel 设置日志输出级别
func (bee *BeegoLogger) SetLevel(level wlog.Level) {
	bee.Level = level
}

func CreateBeegoLogger(opt *Options) *logs.BeeLogger {
	if opt == nil {
		opt = NewOptions()
	}
	log := logs.NewLogger()
	log.SetLogFuncCallDepth(4)
	if opt.Debug {
		//控制台
		log.SetLogger(logs.AdapterConsole)
		return log
	}
	log.SetContentType("application/json")
	config, err := json.Marshal(opt)
	if err != nil {
		logs.Error(err)
		return nil
	}
	log.SetLogger(logs.AdapterMultiFile, string(config))
	return log
}
