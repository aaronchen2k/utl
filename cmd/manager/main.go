package main

import (
	"flag"
	"github.com/utlai/utl/cmd/manager/program"
	managerConf "github.com/utlai/utl/cmd/manager/utils/conf"
	managerConst "github.com/utlai/utl/cmd/manager/utils/const"
	agentConst "github.com/utlai/utl/internal/agent/utils/const"
	_const "github.com/utlai/utl/internal/pkg/const"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_stringUtils "github.com/utlai/utl/internal/pkg/libs/string"
	"github.com/kardianos/service"
	"os"
	"os/signal"
	"syscall"
)

var (
	help    bool
	flagSet *flag.FlagSet

	action string
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

	flagSet.StringVar(&action, "a", "", "")

	if action != "" && !_stringUtils.StrInArr(action, agentConst.ControlActions) {
		_logUtils.Warn(_i118Utils.I118Prt.Sprintf("invalid_actions", action, service.ControlAction))
		return
	}

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	config := &service.Config{
		Name:        _const.AppName,
		DisplayName: _const.AppName,
		Description: _const.AppName + " service.",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}

	prg := &program.Program{}
	srv, err := service.New(prg, config)
	if err != nil {
		_logUtils.Error(err.Error())
		os.Exit(1)
	}
	errs := make(chan error, 5)
	program.Logger, err = srv.Logger(errs)
	if err != nil {
		_logUtils.Error(err.Error())
		os.Exit(1)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				_logUtils.Info(err.Error())
			}
		}
	}()

	if action != "" {
		err := service.Control(srv, action)
		if err != nil {
			_logUtils.Info(_i118Utils.I118Prt.Sprintf("invalid_actions", action, service.ControlAction))

			_logUtils.Error(err.Error())
			os.Exit(1)
		}
		return
	}

	err = srv.Run()
	if err != nil {
		program.Logger.Error(err)
	}
}

func init() {
	_logUtils.Init(managerConst.AppName)
	managerConf.Init()

	cleanup()
}

func cleanup() {
}
