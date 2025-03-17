package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ub1vashka/marketplace/internal/domain/models"
	"github.com/ub1vashka/marketplace/internal/logger"
	"github.com/ub1vashka/marketplace/internal/storage/storageerror"
	"golang.org/x/crypto/bcrypt"
)

type MapStorage struct {
	stor         map[string]models.User
	productStor  map[string]models.Product
	purchaseStor map[string]models.Purchase
}

func New() *MapStorage {
	return &MapStorage{
		stor:         make(map[string]models.User),
		productStor:  make(map[string]models.Product),
		purchaseStor: make(map[string]models.Purchase),
	}
}

func (ms *MapStorage) SaveUser(user models.User) (string, error) {
	log := logger.Get()
	for _, usr := range ms.stor {
		if user.Email == usr.Email {
			return ``, errors.New("user alredy exist")
		}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return ``, err
	}
	user.Password = string(hash)
	uid := uuid.New()
	user.UID = uid
	ms.stor[user.UID.String()] = user
	log.Debug().Any("storage", ms.stor).Msg("check storage")
	return uid.String(), nil
}

func (ms *MapStorage) DeleteUser(uid string) error {
	_, ok := ms.stor[uid]
	if !ok {
		return storageerror.ErrUserNotFound
	}
	delete(ms.stor, uid)
	return nil
}

func (ms *MapStorage) GetUsersProfile() ([]models.User, error) {
	if len(ms.stor) == 0 {
		return nil, storageerror.ErrEmptyUserStorage
	}
	var users []models.User
	for _, user := range ms.stor {
		users = append(users, user)
	}
	return users, nil
}

func (ms *MapStorage) GetUserProfile(uid string) (models.User, error) {
	user, ok := ms.stor[uid]
	if !ok {
		return models.User{}, storageerror.ErrUserNotFound
	}
	return user, nil
}

func (ms *MapStorage) ValidateUser(user models.UserLogin) (string, error) {
	for key, usr := range ms.stor {
		if user.Email == usr.Email {
			if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password)); err != nil {
				return ``, errors.New("invalid user password")
			}
			return key, nil
		}
	}
	return ``, errors.New("user no exist")
}

func (ms *MapStorage) SaveProduct(product models.Product) (string, error) {
	log := logger.Get()
	for _, p := range ms.productStor {
		if product.Name == p.Name && product.Description == p.Description {
			return ``, storageerror.ErrProductAlredyExist
		}
	}
	productID := uuid.New()
	product.ProductID = productID
	ms.productStor[product.ProductID.String()] = product
	log.Debug().Any("product storage", ms.stor).Msg("check storage")
	return productID.String(), nil
}

func (ms *MapStorage) GetAllProducts() ([]models.Product, error) {
	if len(ms.productStor) == 0 {
		return nil, storageerror.ErrEmptyStorage
	}
	var products []models.Product
	for _, product := range ms.productStor {
		products = append(products, product)
	}
	return products, nil
}

func (ms *MapStorage) GetProductByID(productID string) (models.Product, error) {
	product, ok := ms.productStor[productID]
	if !ok {
		return models.Product{}, storageerror.ErrProductIDNotFound
	}
	return product, nil
}

func (ms *MapStorage) DeleteProduct(productID string) error {
	_, ok := ms.productStor[productID]
	if !ok {
		return storageerror.ErrProductIDNotFound
	}
	delete(ms.productStor, productID)
	return nil
}

func (ms *MapStorage) UpdateProduct(productID string, product models.Product) error {
	_, ok := ms.productStor[productID]
	if !ok {
		return storageerror.ErrProductIDNotFound
	}
	ms.productStor[productID] = product
	return nil
}
