package core

import (
	"fmt"
	"gomaster/common"
	"gomaster/config"
	infra "gomaster/initialize"
	"gomaster/utils"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	oplogging "github.com/op/go-logging"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "gomaster"
)

var (
	//defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
	defaultFormatter = `%{time:2006/01/02 - 15:04:05} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

type LogStarter struct {
	infra.BaseStarter
}

func (l *LogStarter) Init(ctx infra.StarterContext) {
	initLog()
}

func initLog() {
	c := common.Conf().Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := oplogging.MustGetLogger(module)
	var backends []oplogging.Backend
	registerStdout(c, &backends)
	if fileWriter := registerFile(c, &backends); fileWriter != nil {
		// 如果需要将日志同时写入文件和控制台，请使用以下代码
		gin.DefaultWriter = io.MultiWriter(fileWriter, os.Stdout)
	}
	oplogging.SetBackend(backends...)
	common.Set_LOG(logger)
}

func registerStdout(c config.Log, backends *[]oplogging.Backend) {
	if c.Stdout != "" {
		level, err := oplogging.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		*backends = append(*backends, createBackend(os.Stdout, c, level))
	}
}

func registerFile(c config.Log, backends *[]oplogging.Backend) io.Writer {
	if c.File != "" {
		if ok, _ := utils.PathExists(logDir); !ok {
			// directory not exist
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		fileWriter, err := rotatelogs.New(
			logDir+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
			// generate soft link, point to latest log file
			rotatelogs.WithLinkName(logSoftLink),
			// maximum time to save log files
			rotatelogs.WithMaxAge(7*24*time.Hour),
			// time period of log file switching
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
		}
		level, err := oplogging.LogLevel(c.File)
		if err != nil {
			fmt.Println(err)
		}
		*backends = append(*backends, createBackend(fileWriter, c, level))
		return fileWriter
	}

	return nil
}

func createBackend(w io.Writer, c config.Log, level oplogging.Level) oplogging.Backend {
	backend := oplogging.NewLogBackend(w, c.Prefix, 0)
	stdoutWriter := false
	if w == os.Stdout {
		stdoutWriter = true
	}
	format := getLogFormatter(c, stdoutWriter)
	backendLeveled := oplogging.AddModuleLevel(oplogging.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled
}

func getLogFormatter(c config.Log, stdoutWriter bool) oplogging.Formatter {
	pattern := defaultFormatter
	if !stdoutWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		// Remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return oplogging.MustStringFormatter(pattern)
}
