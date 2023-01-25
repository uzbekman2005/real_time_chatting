package repo

import "github.com/uzbekman2005/real_time_chatting/api/models"

type ChatStorageI interface {
	CreateAdmin(*models.CreateAdminReq) (*models.AdminRes, error)
	UpdateAdmin(*models.UpdateAdminInfo) error
	GetAllAdmins() ([]*models.AdminRes, error)
	GetAdmin(id int) (*models.AdminRes, error)
	GetAdminByName(name string) (*models.AdminRes, error)
	DeleteAdmin(id int) error
	IsUniqueAdminName(admin_name string) (bool, error)
}
