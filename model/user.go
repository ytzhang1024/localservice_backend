package model

import (
	"6251/localservice/utils"
	"gorm.io/gorm"
)

// ===============================================================================
// = User Info
type User struct {
	gorm.Model
	Email    string `json:"email" form:"email" gorm:"size:128;not null"`
	Password string `json:"password" form:"password" gorm:"size:128;not null"`
	NickName string `json:"nick_name" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Mobile   string `json:"mobile" gorm:"size:11"`
	Role     string `json:"role" form:"role" gorm:"size:128;not null"`
}

// ===============================================================================
// Encrypt the password
func (m *User) Encrypt() error {
	stHash, err := utils.Encrypt(m.Password)
	if err == nil {
		m.Password = stHash
	}

	return err
}

func (m *User) BeforeCreate(orm *gorm.DB) error {
	return m.Encrypt()
}

// ===============================================================================
// = User Login Info
type LoginUser struct {
	ID    uint
	Email string
}
