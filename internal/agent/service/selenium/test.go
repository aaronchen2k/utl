package seleniumService

import (
	"fmt"
	commonService "github.com/utlai/utl/internal/agent/service/common"
	execService "github.com/utlai/utl/internal/agent/service/exec"
	agentConst "github.com/utlai/utl/internal/agent/utils/const"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	"path/filepath"
)

func ExecTest(build *_domain.BuildTo) {
	result := _domain.RpcResult{}

	commonService.SetBuildWorkDir(build)

	// get script
	execService.GetTestScript(build)
	if build.ProjectDir == "" {
		result.Fail(fmt.Sprintf("failed to get test script, %#v。", build))
		return
	}

	// exec test
	parseBuildCommand(build)
	result = execService.ExcCommand(build)
	if !result.IsSuccess() {
		result.Fail(fmt.Sprintf("failed to ext test,\n dir: %s\n  cmd: \n%s",
			build.ProjectDir, build.BuildCommands))
	}

	// submit result
	execService.UploadResult(*build, result)
}

func parseBuildCommand(build *_domain.BuildTo) {
	// mvn clean test -Dtestng.suite=target/test-classes/baidu-test.xml
	//		 		  -DdriverPath=${driverPath}
	dir := ""
	if _commonUtils.IsWin() {
		dir = agentConst.ResPathWin
	} else {
		dir = agentConst.ResPathLinux
	}
	driverFolder := filepath.Join(dir, agentConst.BrowserDriverDir, string(build.BrowserType))
	driverFile := fmt.Sprintf("%s-%s", _commonUtils.GetOs(), build.BrowserVer)
	if _commonUtils.IsWin() {
		driverFile += ".exe"
	}
	build.BuildCommands = build.BuildCommands + " -DdriverPath=" + filepath.Join(driverFolder, driverFile)
}
