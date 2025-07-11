package models

type Role string

const (
    RolePayer Role = "payer"
    RolePayee Role = "payee"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"not null"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role     Role   `gorm:"not null"`
}
