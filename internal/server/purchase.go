package server

import (
	"github.com/ub1vashka/marketplace/internal/domain/models"
	"github.com/ub1vashka/marketplace/internal/logger"
	"github.com/ub1vashka/marketplace/internal/storage/storageerror"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) makePurchaseHandler(ctx *gin.Context) {
	log := logger.Get()
	var purchase models.Purchase
	err := ctx.ShouldBindBodyWithJSON(&purchase)
	if err != nil {
		log.Error().Err(err).Msg("unmarshall purchase body failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = s.valid.Struct(purchase); err != nil {
		log.Error().Err(err).Msg("validate purchase input data failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchaseID, err := s.purchaseService.MakePurchase(purchase)
	if err != nil {
		log.Error().Err(err).Msg("make purchase failed")
		if errors.Is(err, storageerror.ErrProductIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusCreated, "Purchase created with ID: %s", purchaseID)
}

func (s *Server) getUserPurchasesHandler(ctx *gin.Context) {
	log := logger.Get()
	userID := ctx.Query("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}
	purchases, err := s.purchaseService.GetUserPurchases(userID)
	if err != nil {
		log.Error().Err(err).Msg("get user purchases failed")
		if errors.Is(err, storageerror.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(purchases) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "no purchases found"})
		return
	}
	ctx.JSON(http.StatusOK, purchases)
}

func (s *Server) getProductPurchasesHandler(ctx *gin.Context) {
	log := logger.Get()
	productID := ctx.Param("id")
	purchases, err := s.purchaseService.GetProductPurchases(productID)
	if err != nil {
		log.Error().Err(err).Msg("get product purchases failed")
		if errors.Is(err, storageerror.ErrProductIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(purchases) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "no purchases found for this product"})
		return
	}
	ctx.JSON(http.StatusOK, purchases)
}
