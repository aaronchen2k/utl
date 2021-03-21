package main

import (
	"flag"
	"github.com/utlai/utl/cmd/agent/router"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	"github.com/utlai/utl/internal/agent/cron"
	agentUntils "github.com/utlai/utl/internal/agent/utils/common"
	agentConst "github.com/utlai/utl/internal/agent/utils/const"
	_const "github.com/utlai/utl/internal/pkg/const"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	help     bool
	flagSet  *flag.FlagSet
	platform string
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("utl", flag.ContinueOnError)

	flagSet.StringVar(&agentConf.Inst.Server, "s", "", "")
	flagSet.StringVar(&agentConf.Inst.Ip, "i", "", "")
	flagSet.IntVar(&agentConf.Inst.Port, "p", 10, "")
	flagSet.StringVar(&agentConf.Inst.Language, "l", "zh", "")
	flagSet.StringVar(&platform, "t", string(_const.Vm), "")

	flagSet.BoolVar(&help, "h", false, "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	switch os.Args[1] {
	case "start":
		start()

	case "help", "-h":
		agentUntils.PrintUsage()

	default:
		start()
	}
}

func start() {
	_logUtils.Init(agentConst.AppName)
	agentConf.Init()

	if err := flagSet.Parse(os.Args[1:]); err == nil {
		agentConf.Inst.Platform = _const.WorkPlatform(platform)

		agentCron.Init()
		router.App()
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
