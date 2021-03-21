package checkService

import (
	agentConf "github.com/utlai/utl/internal/agent/conf"
	appiumService "github.com/utlai/utl/internal/agent/service/appium"
	deviceService "github.com/utlai/utl/internal/agent/service/device"
	testService "github.com/utlai/utl/internal/agent/service/test"
	vmService "github.com/utlai/utl/internal/agent/service/vm"
)

func Check() {
	if agentConf.IsVmAgent() { // vm
		// is running，register busy
		if testService.CheckTaskRunning() {
			vmService.RegisterVm(true)
			return
		}

		// no task to run, submit free
		if testService.GetTaskSize() == 0 {
			vmService.RegisterVm(false)
			return
		}

		// has task to run，register busy, then run
		build := testService.PeekTask()
		vmService.RegisterVm(true)
		testService.Exec(build)

	} else if agentConf.IsDeviceAgent() { // device
		devices := deviceService.RefreshDevices()
		appiumService.CheckService(devices)

		// task is running，submit busy
		if testService.CheckTaskRunning() {
			deviceService.Register(devices, true)
			return
		}

		// no tasks to run, submit dev status
		if testService.GetTaskSize() == 0 {
			deviceService.Register(devices, false)
			return
		}

		// has task to run，register busy, then run
		build := testService.PeekTask()
		if deviceService.IsValid(devices, build.Serial) { // dev ready, lock, exec and submit status
			deviceService.Register(devices, true)

			testService.Exec(build)
		} else { // not ready，submit dev status
			deviceService.Register(devices, false)
		}

	}
}
