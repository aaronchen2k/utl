package _const

type ResultCode int

const (
	ResultSuccess ResultCode = 1
	ResultFail    ResultCode = 0
)

func (c ResultCode) Int() int {
	return int(c)
}

type BuildProgress string

const (
	ProgressCreated    BuildProgress = "created"
	ProgressLaunchVm   BuildProgress = "launch_vm"
	ProgressPending    BuildProgress = "pending"
	ProgressInProgress BuildProgress = "in_progress"
	ProgressTimeout    BuildProgress = "timeout"
	ProgressCompleted  BuildProgress = "completed"
)

type BuildStatus string

const (
	StatusCreated BuildStatus = "created"
	StatusPass    BuildStatus = "pass"
	StatusFail    BuildStatus = "fail"

	//Unresponsive BuildStatus = "unresponsive"
	//AppiumInvalid BuildStatus = "appium_invalid"
	//DeviceInvalid BuildStatus = "device_invalid"
)

type VmStatus string

const (
	VmCreated       VmStatus = "created"
	VmLaunched      VmStatus = "launched"
	VmActive        VmStatus = "active"
	VmBusy          VmStatus = "busy"
	VmDestroy       VmStatus = "destroy"
	VmFailToCreate  VmStatus = "fail_to_create"
	VmFailToDestroy VmStatus = "fail_to_destroy"
)

type HostStatus string

const (
	HostActive  HostStatus = "active"
	HostOffline HostStatus = "offline"
)

type DeviceStatus string

const (
	DeviceOff     DeviceStatus = "off"
	DeviceOn      DeviceStatus = "on"
	DeviceActive  DeviceStatus = "active"
	DeviceBusy    DeviceStatus = "busy"
	DeviceUnknown DeviceStatus = "unknown"
)

type ServiceStatus string

const (
	ServiceOff    ServiceStatus = "off"
	ServiceOn     ServiceStatus = "on"
	ServiceActive ServiceStatus = "active"
	ServiceBusy   ServiceStatus = "busy"
)

type WorkPlatform string

const (
	Host WorkPlatform = "host"
	Vm   WorkPlatform = "vm"
	Box  WorkPlatform = "box"

	Android WorkPlatform = "android"
	Ios     WorkPlatform = "ios"
)

type BuildType string

const (
	AppiumTest   BuildType = "appium_test"
	SeleniumTest BuildType = "selenium_test"
	UnitTest     BuildType = "unit_test"
)

type OsPlatform string

const (
	OsWindows OsPlatform = "windows"
	OsLinux   OsPlatform = "linux"
	OsMac     OsPlatform = "mac"

	OsAndroid OsPlatform = "android"
	OsIos     OsPlatform = "ios"
)

type OsType string

const (
	Win10 OsType = "win10"
	Win7  OsType = "win7"
	WinXP OsType = "winxp"

	Ubuntu OsType = "ubuntu"
	CentOS OsType = "centos"
	Debian OsType = "debian"

	Mac OsType = "mac"
)

type SysLang string

const (
	EN_US SysLang = "en_us"
	ZH_CN SysLang = "zh_cn"
	ZH_TW SysLang = "zh_tw"
	ZH_HK SysLang = "zh_hk"
	DE_DE SysLang = "de_de"
	JA_JP SysLang = "ja_jp"
	FR_FR SysLang = "fr_fr"
	RU_RU SysLang = "ru_ru"
)

type BrowserType string

const (
	Chrome  BrowserType = "chrome"
	Firefox BrowserType = "firefox"
	Edge    BrowserType = "edge"
	IE      BrowserType = "ie"
)

type ResType string

const (
	ResRoot      ResType = "root"
	ResCluster   ResType = "cluster"
	ResComputer  ResType = "computer"
	ResVm        ResType = "vm"
	ResContainer ResType = "container"
	ResImage     ResType = "image"
	ResFolder    ResType = "folder"
)
