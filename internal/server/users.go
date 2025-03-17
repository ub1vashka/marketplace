package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ub1vashka/marketplace/internal/domain/models"
	"github.com/ub1vashka/marketplace/internal/logger"
	"github.com/ub1vashka/marketplace/internal/storage/storageerror"
)

func (s *Server) getUserByIDHandler(ctx *gin.Context) {
	log := logger.Get()
	uid := ctx.Param("id")
	user, err := s.uService.GetUserProfile(uid)
	if err != nil {
		log.Error().Err(err).Msg("get user by ID failed")
		if errors.Is(err, storageerror.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *Server) getUsersHandler(ctx *gin.Context) {
	log := logger.Get()
	users, err := s.uService.GetUsersProfile()
	if err != nil {
		log.Error().Err(err).Msg("get all users form storage failed")
		if errors.Is(err, storageerror.ErrEmptyUserStorage) {
			ctx.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (s *Server) deleteUserHandler(ctx *gin.Context) {
	log := logger.Get()
	uid := ctx.Param("id")
	err := s.uService.DeleteUser(uid)
	if err != nil {
		log.Error().Err(err).Msg("delete user failed")
		if errors.Is(err, storageerror.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (s *Server) loginHandler(ctx *gin.Context) { //nolint:dupl //todo
	log := logger.Get()
	var user models.UserLogin
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error().Err(err).Msg("unmarshall login body failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = s.valid.Struct(user); err != nil {
		log.Error().Err(err).Msg("validate login user input data failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := s.uService.LoginUser(user)
	if err != nil {
		log.Error().Err(err).Msg("user login validate failed")
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid input data", "error": err.Error()})
		return
	}
	ctx.String(http.StatusCreated, "User was logined; user id: %s", uid)
}

func (s *Server) registerHandler(ctx *gin.Context) { //nolint:dupl //todo
	log := logger.Get()
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error().Err(err).Msg("unmarshall body failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = s.valid.Struct(user); err != nil {
		log.Error().Err(err).Msg("validate user input data failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := s.uService.RegisterUser(user)
	if err != nil {
		log.Error().Err(err).Msg("user register failed")
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid input data", "error": err.Error()})
		return
	}
	ctx.String(http.StatusCreated, "User was created; user id: %s", uid)
}
