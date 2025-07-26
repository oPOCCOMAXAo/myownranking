package models

type User struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0"`
	UpdatedAt int64  `gorm:"column:updated_at;not null;default:0"`
	Name      string `gorm:"column:name;not null;default:'';size:255"`
	Email     string `gorm:"column:email;not null;default:'';size:255;index:idx_email,unique"`
	Password  string `gorm:"column:password;not null;default:'';size:255;comment:'hashed password'"`
}

func (User) TableName() string {
	return "users"
}
