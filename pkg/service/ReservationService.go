package service

import (
	"encoding/json"

	"gitlab.com/BPeters58/go-spring-app/pkg/model"
	"gitlab.com/BPeters58/go-spring-app/pkg/repository"
	"gitlab.com/BPeters58/spring-go"
)

func init() {
	spring.AddService(ReservationService{})
}

type ReservationService struct {
	ReservationRepo repository.ReservationRepository
}

func (rs ReservationService) SaveReservation(reservation *model.Reservation) error {
	return rs.ReservationRepo.Save(*reservation)
}

func (rs ReservationService) GetReservation(id string) (*model.Reservation, error) {
	res, err := rs.ReservationRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	jsonbody, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	reservation := model.Reservation{}
	if err := json.Unmarshal(jsonbody, &reservation); err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (rs ReservationService) Test() {
	rs.ReservationRepo.FindByReservationIDAndSomething("", "")
}
