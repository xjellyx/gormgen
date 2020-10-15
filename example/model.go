package example

import "gorm.io/gorm"

//go:generate gormgen -structs User,Admin -inputDir . -importPkgs gorm.io/gorm -transformErr true
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int
	Email string
}

type Admin struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int
	Email string
}
