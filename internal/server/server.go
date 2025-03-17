package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ub1vashka/marketplace/internal/config"
	"github.com/ub1vashka/marketplace/internal/logger"
	"github.com/ub1vashka/marketplace/internal/service"
)

type Server struct {
	serve           *http.Server
	valid           *validator.Validate
	uService        service.UserService
	productService  service.ProductService
	purchaseService service.PurchaseService
}

func New(cfg config.Config, us service.UserService, pros service.ProductService, purch service.PurchaseService) *Server {
	addrStr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server := http.Server{ //nolint:gosec //todo
		Addr: addrStr,
	}
	vald := validator.New()
	return &Server{
		serve:           &server,
		valid:           vald,
		uService:        us,
		productService:  pros,
		purchaseService: purch,
	}
}

func (s *Server) Run() error {
	log := logger.Get()
	router := s.configRouting()
	s.serve.Handler = router
	log.Info().Str("addr", s.serve.Addr).Msg("server start")
	if err := s.serve.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("runing server failed")
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.serve.Shutdown(ctx)
}

func (s *Server) configRouting() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) { ctx.String(http.StatusOK, "Hello, my friend!") })
	users := router.Group("/users")
	{
		users.GET("/:id", s.getUserByIDHandler)
		users.GET("/list", s.getUsersHandler)
		users.DELETE("/:id", s.deleteUserHandler)
		users.POST("/register", s.registerHandler)
		users.POST("/login", s.loginHandler)
	}
	products := router.Group("/products")
	{
		products.GET("/:id", s.getProductByIDHandler)
		products.DELETE("/:id", s.deleteProductHandler)
		products.GET("/catalog", s.getAllProductsHandler)
		products.POST("/add", s.addProductHandler)
		products.POST("/update", s.updateProductHandler)
	}
	purchases := router.Group("/purchases")
	{
		purchases.POST("/:id", s.MakePurchaseHandler)
		purchases.GET("/history", s.GetUserPurchasesHandler)
		purchases.GET("/product/:id", s.GetProductPurchasesHandler)
	}

	return router
}
