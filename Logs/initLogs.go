package Logs

// "github.com/chroblert/JCRandomProxy-GUI/Conf"

var applog = new(FishLogger)

// 20201011: 使用Logs
func InitLogs(logpath string, amaxSize int64, amaxAge, alogCount int) {
	maxSize = amaxSize // 单个文件最大大小
	maxAge = amaxAge   // 单个文件保存2天
	logCount = alogCount
	applog = NewLogger(logpath)
	defer applog.Flush()
	applog.SetLevel(DEBUG)
	applog.SetCallInfo(true)
	applog.SetConsole(true)
	//applog.Info("test")
}
func Println(args ...interface{}) {
	// applog.Info(args)
	applog.println(INFO, args...)
}

func Printf(format string, args ...interface{}) {
	// applog.Infof(format, args...)
	applog.printf(INFO, format, args...)
}
