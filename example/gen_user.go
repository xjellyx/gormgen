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

// GetByAge get one record by Age
func (t *User) GetByAge(db *gorm.DB) (err error) {
	if err = db.First(t, "age = ?", t.Age).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// DeleteByAge delete record by Age
func (t *User) DeleteByAge(db *gorm.DB) (err error) {
	if err = db.Delete(t, "age = ?", t.Age).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}

// GetByEmail get one record by Email
func (t *User) GetByEmail(db *gorm.DB) (err error) {
	if err = db.First(t, "email = ?", t.Email).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// DeleteByEmail delete record by Email
func (t *User) DeleteByEmail(db *gorm.DB) (err error) {
	if err = db.Delete(t, "email = ?", t.Email).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}
