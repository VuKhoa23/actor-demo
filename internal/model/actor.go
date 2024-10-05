package model

import (
	"time"
)

//---------------------------------------------------------------------------------//
// this file including models for actor (for database/request mapping and response)//
//---------------------------------------------------------------------------------//

type ActorRequest struct {
	ID        int64  `db:"actor_id"`
	FirstName string `json:"firstName" db:"first_name" validate:"required"`
	LastName  string `json:"lastName" db:"last_name" validate:"required"`
}

type Actor struct {
	ID        int        `db:"actor_id" json:"id"`
	FirstName string     `db:"first_name" json:"firstName"`
	LastName  string     `db:"last_name" json:"lastName"`
	UpdatedAt *time.Time `db:"last_update" json:"updatedAt,omitempty"`
}
