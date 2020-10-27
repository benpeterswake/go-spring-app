package controller

import (
	"net/url"

	"gitlab.com/BPeters58/go-spring-app/pkg/model"
	"gitlab.com/BPeters58/go-spring-app/pkg/service"
	"gitlab.com/BPeters58/spring-go"
)

func init() {
	spring.AddController(ReservationController{})
}

type ReservationController struct {
	ReservationService service.ReservationService
}

func (rc ReservationController) CreateReservation() spring.Handler {
	return spring.Handler{
		Route: "/",
		Handler: func(requestBody *model.Reservation, vars map[string]string, params url.Values) error {
			return rc.ReservationService.SaveReservation(requestBody)
		},
		Method:   "POST",
		Produces: "application/json",
		Consumes: "application/json",
	}
}

func (rc ReservationController) GetReservationByID() spring.Handler {
	return spring.Handler{
		Route: "/{id}",
		Handler: func(requestBody *model.Reservation, vars map[string]string, params url.Values) (*model.Reservation, error) {
			id := vars["id"]
			return rc.ReservationService.GetReservation(id)
		},
		Method:   "GET",
		Produces: "application/json",
		Consumes: "application/json",
	}
}
