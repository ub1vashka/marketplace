package service

import (
	"time"

	"github.com/ub1vashka/marketplace/internal/domain/models"
	"github.com/ub1vashka/marketplace/internal/logger"
)

type UserStorage interface {
	SaveUser(models.User) (string, error)
	ValidateUser(models.UserLogin) (string, error)
	GetUsersProfile() ([]models.User, error)
	GetUserProfile(string) (models.User, error)
	DeleteUser(string) error
}

type UserService struct {
	stor UserStorage
}

func (us *UserService) RegisterUser(user models.User) (string, error) {
	log := logger.Get()
	user.RegisterDate = time.Now()
	uid, err := us.stor.SaveUser(user)
	if err != nil {
		log.Error().Err(err).Msg("save user failed")
		return ``, err
	}
	return uid, nil
}

func (us *UserService) LoginUser(user models.UserLogin) (string, error) {
	log := logger.Get()
	uid, err := us.stor.ValidateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("validate user failed")
		return ``, err
	}
	return uid, nil
}

func NewUserService(stor UserStorage) UserService {
	return UserService{stor: stor}
}

func (us *UserService) GetUsersProfile() ([]models.User, error) {
	return us.stor.GetUsersProfile()
}

func (us *UserService) GetUserProfile(uid string) (models.User, error) {
	return us.stor.GetUserProfile(uid)
}

func (us *UserService) DeleteUser(uid string) error {
	return us.stor.DeleteUser(uid)
}
