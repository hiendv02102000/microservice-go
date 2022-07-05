package entity

// Users struct
type Users struct {
	ID        int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Email     string `gorm:"column:email;"`
	FirstName string `gorm:"column:first_name;"`
	LastName  string `gorm:"column:last_name"`
	Password  string `gorm:"column:password;not null"`
	Active    int    `gorm:"column:active"`
	// Token          *string     `gorm:"column:token"`
	// TokenExpriedAt *time.Time  `gorm:"column:token_expired_at"`
	BaseModel
}
