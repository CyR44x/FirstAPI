package userService

import "gorm.io/gorm"

type UserRepository interface {
	PostUser(user User) (User, error)
	GetUsers() ([]User, error)
	PatchUserByID(id int, user User) (User, error)
	DeleteUserByID(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) PostUser(user User) (User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) GetUsers() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) PatchUserByID(id int, user User) (User, error) {
	existingUser := User{}
	result := r.db.First(&existingUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	result = r.db.Model(&existingUser).Updates(user)
	return existingUser, result.Error
}

func (r *userRepository) DeleteUserByID(id int) error {
	result := r.db.Delete(&User{}, id)
	return result.Error
}
