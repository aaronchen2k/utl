package router

import (
	"fmt"
	"github.com/utlai/utl/cmd/agent/router/handler"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	_const "github.com/utlai/utl/internal/pkg/const"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/smallnest/rpcx/server"
	"strconv"
)

func App() {
	addr := agentConf.Inst.Ip + ":" + strconv.Itoa(agentConf.Inst.Port)

	srv := server.NewServer()

	if agentConf.Inst.Platform == _const.Vm {
		srv.RegisterName("selenium", new(handler.SeleniumAction), "")

	} else if agentConf.Inst.Platform == _const.Box {
		srv.RegisterName("appium", new(handler.AppiumAction), "")

	}

	_logUtils.Info(fmt.Sprintf("start server on %s ...", addr))
	err := srv.Serve("tcp", addr)
	if err != nil {
		_logUtils.Infof("fail to start server on %s, err is %s", err.Error())
	}
}
