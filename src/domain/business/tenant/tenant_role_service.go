package tenant

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"sgr/api/req"
	"sgr/domain/database/models"
	"sgr/pkg/page"
)

type TenantRoleService struct {
	db *gorm.DB
}

func NewTenantRoleService(db *gorm.DB) *TenantRoleService {
	return &TenantRoleService{
		db: db,
	}
}

func (trs *TenantRoleService) CreateTenantRole(req *req.TenantRoleReq) (bool, *models.SgrTenantRole) {
	var tenantRoelModel = &models.SgrTenantRole{}
	copier.Copy(tenantRoelModel, req)
	trs.db.Create(tenantRoelModel)
	return tenantRoelModel.ID != 0, tenantRoelModel
}

func (trs *TenantRoleService) UpdateTenantRole(req *req.TenantRoleReq) (bool, *models.SgrTenantRole) {
	var tenantRoelModel = &models.SgrTenantRole{}
	copier.Copy(tenantRoelModel, req)
	res := trs.db.Save(tenantRoelModel)
	return res.RowsAffected > 0, tenantRoelModel
}

func (trs *TenantRoleService) DeleteTenantRole(id string) bool {
	res := trs.db.Model(&models.SgrTenantRole{}).Where("id = ?", id).Update(models.SgrTenantRoleColumns.Status, 0)
	return res.RowsAffected > 0
}

func (trs *TenantRoleService) QueryTenantRoleList(roleId string, tenantId int, pageIndex int, pageSize int) *page.Page {
	var data = &[]models.SgrTenantRole{}

	condition := trs.db.Model(&models.SgrTenantRole{})
	if roleId != "" {
		condition.Where("role_name like ?", "%"+roleId+"%")
	}
	if tenantId > 0 {
		condition.Where("tenant_id = ?", tenantId)
	}

	return page.StartPage(condition, pageIndex, pageSize).DoFind(data)
}
