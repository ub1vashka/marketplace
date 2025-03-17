package service

import "github.com/ub1vashka/marketplace/internal/domain/models"

type PurchaseStorage interface {
	MakePurchase(models.Purchase) (string, error)
	GetUserPurchases() ([]models.Purchase, error)
	GetProductPurchases(string) ([]models.Purchase, error)
}

type PurchaseService struct {
	stor PurchaseStorage
}

func NewPurchaseService(stor PurchaseStorage) PurchaseService {
	return PurchaseService{stor: stor}
}
