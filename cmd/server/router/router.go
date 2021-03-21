package router

import (
	"fmt"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/utlai/utl/cmd/server/router/handler"
	"github.com/utlai/utl/internal/server/biz/middleware"
	middlewareUtils "github.com/utlai/utl/internal/server/biz/middleware/misc"
	"github.com/utlai/utl/internal/server/cfg"
	"github.com/utlai/utl/internal/server/repo"
	"github.com/utlai/utl/internal/server/service"
	"net/http"
)

type Router struct {
	api *iris.Application

	InitService   *service.InitService      `inject:""`
	JwtService    *middleware.JwtService    `inject:""`
	TokenService  *middleware.TokenService  `inject:""`
	CasbinService *middleware.CasbinService `inject:""`

	AccountCtrl *handler.AccountCtrl `inject:""`
	FileCtrl    *handler.FileCtrl    `inject:""`

	InitCtrl *handler.InitCtrl `inject:""`
	PermCtrl *handler.PermCtrl `inject:""`
	RoleCtrl *handler.RoleCtrl `inject:""`
	UserCtrl *handler.UserCtrl `inject:""`

	RpcCtrl *handler.RpcCtrl `inject:""`

	PlanCtrl *handler.ProjectCtrl `inject:""`

	WsCtrl *handler.WsCtrl `inject:""`

	TokenRepo *repo.TokenRepo `inject:""`
}

func NewRouter(app *iris.Application) *Router {
	router := &Router{}
	router.api = app

	return router
}

func (r *Router) App() {
	iris.LimitRequestBodySize(serverConf.Config.Options.UploadMaxSize)
	r.api.UseRouter(middlewareUtils.CrsAuth())

	app := r.api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		// 二进制模式 ， 启用项目入口
		if serverConf.Config.BinData {
			app.Get("/", func(ctx iris.Context) { // 首页模块
				_ = ctx.View("index.html")
			})
		}

		v1 := app.Party("/v1")
		{
			v1.PartyFunc("/rpc", func(party iris.Party) {
				party.Post("/request", r.RpcCtrl.Request).Name = "转发RPC请求"
			})

			v1.PartyFunc("/admin", func(admin iris.Party) {
				admin.Get("/init", r.InitCtrl.InitData)
				admin.Post("/login", r.AccountCtrl.UserLogin)

				//登录验证
				admin.Use(r.JwtService.Serve, r.TokenService.Serve, r.CasbinService.Serve)

				admin.Post("/logout", r.AccountCtrl.UserLogout).Name = "退出"
				admin.Get("/expire", r.AccountCtrl.UserExpire).Name = "刷新Token"
				admin.Get("/profile", r.UserCtrl.GetProfile).Name = "个人信息"

				admin.PartyFunc("/projects", func(party iris.Party) {
					party.Get("/", r.PlanCtrl.List).Name = "项目列表"
					party.Get("/{id:uint}", r.PlanCtrl.Get).Name = "项目详情"
					party.Post("/", r.PlanCtrl.Create).Name = "创建项目"
					party.Put("/{id:uint}", r.PlanCtrl.Update).Name = "更新项目"
					party.Delete("/{id:uint}", r.PlanCtrl.Delete).Name = "删除项目"
				})

				admin.PartyFunc("/users", func(party iris.Party) {
					party.Get("/", r.UserCtrl.GetAllUsers).Name = "用户列表"
					party.Get("/{id:uint}", r.UserCtrl.GetUser).Name = "用户详情"
					party.Post("/", r.UserCtrl.CreateUser).Name = "创建用户"
					party.Put("/{id:uint}", r.UserCtrl.UpdateUser).Name = "编辑用户"
					party.Delete("/{id:uint}", r.UserCtrl.DeleteUser).Name = "删除用户"
				})
				admin.PartyFunc("/roles", func(party iris.Party) {
					party.Get("/", r.RoleCtrl.GetAllRoles).Name = "角色列表"
					party.Get("/{id:uint}", r.RoleCtrl.GetRole).Name = "角色详情"
					party.Post("/", r.RoleCtrl.CreateRole).Name = "创建角色"
					party.Put("/{id:uint}", r.RoleCtrl.UpdateRole).Name = "编辑角色"
					party.Delete("/{id:uint}", r.RoleCtrl.DeleteRole).Name = "删除角色"
				})
				admin.PartyFunc("/permissions", func(party iris.Party) {
					party.Get("/", r.PermCtrl.GetAllPermissions).Name = "权限列表"
					party.Get("/{id:uint}", r.PermCtrl.GetPermission).Name = "权限详情"
					party.Post("/", r.PermCtrl.CreatePermission).Name = "创建权限"
					party.Put("/{id:uint}", r.PermCtrl.UpdatePermission).Name = "编辑权限"
					party.Delete("/{id:uint}", r.PermCtrl.DeletePermission).Name = "删除权限"
				})
			})
		}

		websocketAPI := r.api.Party("/api/v1/ws")
		m := mvc.New(websocketAPI)
		m.Register(
			&prefixedLogger{prefix: "DEV"},
		)
		m.HandleWebsocket(handler.NewWsCtrl())
		websocketServer := websocket.New(
			gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}),
			m)
		websocketAPI.Get("/", websocket.Handler(websocketServer))
	}
}

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}
