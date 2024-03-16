package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lesion45/pinterest-clone/internal/lib/api/request"
	_ "github.com/lesion45/pinterest-clone/internal/lib/api/request"
	"github.com/lesion45/pinterest-clone/internal/lib/api/response"
	"github.com/lesion45/pinterest-clone/internal/lib/logger/sl"
	"github.com/lesion45/pinterest-clone/storage"
	"github.com/lesion45/pinterest-clone/storage/postgres"
	"log/slog"
	"net/http"
)

// Request fields
// * sessionInfo - NOT RELEASED - IN PROCESS
func HomePage(ctx *gin.Context) {
	return
}

// Request fields
// * sessionInfo - NOT RELEASED - IN PROCESS
func GetAllPins(ctx *gin.Context) {
	return
}

// Request fields
// * login
// * password
// * sessionInfo - NOT RELEASED - IN PROCESS
func Login(ctx *gin.Context) {
	return
}

// Request fields
// * login
// * password
// * sessionInfo - NOT RELEASED - IN PROCESS
func Register(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.Register"
		log.With(slog.String("op", op))

		var req request.RegisterReq

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.Error("failed to decode request"))

			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = s.AddUser(req.Nickname, req.Password)

		// TODO: add errors "Nickname already exists" and "Login already exists" instead "User already exists"
		if errors.Is(err, storage.ErrUserExists) {
			log.Info("user already exists")
			ctx.IndentedJSON(http.StatusOK, response.Error("user already exists"))

			return
		}

		if err != nil {
			log.Error("failed to register", err)
			ctx.IndentedJSON(http.StatusOK, response.Error("failed to register"))

			return
		}

		log.Info("user added")

		ctx.IndentedJSON(http.StatusOK, response.Response{
			Status: response.StatusOK,
		})
	}
}

// Request fields
// * userID
// * imageURL
// * description
// * sessionInfo - NOT RELEASED - IN PROCESS
func CreatePin(ctx *gin.Context) {

	return
}

// Request fields
// * pinID
// * sessionInfo - NOT RELEASED - IN PROCESS
func GetPin(ctx *gin.Context) {
	return
}
