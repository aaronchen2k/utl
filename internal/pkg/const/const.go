package _const

import "os"

const (
	AppName = "utl"
	PthSep  = string(os.PathSeparator)

	LanguageDefault = "en"
	LanguageEN      = "en"
	LanguageZH      = "zh"

	UserTokenExpireTime = 365 * 24 * 60 * 60 * 1000

	RegisterExpireTime = 5  // min
	WaitForExecTime    = 60 // min
	VmTimeout          = 30 // min

	RetryTime    = 3
	AgentRunTime = 20 // sec

	WebCheckQueueInterval    = 10 // sec
	AgentCheckDeviceInterval = 10 // sec

	MaxVmOnHost = 3
	RpcPort     = 8972

	UploadDir          = "uploads"
	Sep_Of_Mac_Address = ":"

	Build_Command_Param_AppPath        = "${appPath}"
	Build_Command_Param_AppPackage     = "${appPackage}"
	Build_Command_Param_AppActivity    = "${appActivity}"
	Build_Command_Param_AppiumPort     = "${appiumPort}"
	Build_Command_Param_SeleniumDriver = "${driverPath}"

	PageSize = 10
)
