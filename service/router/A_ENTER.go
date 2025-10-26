package router

import (
	"sun-panel/global"
	// "sun-panel/router/admin"
	"sun-panel/router/openness"
	"sun-panel/router/panel"
	"sun-panel/router/system"

	"github.com/gin-gonic/gin"
)

// 初始化总路由
func InitRouters(addr string) error {
	router := gin.Default()
	// rootRouter := router.Group("/")   // 不必须
	// 创建两个前缀的 API 分组：/api 与 /sunpanel/api
	apiGroup := router.Group("/api")
	apiGroupSun := router.Group("/sunpanel/api")

	// 接口：对每个模块同时注册到 /api 和 /sunpanel/api
	system.Init(apiGroup)
	system.Init(apiGroupSun)

	panel.Init(apiGroup)
	panel.Init(apiGroupSun)

	openness.Init(apiGroup)
	openness.Init(apiGroupSun)

	// 将根路径重定向到 /sunpanel
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/sunpanel")
	})

	// WEB文件服务（/sunpanel 前缀）
	{
		webPath := "./web"
		router.StaticFile("/sunpanel", webPath+"/index.html")
		router.Static("/sunpanel/assets", webPath+"/assets")
		router.Static("/sunpanel/custom", webPath+"/custom")
		router.StaticFile("/sunpanel/favicon.ico", webPath+"/favicon.ico")
		router.StaticFile("/sunpanel/favicon.svg", webPath+"/favicon.svg")
	}

	// 上传的文件
	sourcePath := global.Config.GetValueString("base", "source_path")
	router.Static(sourcePath[1:], sourcePath)

	global.Logger.Info("Sun-Panel is Started.  Listening and serving HTTP on ", addr)
	return router.Run(addr)
}
