package model

type Category struct {
	//gorm.Model
	ID   uint `gorm:"primarykey"`
	Name string
}
