package entity

// Users struct
type User struct {
	ID        int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Email     string `gorm:"column:email;not null"`
	FirstName string `gorm:"column:first_name;"`
	LastName  string `gorm:"column:last_name"`
	Password  string `gorm:"column:password;not null"`
	Active    int    `gorm:"column:active"`
	// Token          *string     `gorm:"column:token"`
	// TokenExpriedAt *time.Time  `gorm:"column:token_expired_at"`
	BaseModel
}

func (u *User) TableName() string {
	return "users"
}
