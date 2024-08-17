package model

type GoCrmUser struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement"`
	Username string `gorm:"column:username;type:varchar(50);not null;"`
	Password string `gorm:"column:password;type:varchar(255);not null;"`
	Email    string `gorm:"column:email;type:varchar(50);not null;"`
	Phone    string `gorm:"column:phone;type:varchar(20);not null;"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null;"`
}

func (*GoCrmUser) TableName() string {
	return "go_crm_user"
}
