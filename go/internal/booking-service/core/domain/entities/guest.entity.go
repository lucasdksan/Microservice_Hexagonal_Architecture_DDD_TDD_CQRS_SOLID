package entities

type Guest struct {
	Id      uint `gorm:"primaryKey"`
	Name    string
	Surname string
	Email   string
}
