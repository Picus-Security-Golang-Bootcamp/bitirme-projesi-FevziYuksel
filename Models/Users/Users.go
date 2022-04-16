package Users

import "gorm.io/gorm"

//unique id ve cart'ı düzelt
type Users struct {
	gorm.Model
	//Id       uint   `json:"id" gorm:"primarykey"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=2,max=20"`
	Role     string `json:"role" gorm:"default:costumer"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt sql.NullTime `gorm:"index"`
}

func NewUser(email string, password string, role string) *Users {
	return &Users{
		Email:    email,
		Password: password,
		Role:     role,
	}
}
func (u *Users) GetEmail() string {
	return u.Email
}
func (u *Users) GetPassword() string {
	return u.Password
}
func (u *Users) GetRole() string {
	return u.Role
}
