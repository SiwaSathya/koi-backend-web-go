package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Username  *string        `gorm:"null" json:"username"`
	Role      string         `gorm:"not null" json:"role"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Ormawa    *Ormawa        `json:"ormawa"`
	Mahasiswa *Mahasiswa     `json:"mahasiswa"`
}

type CreateUser struct {
	ID              uint    `json:"id"`
	Nim             *uint   `json:"nim"`
	Username        *string `json:"username"`
	Role            string  `json:"role"`
	NamaOrmawa      *string `json:"nama_ormawa"`
	Status          *int    `json:"status"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
}

type TokenClaims struct {
	User *User `json:"user"`
	jwt.StandardClaims
}

type LoginPayload struct {
	Username *string `json:"username"`
	Password string  `json:"password"`
}

type UserRepository interface {
	CreateUser(req *User) (*User, error)
	GetUser(username string) (*User, error)
}

type UserUseCase interface {
	CreateUser(ctx context.Context, req *CreateUser) (*User, error)
	LoginUser(ctx context.Context, req *LoginPayload) (*User, string, error)
}
