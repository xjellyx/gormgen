package example

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCreateuser = errors.New("create user failed")
	ErrDeleteuser = errors.New("delete user failed")
	ErrGetuser    = errors.New("get user failed")
	ErrUpdateuser = errors.New("update user failed")
)

// Newuser new
func Newuser() *user {
	return new(user)
}

// Add add one record
func (t *user) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {

		err = ErrCreateuser
		return
	}
	return
}

// Delete delete record
func (t *user) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {

		err = ErrDeleteuser
		return
	}
	return
}

// Updates update record
func (t *user) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Where("id = ?", t.ID).Updates(m).Error; err != nil {

		err = ErrUpdateuser
		return
	}
	return
}

// GetuserAll get all record
func GetuserAll(db *gorm.DB) (ret []*user, err error) {
	if err = db.Find(&ret).Error; err != nil {

		err = ErrGetuser
		return
	}
	return
}

// GetuserCount get count
func GetuserCount(db *gorm.DB) (ret int64) {
	db.Model(&user{}).Count(&ret)
	return
}

// QueryByID query cond by ID
func (t *user) SetQueryByID(id uint) *user {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *user) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {

		err = ErrGetuser
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *user) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != nil {

		err = ErrDeleteuser
		return
	}
	return
}

// QueryByName query cond by Name
func (t *user) SetQueryByName(name string) *user {
	t.Name = name
	return t
}

// GetByName get one record by Name
func (t *user) GetByName(db *gorm.DB) (err error) {
	if err = db.First(t, "name = ?", t.Name).Error; err != nil {

		err = ErrGetuser
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *user) DeleteByName(db *gorm.DB) (err error) {
	if err = db.Delete(t, "name = ?", t.Name).Error; err != nil {

		err = ErrDeleteuser
		return
	}
	return
}
