package postgres

import (
	"time"

	"github.com/uzbekman2005/real_time_chatting/api/models"
)

func (p *ChatRepo) CreateAdmin(req *models.CreateAdminReq) (*models.AdminRes, error) {
	query := `
		INSERT INTO admins (
			admin_name, 
			admin_password,
			admin_role
		) VALUES($1, $2, $3)
		RETURNING 
		id,
		created_at,
		updated_at
	`
	res := &models.AdminRes{}
	err := p.Db.QueryRow(query,
		req.Name,
		req.Password,
		req.Role).Scan(
		&res.Id,
		&CreatedAt,
		&UpdatedAt,
	)
	res.Name = req.Name
	res.Role = req.Role
	res.CreatedAt = CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = UpdatedAt.Format(time.RFC3339)

	return res, err
}

func (p *ChatRepo) UpdateAdmin(req *models.UpdateAdminInfo) error {
	query := `
		UPDATE admins SET 
			admin_name=$1,
			admin_password=$2,
			admin_role=$3,
			updated_at=now()
		WHERE id=$4
	`
	err := p.Db.QueryRow(query,
		req.NewName,
		req.NewPassword,
		req.NewRole,
		req.Id).Err()

	return err
}

func (p *ChatRepo) GetAllAdmins() ([]*models.AdminRes, error) {
	query := `
		SELECT
			id,
			admin_name, 
			admin_role,
			created_at,
			updated_at
		FROM admins
	`
	res := []*models.AdminRes{}
	rows, err := p.Db.Query(query)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := &models.AdminRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Role,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		temp.CreatedAt = CreatedAt.Format(time.RFC3339)
		temp.UpdatedAt = UpdatedAt.Format(time.RFC3339)
		res = append(res, temp)
	}

	return res, nil
}

func (p *ChatRepo) GetAdmin(id int) (*models.AdminRes, error) {
	query := `
		SELECT
			id,
			admin_name, 
			admin_role,
			admin_password,
			created_at,
			updated_at
		FROM admins WHERE id=$1
	`
	res := &models.AdminRes{}
	err := p.Db.QueryRow(query, id).Scan(
		&res.Id,
		&res.Name,
		&res.Role,
		&res.Password,
		&CreatedAt,
		&UpdatedAt,
	)

	res.CreatedAt = CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = UpdatedAt.Format(time.RFC3339)

	return res, err
}

func (p *ChatRepo) GetAdminByName(name string) (*models.AdminRes, error) {
	query := `
		SELECT
			id,
			admin_name, 
			admin_role,
			admin_password,
			created_at,
			updated_at
		FROM admins WHERE admin_name=$1
	`
	res := &models.AdminRes{}
	err := p.Db.QueryRow(query, name).Scan(
		&res.Id,
		&res.Name,
		&res.Role,
		&res.Password,
		&CreatedAt,
		&UpdatedAt,
	)

	res.CreatedAt = CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = UpdatedAt.Format(time.RFC3339)

	return res, err
}

func (p *ChatRepo) DeleteAdmin(id int) error {
	query := `DELETE FROM admins WHERE id=$1`
	_, err := p.Db.Exec(query, id)

	return err
}

func (p *ChatRepo) IsUniqueAdminName(adminName string) (bool, error) {
	query := `
	SELECT $1 in (SELECT admin_name from admins)
	`
	response := false
	err := p.Db.QueryRow(query, adminName).Scan(
		&response,
	)

	return !response, err
}
