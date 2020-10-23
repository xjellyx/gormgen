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
	if err = db.Model(&User{}).Where("id = ?", t.ID).Updates(m).Error; err != nil {

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

type QueryUserForm struct {
	CreatedAt *FieldData `json:"createdAt" form:"createdAt"`
	UpdatedAt *FieldData `json:"updatedAt" form:"updatedAt"`
	Age       *FieldData `json:"age" form:"age"`
	Email     *FieldData `json:"email" form:"email"`
	Order     []string   `json:"order" form:"order"`
	PageNum   int        `json:"pageNum" form:"pageNum"`
	PageSize  int        `json:"pageSize" form:"pageSize"`
}

//  GetUserList get User list some field value or some condition
func GetUserList(q *QueryUserForm, db *gorm.DB) (ret []*User, err error) {
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
