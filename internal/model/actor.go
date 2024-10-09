package model

import (
	"github.com/guregu/null/v5"
)

//---------------------------------------------------------------------------------//
// this file including models for actor (for database/request mapping and response)//
//---------------------------------------------------------------------------------//

type ActorRequest struct {
	ID        int64  `db:"actor_id"`
	FirstName string `json:"firstName" db:"first_name" binding:"required"`
	LastName  string `json:"lastName" db:"last_name" binding:"required"`
}

type Actor struct {
	ID        int         `db:"actor_id" json:"id"`
	FirstName null.String `db:"first_name" json:"firstName"`
	LastName  null.String `db:"last_name" json:"lastName"`
	UpdatedAt null.Time   `db:"last_update" json:"updatedAt"`
}

type HttpResponse[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data"`
}
