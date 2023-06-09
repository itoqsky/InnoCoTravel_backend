package repository

import (
	"github.com/itoqsky/InnoCoTravel-backend/internal/core"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GetUserId(user core.User) (int, error)
	// LoginTg(core.User) (string, error)
}
type User interface {
	GetUserInfo(id int) (core.User, error)
}

type Trip interface {
	Create(trip core.Trip) (int, error)
	Delete(userId, tripId int) (int, error)
	GetById(userId, tripId int) (core.Trip, error)

	GetAdjTrips(input core.InputAdjTrips) ([]core.Trip, error)
	GetJoinedTrips(userId int) ([]core.Trip, error)
	GetJoinedUsers(userId, tripId int) ([]core.UserCtx, error)
}

type Repository struct {
	Authorization
	User
	Trip
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Trip:          NewTripPostgres(db),
	}
}
