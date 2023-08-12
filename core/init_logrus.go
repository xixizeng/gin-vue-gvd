package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogRequest struct {
	LogPath  string //日志目录
	AppName  string //app的名字
	NoDate   bool   //是否关闭按照时间分割
	NoErr    bool   //是否关闭需要单独存放的BUG等级
	NoGlobal bool   //是否关闭替换全局
}

type LogFormatter struct {
}

type ErrorHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

// 对error级别的错误文件进行单独输出
func (hook ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook ErrorHook) Fire(entry *logrus.Entry) (err error) {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	//判断日期是否为当前日期
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		fmt.Println("这是一次>>>>>>>>>>>>>>")
		return nil
	}
	hook.file.Close()
	//重新创建文件目录
	os.MkdirAll(path.Join(hook.logPath, timer), os.ModeAppend)
	filePath := path.Join(hook.logPath, timer, "err.log")
	hook.file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	hook.fileDate = timer
	hook.file.Write([]byte(line))

	return nil
}

// 按时间分割写入日志文件的hook
type DateHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

func (hook DateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook DateHook) Fire(entry *logrus.Entry) (err error) {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	//判断日期是否为当前日期
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	hook.file.Close()
	//重新创建文件目录
	os.MkdirAll(path.Join(hook.logPath, timer), os.ModeAppend)
	filePath := path.Join(hook.logPath, timer, hook.appName+".log")

	hook.file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	hook.fileDate = timer
	hook.file.Write([]byte(line))

	return nil
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

var log *logrus.Logger

// 创建一个logrus
func NewLog(requestList ...LogRequest) *logrus.Logger {
	var request LogRequest
	if len(requestList) > 0 {
		request = requestList[0]
	}
	if request.LogPath == "" {
		request.LogPath = "logs"
	}
	if request.AppName == "" {
		request.AppName = "gvd"
	}
	mLog := logrus.New()               //新建一个实例
	mLog.SetOutput(os.Stdout)          //设置输出类型
	mLog.SetReportCaller(true)         //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)   //设置最低的Level

	if !request.NoDate {

		mLog.AddHook(&DateHook{
			logPath: request.LogPath,
			appName: request.AppName,
		}) //加载日志按时间分割的输出的方式
	}
	if !request.NoErr {
		mLog.AddHook(&ErrorHook{
			logPath: request.LogPath,
			appName: request.AppName,
		}) //加载日志单独输出ERROR等级的方式
	}
	if !request.NoGlobal {
		InitDefaultLogger()
	}
	return mLog
}

// 定义全局的logrus
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)          //设置输出类型
	logrus.SetReportCaller(true)         //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	logrus.SetLevel(logrus.DebugLevel)   //设置最低的Level
	logrus.AddHook(&DateHook{})          //选择日志输出的方式
	logrus.AddHook(&ErrorHook{})         //将error另外加载写在下面，每天第一个错误是error就不会报错
}
