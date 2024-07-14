package po

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `gorm:"column:uuid; type:char(36); unique; not null; index:idx_uuid"`
	UserName string    `gorm:"column:user_name; type:varchar(50); not null; index:idx_user_name"`
	IsActive bool      `gorm:"column:is_active; type:boolean; not null; default:true"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_users"
}
