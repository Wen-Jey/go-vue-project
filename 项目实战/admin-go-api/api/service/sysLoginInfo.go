// 登录日志 服务层

package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

type ISysLoginInfoService interface {
	GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, pageNum int)
	BatchDeleteSysLoginInfo(c *gin.Context, dto entity.DelSysLoginInfoDto)
	DeleteSysLoginInfo(c *gin.Context, dto entity.SysLoginInfoIdDto)
	CleanSysLoginInfo(c *gin.Context)
}

type SysLoginInfoServiceImpl struct{}

// 清空登录日志
func (s SysLoginInfoServiceImpl) CleanSysLoginInfo(c *gin.Context) {
	dao.CleanSysLoginInfo()
	result.Success(c, true)
}

// 批量删除登录日志
func (s SysLoginInfoServiceImpl) BatchDeleteSysLoginInfo(c *gin.Context, dto entity.DelSysLoginInfoDto) {
	dao.BatchDeleteSysLoginInfo(dto)
	result.Success(c, true)
}

// 根据id删除登录日志
func (s SysLoginInfoServiceImpl) DeleteSysLoginInfo(c *gin.Context, dto entity.SysLoginInfoIdDto) {
	dao.DeleteSysLoginInfoById(dto)
	result.Success(c, true)
}

// 分页获取登录日志列表
func (s SysLoginInfoServiceImpl) GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, pageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if pageNum < 1 {
		pageNum = 1
	}
	sysLoginInfo, count := dao.GetSysLoginInfoList(Username, LoginStatus, BeginTime, EndTime, PageSize, pageNum)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": pageNum, "list": sysLoginInfo})
}

var sysLoginInfoService = SysLoginInfoServiceImpl{}

func SysLoginInfoService() ISysLoginInfoService {
	return &sysLoginInfoService
}
