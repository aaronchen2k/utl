package agentUntils

import (
	"fmt"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	agentRes "github.com/utlai/utl/res/agent"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

var (
	usageFile = fmt.Sprintf("res%sdoc%susage.txt", string(os.PathSeparator), string(os.PathSeparator))
)

func PrintUsage() {
	_logUtils.PrintColor("Usage: ", color.FgCyan)

	usage := ReadResData(usageFile)

	app := "utl-agent"
	if _commonUtils.IsWin() {
		app += ".exe"
	}
	usage = fmt.Sprintf(usage, app)
	fmt.Printf("%s\n", usage)
}

func ReadResData(path string) string {
	isRelease := _commonUtils.IsRelease()

	var jsonStr string
	if isRelease {
		data, _ := agentRes.Asset(path)
		jsonStr = string(data)
	} else {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			jsonStr = "fail to read " + path
		} else {
			str := string(buf)
			jsonStr = _commonUtils.RemoveBlankLine(str)
		}
	}

	return jsonStr
}
