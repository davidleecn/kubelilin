package req

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"kubelilin/pkg/page"
)

type DevopsProjectReq struct {
	mvc.RequestBody
	page.PageRequest

	Name string `json:"name" uri:"name"`
}
