package example

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCreateadmin = errors.New("create admin failed")
	ErrDeleteadmin = errors.New("delete admin failed")
	ErrGetadmin    = errors.New("get admin failed")
	ErrUpdateadmin = errors.New("update admin failed")
)

// Newadmin new
func Newadmin() *admin {
	return new(admin)
}

// Add add one record
func (t *admin) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {

		err = ErrCreateadmin
		return
	}
	return
}

// Delete delete record
func (t *admin) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {

		err = ErrDeleteadmin
		return
	}
	return
}

// Updates update record
func (t *admin) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Where("id = ?", t.ID).Updates(m).Error; err != nil {

		err = ErrUpdateadmin
		return
	}
	return
}

// GetadminAll get all record
func GetadminAll(db *gorm.DB) (ret []*admin, err error) {
	if err = db.Find(&ret).Error; err != nil {

		err = ErrGetadmin
		return
	}
	return
}

// GetadminCount get count
func GetadminCount(db *gorm.DB) (ret int64) {
	db.Model(&admin{}).Count(&ret)
	return
}

// QueryByID query cond by ID
func (t *admin) SetQueryByID(id uint) *admin {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *admin) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {

		err = ErrGetadmin
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *admin) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != nil {

		err = ErrDeleteadmin
		return
	}
	return
}

// QueryByName query cond by Name
func (t *admin) SetQueryByName(name string) *admin {
	t.Name = name
	return t
}

// GetByName get one record by Name
func (t *admin) GetByName(db *gorm.DB) (err error) {
	if err = db.First(t, "name = ?", t.Name).Error; err != nil {

		err = ErrGetadmin
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *admin) DeleteByName(db *gorm.DB) (err error) {
	if err = db.Delete(t, "name = ?", t.Name).Error; err != nil {

		err = ErrDeleteadmin
		return
	}
	return
}
