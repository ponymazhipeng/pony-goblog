package controller

import (
	"net/http"
	"pony-goblog/common"
	"pony-goblog/config"
	"pony-goblog/dao"
	"pony-goblog/models"
	"pony-goblog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole
	var pd = &models.PigeonholeData{
		config.Cfg.Viewer,
		config.Cfg.System,
		dao.GetCategorys(),
		service.PostByMonth(),
	}
	pigeonhole.WriteData(w, pd)
}
