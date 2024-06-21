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
	NamaMahasiswa   *string `json:"nama_mahasiswa"`
	Status          *int    `json:"status"`
	Password        string  `json:"password"`
	Deskripsi       *string `gorm:"not null" json:"deskripsi"`
	JenisOrmawa     *string `gorm:"not null" json:"jenis_ormawa"`
	NoTelepon       *string `gorm:"not null" json:"no_telepon"`
	Email           *string `gorm:"not null" json:"email"`
	TanggalLahir    *string `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin    *uint   `gorm:"not null" json:"jenis_kelamin"`
	TempatLahir     *string `gorm:"not null" json:"tempat_lahir"`
	AlamatTinggal   *string `gorm:"not null" json:"alamat_tinggal"`
	ConfirmPassword string  `json:"confirm_password"`
	Logo            *string `json:"logo"`
	Photo           *string `json:"photo"`
	Cover           *string `json:"cover"`
}

type TokenClaims struct {
	User *User `json:"user"`
	jwt.StandardClaims
}

type LoginPayload struct {
	Username *string `json:"username"`
	Password string  `json:"password"`
}

type ResetPassword struct {
	UserID               string `json:"user_id"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type UserRepository interface {
	CreateUser(req *User) (*User, error)
	GetUser(username string) (*User, error)
	GetUserById(id uint) (*User, error)
	UpdatePassword(req *ResetPassword) error
	UpdateProfile(req *User) error
}

type UserUseCase interface {
	CreateUser(ctx context.Context, req *CreateUser) (*User, error)
	LoginUser(ctx context.Context, req *LoginPayload) (*User, string, error)
	GetUserById(ctx context.Context, id uint) (*User, error)
	PengajuanEventOrmawa(ctx context.Context) (map[string]any, error)
	ResetPassword(ctx context.Context, req *ResetPassword) error
	UpdateProfile(ctx context.Context, req *CreateUser) error
}
