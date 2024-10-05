package database

import "actor-demo/internal/model"

//------------------------------------------------//
// this file including repository for actor (CRUD)//
//------------------------------------------------//

func (s *service) CreateActor(request model.ActorRequest) (model.Actor, error) {
	query := `INSERT INTO actor(first_name, last_name) VALUES (:first_name, :last_name)`

	result, err := s.db.NamedExec(query, request)
	if err != nil {
		return model.Actor{}, err
	}

	actorId, err := result.LastInsertId()
	if err != nil {
		return model.Actor{}, err
	}

	actor, err := s.GetActorById(actorId)
	if err != nil {
		return model.Actor{}, err
	}

	return actor, nil
}

func (s *service) GetAllActor() ([]model.Actor, error) {
	query := `SELECT actor_id, first_name, last_name, last_update FROM actor`
	var result []model.Actor

	err := s.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetActorById(actorId int64) (model.Actor, error) {
	query := `SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = ?`
	var result model.Actor
	err := s.db.Get(&result, query, actorId)
	if err != nil {
		return model.Actor{}, err
	}

	return result, nil
}

func (s *service) UpdateActor(request model.ActorRequest) (model.Actor, error) {
	query := `UPDATE actor SET first_name = :first_name, last_name = :last_name WHERE actor_id = :actor_id`

	_, err := s.db.NamedExec(query, request)
	if err != nil {
		return model.Actor{}, err
	}

	actor, err := s.GetActorById(request.ID)
	if err != nil {
		return model.Actor{}, err
	}

	return actor, nil
}

func (s *service) DeleteActorById(actorId int64) error {
	if _, err := s.db.Exec("DELETE FROM actor WHERE actor_id = ?", actorId); err != nil {
		return err
	}
	return nil
}
