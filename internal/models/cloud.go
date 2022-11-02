package models

import "gorm.io/gorm"

type ConfigModel struct {
	gorm.Model

	SecureID string `gorm:"unique"`
	Name     string
	Author   string
	Owner    string
	Content  string
	Deleted  bool
}

type ScriptModel struct {
	gorm.Model

	SecureID   string `gorm:"unique"`
	Name       string
	Author     string
	Owner      string
	Content    string
	OpenSource bool
	Deleted    bool
}
