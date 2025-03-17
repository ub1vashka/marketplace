package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ub1vashka/marketplace/internal/domain/models"
	"github.com/ub1vashka/marketplace/internal/logger"
	"github.com/ub1vashka/marketplace/internal/storage/storageerror"
)

func (s *Server) getProductByIDHandler(ctx *gin.Context) {
	log := logger.Get()
	productID := ctx.Param("id")
	product, err := s.productService.GetProductByID(productID)
	if err != nil {
		log.Error().Err(err).Msg("get product by ID failed")
		if errors.Is(err, storageerror.ErrProductIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (s *Server) deleteProductHandler(ctx *gin.Context) {
	log := logger.Get()
	productID := ctx.Param("id")
	err := s.productService.DeleteProduct(productID)
	if err != nil {
		log.Error().Err(err).Msg("delete product failed")
		if errors.Is(err, storageerror.ErrProductIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

func (s *Server) addProductHandler(ctx *gin.Context) {
	log := logger.Get()
	var product models.Product
	err := ctx.ShouldBindBodyWithJSON(&product)
	if err != nil {
		log.Error().Err(err).Msg("unmarshall body failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productID, err := s.productService.SaveProduct(product)
	if err != nil {
		log.Error().Err(err).Msg("save product failed")
		if errors.Is(err, storageerror.ErrProductAlredyExist) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusCreated, "product %s was saved", productID)
}

func (s *Server) getAllProductsHandler(ctx *gin.Context) {
	log := logger.Get()
	products, err := s.productService.GetAllProducts()
	if err != nil {
		log.Error().Err(err).Msg("get all products form storage failed")
		if errors.Is(err, storageerror.ErrEmptyStorage) {
			ctx.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (s *Server) updateProductHandler(ctx *gin.Context) {
	log := logger.Get()
	var product models.Product
	err := ctx.ShouldBindBodyWithJSON(&product)
	if err != nil {
		log.Error().Err(err).Msg("unmarshall body failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.productService.UpdateProduct(product.ProductID.String(), product)
	if err != nil {
		log.Error().Err(err).Msg("update product failed")
		if errors.Is(err, storageerror.ErrProductIDNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}
