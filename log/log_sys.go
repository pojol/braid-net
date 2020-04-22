package log

import (
	"go.uber.org/zap"
)

/*
	诊断日志
	包含如
	1. 系统&进程状态日志
	2. 定期任务状态时间日志
	3. 程序追踪日志（警告，错误
	4. request慢日志
*/

// SysError 系统错误日志
func SysError(module string, function string, desc string) {
	logPtr.gSysLog.Error(desc, // err msg
		zap.String("module", module), // 模块
		zap.String("func", function), // 函数
	)
}

// SysSlow 慢日志
func SysSlow(apiName string, requestid string, et int, desc string) {

	logPtr.gSysLog.Warn(desc,
		zap.String("api", apiName),         // trace name
		zap.String("requestID", requestid), // trace id
		zap.Int("executionTime", et),       // 总计执行时间
	)

}

// SysRoutingError 路由警告日志
func SysRoutingError(serviceName string, desc string) {
	logPtr.gSysLog.Warn(desc,
		zap.String("service", serviceName),
	)
}

// SysWelcome 欢迎日志
func SysWelcome(box string, mode string, ty string, info string) {
	logPtr.gSysLog.Info(info,
		zap.String("box", box),
		zap.String("mode", mode),
		zap.String("type", ty),
	)
}