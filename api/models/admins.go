package models

type AdminRes struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`
	Password    string `json:"-"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateAdminReq struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UpdateAdminInfo struct {
	Id          int `json:"id"`
	NewName     string `json:"new_name"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	NewRole     string `json:"new_role"`
}

type DeleteAdminReq struct {
	Id int `json:"id"`
}
