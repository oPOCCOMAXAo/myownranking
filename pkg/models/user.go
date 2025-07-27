package models

type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name;not null;default:'';size:255"`
	Email    string `gorm:"column:email;not null;default:'';size:255;index:idx_email,unique"`
	Password string `gorm:"column:password;not null;default:'';size:255;comment:'hashed password'"`
}

func (User) TableName() string {
	return "users"
}

type UserAPI struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (UserAPI) FromModel(user *User, _ int) *UserAPI {
	return &UserAPI{
		ID:   user.ID,
		Name: user.Name,
	}
}
