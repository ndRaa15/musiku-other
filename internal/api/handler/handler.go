package handler

import (
	"github.com/Ndraaa15/musiku/internal/domain/service"
)

type Handler struct {
	User       service.UserServiceImpl
	Venue      service.VenueServiceImpl
	Instrument service.InstrumentServiceImpl
	Studio     service.StudioServiceImpl
}

func NewHandler(user service.UserServiceImpl, venue service.VenueServiceImpl, instrument service.InstrumentServiceImpl, studio service.StudioServiceImpl) *Handler {
	return &Handler{
		User:       user,
		Venue:      venue,
		Instrument: instrument,
		Studio:     studio,
	}
}

//indra
