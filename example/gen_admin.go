package example

import (
	"errors"
	"time"

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
	if err = db.Model(&Admin{}).Where("id = ?", t.ID).Updates(m).Error; err != nil {

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

type QueryAdminForm struct {
	CreatedAt *fieldData `json:"createdAt" form:"createdAt"`
	UpdatedAt *fieldData `json:"updatedAt" form:"updatedAt"`
	Age       *fieldData `json:"age" form:"age"`
	Email     *fieldData `json:"email" form:"email"`
	Order     []string   `json:"order" form:"order"`
	PageNum   int        `json:"pageNum" form:"pageNum"`
	PageSize  int        `json:"pageSize" form:"pageSize"`
}

//  GetAdminList get Admin list some field value or some condition
func GetAdminList(q *QueryAdminForm, db *gorm.DB) (ret []*Admin, err error) {
	// order
	if len(q.Order) > 0 {
		for _, v := range q.Order {
			db = db.Order(v)
		}
	}
	// pageSize
	if q.PageSize != 0 {
		db = db.Limit(q.PageSize)
	}
	// pageNum
	if q.PageNum != 0 {
		q.PageNum = (q.PageNum - 1) * q.PageSize
		db = db.Offset(q.PageNum)
	}

	// CreatedAt
	if q.CreatedAt != nil {
		db = db.Where("created_at"+q.CreatedAt.Symbol+"?", q.CreatedAt.Value)
	}
	// UpdatedAt
	if q.UpdatedAt != nil {
		db = db.Where("updated_at"+q.UpdatedAt.Symbol+"?", q.UpdatedAt.Value)
	}
	// Age
	if q.Age != nil {
		db = db.Where("age"+q.Age.Symbol+"?", q.Age.Value)
	}
	// Email
	if q.Email != nil {
		db = db.Where("email"+q.Email.Symbol+"?", q.Email.Value)
	}
	if err = db.Find(&ret).Error; err != nil {
		return
	}
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
