package example

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCreateUser = errors.New("create User failed")
	ErrDeleteUser = errors.New("delete User failed")
	ErrGetUser    = errors.New("get User failed")
	ErrUpdateUser = errors.New("update User failed")
)

// NewUser new
func NewUser() *User {
	return new(User)
}

// Add add one record
func (t *User) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {

		err = ErrCreateUser
		return
	}
	return
}

// Delete delete record
func (t *User) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}

// Updates update record
func (t *User) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Where("id = ?", t.ID).Updates(m).Error; err != nil {

		err = ErrUpdateUser
		return
	}
	return
}

// GetUserAll get all record
func GetUserAll(db *gorm.DB) (ret []*User, err error) {
	if err = db.Find(&ret).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// GetUserCount get count
func GetUserCount(db *gorm.DB) (ret int64) {
	db.Model(&User{}).Count(&ret)
	return
}

// QueryByID query cond by ID
func (t *User) SetQueryByID(id uint) *User {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *User) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *User) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}

// QueryByName query cond by Name
func (t *User) SetQueryByName(name string) *User {
	t.Name = name
	return t
}

// GetByName get one record by Name
func (t *User) GetByName(db *gorm.DB) (err error) {
	if err = db.First(t, "name = ?", t.Name).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *User) DeleteByName(db *gorm.DB) (err error) {
	if err = db.Delete(t, "name = ?", t.Name).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}
