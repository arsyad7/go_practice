package repository

import (
	"gorm/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, req *models.User) error
	DeleteUser(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(req *models.User) error {
	return r.db.Create(req).Error
}

func (r *userRepo) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepo) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id=?", id).Error
	return &user, err
}

func (r *userRepo) UpdateUser(id uint, req *models.User) error {
	var user models.User
	err := r.db.Model(&user).Where("id = ?", id).Updates(models.User{Email: req.Email}).Error
	return err
}

func (r *userRepo) DeleteUser(id uint) error {
	var user models.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	return err
}
