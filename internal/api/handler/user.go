package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/musiku/global/errors"

	"github.com/Ndraaa15/musiku/global/response"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (h *Handler) Register(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	req := entity.UserRegister{}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	user, err := h.User.Register(&req, c)

	if err != nil {
		message = errors.ErrInternalServer.Error()
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to register user and please verify your email"
		data = user
	}
}

func (h *Handler) VerifyAccount(ctx *gin.Context) {
	param := ctx.Param("id")
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	uuid, err := uuid.FromString(param)
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	user, err := h.User.VerifyAccount(uuid, c)

	if err != nil {
		message = errors.ErrInternalServer.Error()
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Your account has been verified, please login"
		data = user
	}
}

func (h *Handler) Login(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	req := entity.UserLogin{}
	if err = ctx.ShouldBindJSON(&req); err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	res, err := h.User.Login(&req, c)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
		return
	default:
		message = "Success to login"
		data = res
	}
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	req := entity.UserUpdate{}
	if err = ctx.ShouldBindJSON(&req); err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	select {
	case <-c.Done():
		message = errors.ErrRequestTimeout.Error()
		code = http.StatusRequestTimeout
		return
	default:
		message = "Success to login"
		data = nil
	}
}

func (h *Handler) UploadPhotoProfile(ctx *gin.Context) {

}
