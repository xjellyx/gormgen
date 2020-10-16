package example

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCreateAdmin = errors.New("create Admin failed")
	ErrDeleteAdmin = errors.New("delete Admin failed")
	ErrGetAdmin    = errors.New("get Admin failed")
	ErrUpdateAdmin = errors.New("update Admin failed")
)

// NewAdmin new
func NewAdmin() *Admin {
	return new(Admin)
}

// Add add one record
func (t *Admin) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {

		err = ErrCreateAdmin
		return
	}
	return
}

// Delete delete record
func (t *Admin) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {

		err = ErrDeleteAdmin
		return
	}
	return
}

// Updates update record
func (t *Admin) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Where("id = ?", t.ID).Updates(m).Error; err != nil {

		err = ErrUpdateAdmin
		return
	}
	return
}

// GetAdminAll get all record
func GetAdminAll(db *gorm.DB) (ret []*Admin, err error) {
	if err = db.Find(&ret).Error; err != nil {

		err = ErrGetAdmin
		return
	}
	return
}

// GetAdminCount get count
func GetAdminCount(db *gorm.DB) (ret int64) {
	db.Model(&Admin{}).Count(&ret)
	return
}

// QueryByID query cond by ID
func (t *Admin) SetQueryByID(id uint) *Admin {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *Admin) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {

		err = ErrGetAdmin
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *Admin) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != nil {

		err = ErrDeleteAdmin
		return
	}
	return
}

// QueryByName query cond by Name
func (t *Admin) SetQueryByName(name string) *Admin {
	t.Name = name
	return t
}

// GetByName get one record by Name
func (t *Admin) GetByName(db *gorm.DB) (err error) {
	if err = db.First(t, "name = ?", t.Name).Error; err != nil {

		err = ErrGetAdmin
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *Admin) DeleteByName(db *gorm.DB) (err error) {
	if err = db.Delete(t, "name = ?", t.Name).Error; err != nil {

		err = ErrDeleteAdmin
		return
	}
	return
}