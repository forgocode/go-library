package logger

import (
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//TODO: https://learnku.com/articles/70639 可参考

type logModule struct {
	Modules map[string]*NewLogger
	sync.RWMutex
}

var logModules = &logModule{
	Modules: make(map[string]*NewLogger),
}

var logLevel = []string{"info", "debug"}

func (m *logModule) add(name string, logger *NewLogger) {
	logModules.Lock()
	defer logModules.Unlock()
	m.Modules[name] = logger
}

type NewLogger struct {
	fileName string
	Level    string
	*zap.SugaredLogger
	zapLevel   zap.AtomicLevel
	moduleName string
}

type Options struct {
	FileName   string
	Level      string
	ModuleName string
}

func New(opts ...Options) *NewLogger {
	logger := &NewLogger{}
	for _, opt := range opts {
		logger.fileName = opt.FileName
		logger.Level = opt.Level
		logger.moduleName = opt.ModuleName
	}
	if len(opts) == 0 {
		logger.defaultOptions()
	}
	logger.zapLevel = getLevelEnabler()
	logger.SetLevel(logger.Level)
	logger.SugaredLogger = initLogger(logger.fileName, logger.zapLevel, logger.moduleName)
	logModules.add(logger.moduleName, logger)
	return logger
}

func HandleLog(args ...string) interface{} {
	if len(args) == 0 {
		return "\nplease input: mylog=[moduleName]=[loglevel]\n" +
			"example: mylog=atp=info\n" +
			logModules.printModule() +
			printLogLevel()
	}
	if len(args) > 2 {
		return "log args is more than 2"
	}
	if !logModules.isModuleExist(args[0]) {
		return logModules.printModule()
	}
	if len(args) == 1 {
		return args[0] + " logLevel is " + logModules.getModuleLogLevel(args[0])
	}
	logModules.setModuleLogLevel(args[0], args[1])
	return "module:" + args[0] + ", loglevel is setted: " + args[1]
}
func (m *logModule) printModule() string {
	m.RLock()
	defer m.RUnlock()
	result := ""
	for k := range m.Modules {
		result = result + ", " + k
	}
	return "moduleName: [" + result[1:] + " ]\n"
}

func printLogLevel() string {
	result := ""
	for _, v := range logLevel {
		result = result + ", " + v
	}
	return "logLevel: [" + result[1:] + " ]\n"
}

func (m *logModule) isModuleExist(module string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.Modules[module]
	return ok
}

func (m *logModule) setModuleLogLevel(module, logLevel string) {
	m.RLock()
	defer m.RUnlock()
	m.Modules[module].SetLevel(logLevel)
}

func (l *NewLogger) SetLevel(level string) {
	l.Level = level
	switch level {
	case "info":
		l.zapLevel.SetLevel(zapcore.InfoLevel)
	case "debug":
		l.zapLevel.SetLevel(zapcore.DebugLevel)
	}
}

func (m *logModule) getModuleLogLevel(moduleName string) string {
	m.RLock()
	defer m.RUnlock()
	return m.Modules[moduleName].Level
}

func (l *NewLogger) defaultOptions() {
	l.fileName = "running.log"
	l.Level = "info"
}

func initLogger(fileName string, level zapcore.LevelEnabler, moduleName string) *zap.SugaredLogger {
	writerSyncer := getLogWriter(fileName)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, level)
	loger := zap.New(core, zap.AddCaller())
	return loger.Sugar().Named(moduleName)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	logConf := zap.NewProductionEncoderConfig()
	logConf.EncodeTime = zapcore.ISO8601TimeEncoder
	logConf.EncodeLevel = zapcore.CapitalLevelEncoder
	logConf.NameKey = "module"
	logConf.SkipLineEnding = true
	return zapcore.NewConsoleEncoder(logConf)
	//return zapcore.NewJSONEncoder(logConf)
}

func getLevelEnabler() zap.AtomicLevel {
	return zap.NewAtomicLevel()
}
