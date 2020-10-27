package repository

import (
	"gitlab.com/BPeters58/go-spring-app/pkg/model"
	"gitlab.com/BPeters58/spring-go"
)

func init() {
	spring.AddRepository(ReservationRepository{}, model.Reservation{})
}

type ReservationRepository struct {
	spring.CRUDRepository
	FindByReservationIDAndSomething func(id string, something string) *model.Reservation `query:"select * from reservation where id = ?0 and something = ?1"`
}
