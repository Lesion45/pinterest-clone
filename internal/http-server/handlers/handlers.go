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

func HomePage(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		return
	}
}

func AddPin(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.AddPin"

		log.With(slog.String("op", op))

		var req request.AddPinReq

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = s.AddPin(req.Username, req.ImageURL)

		// TODO: ADD MORE TYPES OF HANDLING ERRORS
		if err != nil {
			log.Info("failed to add pin")
			ctx.IndentedJSON(http.StatusOK, response.Error("failed to add pin"))
		}

		log.Info("pin added")

		ctx.IndentedJSON(http.StatusOK, response.Response{
			Status: response.StatusOK,
		})

		return
	}
}

func GetPin(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.GetPin"
	}
}

func GetAllPins(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		return
	}
}

func DeletePin(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.DeletePin"
		log.With(slog.String("op", op))

		var req request.DeletePinReq

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = s.DeletePin(req.PinID)
		if err != nil {
			if errors.Is(err, storage.ErrPinNotFound) {
				log.Info("pin not found")
				ctx.IndentedJSON(http.StatusOK, response.Error("pin not found"))

				return
			}
		}
		log.Info("pin deleted")

		ctx.IndentedJSON(http.StatusOK, response.Response{
			Status: response.StatusOK,
		})

		return
	}
}

// TODO: REFACTOR
func Login(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.Login"

		var req request.LoginReq

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = s.ValidatePassword(req.Username, req.Password)

		if errors.Is(err, storage.ErrInvalidPassword) {
			log.Info("invalid password")
			ctx.IndentedJSON(http.StatusOK, response.Error("invalid password"))

			return
		}

		if err != nil {
			log.Error("failed to login", err)
			ctx.IndentedJSON(http.StatusOK, response.Error("failed to login"))

			return
		}

		log.Info("user logged in")

		ctx.IndentedJSON(http.StatusOK, response.Response{
			Status: response.StatusOK,
		})

		return
	}
}

func Register(log *slog.Logger, s *postgres.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.Register"
		log.With(slog.String("op", op))

		var req request.RegisterReq

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = s.AddUser(req.Username, req.Password)

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
