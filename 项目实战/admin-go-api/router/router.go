// 访问接口路由配置

package router

import (
	"admin-api/api/controller"
	"admin-api/common/config"
	"admin-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 跌机恢复
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	router.Use(middleware.Logger())
	register(router)
	return router
}

// 路由注册
func register(router *gin.Engine) {
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controller.Login)
	// jwt鉴权接口
	jwt := router.Group("/api", middleware.AuthMiddleware(), middleware.LogMiddleware())
	{
		jwt.POST("/post/add", controller.CreateSysPost)
		jwt.GET("/post/list", controller.GetSysPostList)
		jwt.GET("/post/info", controller.GetSysPostById)
		jwt.PUT("/post/update", controller.UpdateSysPost)
		jwt.DELETE("/post/delete", controller.DeleteSysPostById)
		jwt.DELETE("/post/batch/delete", controller.BatchDeleteSysPost)
		jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/post/vo/list", controller.QuerySysPostVoList)
		jwt.GET("/dept/list", controller.GetSysDeptList)
		jwt.POST("/dept/add", controller.CreateSysDept)
		jwt.GET("/dept/info", controller.GetSysDeptById)
		jwt.PUT("/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/dept/vo/list", controller.QuerySysDeptVoList)
		jwt.POST("/menu/add", controller.CreateSysMenu)
		jwt.GET("/menu/vo/list", controller.QuerySysMenuVoList)
		jwt.GET("/menu/info", controller.GetSysMenu)
		jwt.PUT("/menu/update", controller.UpdateSysMenu)
		jwt.DELETE("/menu/delete", controller.DeleteSysMenu)
		jwt.GET("/menu/list", controller.GetSysMenuList)
		jwt.POST("/role/add", controller.CreateSysRole)
		jwt.GET("/role/info", controller.GetSysRoleById)
		jwt.PUT("/role/update", controller.UpdateSysRole)
		jwt.DELETE("/role/delete", controller.DeleteSysRoleById)
		jwt.PUT("/role/updateStatus", controller.UpdateSysRoleStatus)
		jwt.GET("/role/list", controller.GetSysRoleList)
		jwt.GET("/role/vo/list", controller.QuerySysRoleVoList)
		jwt.GET("/role/vo/idList", controller.QueryRoleMenuIdList)
		jwt.PUT("/role/assignPermissions", controller.AssignPermissions)
		jwt.POST("/admin/add", controller.CreateSysAdmin)
		jwt.GET("/admin/info", controller.GetSysAdminInfo)
		jwt.PUT("/admin/update", controller.UpdateSysAdmin)
		jwt.DELETE("/admin/delete", controller.DeleteSysAdminById)
		jwt.PUT("/admin/updateStatus", controller.UpdateSysAdminStatus)
		jwt.PUT("/admin/updatePassword", controller.ResetSysAdminPassword)
		jwt.GET("/admin/list", controller.GetSysAdminList)
		jwt.POST("/upload", controller.Upload)
		jwt.PUT("/admin/updatePersonal", controller.UpdatePersonal)
		jwt.PUT("/admin/updatePersonalPassword", controller.UpdatePersonalPassword)
		jwt.GET("/sysLoginInfo/list", controller.GetSysLoginInfoList)
		jwt.DELETE("/sysLoginInfo/batch/delete", controller.BatchDeleteSysLoginInfo)
		jwt.DELETE("/sysLoginInfo/delete", controller.DeleteSysLoginInfoById)
		jwt.DELETE("/sysLoginInfo/clean", controller.CleanSysLoginInfo)
		jwt.GET("/sysOperationLog/list", controller.GetSysOperationLogList)
		jwt.DELETE("/sysOperationLog/delete", controller.DeleteSysOperationLogById)
		jwt.DELETE("/sysOperationLog/batch/delete", controller.BatchDeleteSysOperationLog)
		jwt.DELETE("/sysOperationLog/clean", controller.CleanSysOperationLog)
	}
}
