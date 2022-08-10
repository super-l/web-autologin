package model

import "svjia-cookie/core"

type mSystemConfig struct{}

var MConfig = mSystemConfig{}

type systemConfig struct {
	Id       uint64 `json:"id" gorm:"primary_key column:id"`
	Category uint   `json:"category" gorm:"column:category"`
	Name     string `json:"name"  gorm:"column:name"`
	Value    string `json:"value" gorm:"column:value"`
	Remark   string `json:"remark" gorm:"column:remark"`
}

func (mSystemConfig) Update(data systemConfig) error {
	err := core.DB.Model(&systemConfig{}).Where("category = ?", data.Category).Where("name = ?", data.Name).Updates(map[string]interface{}{"value": data.Value}).Error
	return err
}

func (mSystemConfig) UpdateConfig(category string, name string, value string) error {
	err := core.DB.Model(&systemConfig{}).Where("category = ?", category).Where("name = ?", name).Updates(map[string]interface{}{"value": value}).Error
	return err
}
