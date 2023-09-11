package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ndraaa15/musiku/internal/api/handler"
	"github.com/Ndraaa15/musiku/internal/application/repository"
	"github.com/Ndraaa15/musiku/internal/application/service"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
	"github.com/Ndraaa15/musiku/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = iota
	ErrBadConfig
	ErrInternalServer
)

type server struct {
	router  *gin.Engine
	server  *http.Server
	handler *handler.Handler
}

func New() (*server, error) {
	s := &server{
		router: gin.Default(),
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	db, err := mysql.NewMySqlClient()
	if err != nil {
		log.Printf("[musiku-server] failed to initialize musiku database : %v\n", err)
		return nil, err
	}
	log.Printf("[musiku-server] succes to initialize musiku database. Database connected\n")

	if err := mysql.Migration(db); err != nil {
		log.Printf("[musiku-server] failed to migrate musiku database : %v\n", err)
		return nil, err
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	{
		var totalDays int64
		if err := db.Model(&entity.Day{}).Count(&totalDays).Error; err != nil {
			log.Printf("[musiku-server] failed to count total days : %v\n", err)
			return nil, err
		}

		if totalDays == 0 {
			if err := repository.SeedDays(db); err != nil {
				log.Printf("[musiku-server] failed to seed days : %v\n", err)
				return nil, err
			}
		}
	}

	{
		var totalVenue int64
		if err := db.Model(&entity.Venue{}).Count(&totalVenue).Error; err != nil {
			log.Printf("[musiku-server] failed to count total venue : %v\n", err)
			return nil, err
		}
		if totalVenue == 0 {
			if err := repository.SeedVenue(db); err != nil {
				log.Printf("[musiku-server] failed to seed venue : %v\n", err)
				return nil, err
			}
		}
	}

	venueRepository := repository.NewVenueRepository(db)
	venueService := service.NewVenueService(venueRepository)

	{
		var totalInstrument int64
		if err := db.Model(&entity.Instrument{}).Count(&totalInstrument).Error; err != nil {
			log.Printf("[musiku-server] failed to count total instrument : %v\n", err)
			return nil, err
		}

		if totalInstrument == 0 {
			if err := repository.SeedInstrument(db); err != nil {
				log.Printf("[musiku-server] failed to seed instrument : %v\n", err)
				return nil, err
			}
		}
	}

	instrumentRepository := repository.NewInstrumentRepository(db)
	instrumentService := service.NewInstrumentService(instrumentRepository)

	studioRepository := repository.NewStudioRepository(db)
	studioService := service.NewStudioService(studioRepository)

	s.handler = handler.NewHandler(userService, venueService, instrumentService, studioService)

	s.router = gin.Default()

	return s, nil
}

func Run() int {
	s, err := New()

	if err != nil {
		return ErrBadConfig
	}

	s.Start()
	s.router.Run(fmt.Sprintf(":%s", os.Getenv("CONFIG_SERVER_PORT")))
	return CodeSuccess
}

func (s *server) Start() {
	log.Printf("[musiku-server] Server is running at %s:%s", os.Getenv("CONFIG_SERVER_HOST"), os.Getenv("CONFIG_SERVER_PORT"))
	log.Println("[musiku-server] starting server...")

	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi, i'm musiku server"})
	})

	route := s.router.Group("/api/v1")

	route.POST("/register", s.handler.Register)
	route.POST("/login", s.handler.Login)
	route.GET("/verify/:id", s.handler.VerifyAccount)

	route.GET("/venue", s.handler.GetAllVenue)
	route.GET("/venue/:id", s.handler.GetVenueByID)
	route.PATCH("/venue/:id", s.handler.RentVenue)

	route.GET("/instruments", s.handler.GetAllInstrument)
	route.GET("/instruments/:id", s.handler.GetInstrumentByID)
	route.PATCH("/instruments/:id", s.handler.RentInstrument)
	route.GET("/rent/instrument/province", s.handler.GetProvince)
	route.GET("/rent/instrument/city", s.handler.GetCity)
	route.GET("/rent/instrument/cost", s.handler.GetCost)

	route.Use(middleware.ValidateJWTToken())
}
