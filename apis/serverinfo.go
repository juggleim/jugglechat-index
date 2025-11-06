package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/juggleim/jugglechat-index/dbs"
	"github.com/juggleim/jugglechat-index/errs"
	"github.com/juggleim/jugglechat-index/tools"
)

func GetServerInfo(ctx *gin.Context) {
	var app *dbs.AppNavDao
	var err error
	dao := dbs.AppNavDao{}
	appkey := ctx.Query("app_key")
	aliasNo := ctx.Query("no")
	if appkey == "" && aliasNo == "" {
		tools.ErrorHttpResp(ctx, errs.IMErrorCode_API_PARAM_REQUIRED)
		return
	}
	if appkey != "" {
		app, err = dao.FindByAppkey(appkey)
		if err != nil || app == nil {
			tools.SuccessHttpResp(ctx, map[string]string{
				"server_info": "",
			})
			return
		}
	} else {
		app, err = dao.FindByAliasNo(aliasNo)
		if err != nil || app == nil {
			tools.SuccessHttpResp(ctx, map[string]string{
				"server_info": "",
			})
			return
		}
	}
	tools.SuccessHttpResp(ctx, map[string]string{
		"server_info_plain": tools.ToJson(&ServerInfo{
			AppKey: app.AppKey,
			// AliasNo:    app.AliasNo,
			ImServers:  []string{app.WsUrl},
			AppServers: []string{app.AppUrl},
		})})
}

type ServerInfo struct {
	AppKey string `json:"app_key"`
	// AliasNo    string   `json:"alias_no"`
	ImServers  []string `json:"im_servers"`
	AppServers []string `json:"app_servers"`
}
