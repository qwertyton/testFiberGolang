package model

type Respond struct {
	ActivePower int `gorm:"column:active_power;"`
	PowerInput  int `gorm:"column:power_input;"`
}
