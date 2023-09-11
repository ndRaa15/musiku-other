package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/global/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllInstrument(ctx *gin.Context) {
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

	instruments, err := h.Instrument.GetAllInstrument(c)
	if err != nil {
		message = errors.ErrInternalServer.Error()
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
		return
	default:
		message = "Success to get all instrument"
		data = instruments
	}
}

func (h *Handler) GetInstrumentByID(ctx *gin.Context) {
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

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	res, err := h.Instrument.GetByID(c, uint(id))
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
		message = "Success to get instrument by id"
		data = res
	}
}

func (h *Handler) RentInstrument(ctx *gin.Context) {
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

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
		return
	default:
		message = "Success to get instrument"
		data = nil
		return
	}
}

func (h *Handler) GetProvince(ctx *gin.Context) {
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

	idProvince := ctx.Query("id")
	res, err := h.Instrument.GetProvince(c, idProvince)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get province"
		data = res
	}
}

func (h *Handler) GetCity(ctx *gin.Context) {
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

	idProvince := ctx.Query("province")
	idCity := ctx.Query("id")
	res, err := h.Instrument.GetCity(c, idCity, idProvince)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get city"
		data = res
	}
}

func (h *Handler) GetCost(ctx *gin.Context) {
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

	origin := ctx.Query("origin") //city id
	weight := ctx.Query("weight")
	courier := ctx.Query("courier")

	res, err := h.Instrument.GetCost(c, origin, weight, courier)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get cost"
		data = res
	}
}
