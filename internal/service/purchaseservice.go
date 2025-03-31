package service

import (
	"github.com/ub1vashka/marketplace/internal/domain/models"
)

type PurchaseStorage interface {
	MakePurchase(models.Purchase) (string, error)
	GetUserPurchases(string) ([]models.Purchase, error)
	GetProductPurchases(string) ([]models.Purchase, error)
}

type PurchaseService struct {
	stor PurchaseStorage
}

func NewPurchaseService(stor PurchaseStorage) PurchaseService {
	return PurchaseService{stor: stor}
}

func (ps *PurchaseService) MakePurchase(purchase models.Purchase) (string, error) {
	return ps.stor.MakePurchase(purchase)
}

func (ps *PurchaseService) GetUserPurchases(userID string) ([]models.Purchase, error) {
	return ps.stor.GetUserPurchases(userID)
}

func (ps *PurchaseService) GetProductPurchases(productID string) ([]models.Purchase, error) {
	return ps.stor.GetProductPurchases(productID)
}
