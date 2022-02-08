package db

import (
	"bookstore-oauth-api/src/clients/cassandra"
	"bookstore-oauth-api/src/domain/access_token"
	"bookstore-oauth-api/src/utils/errors"

	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct{}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError("error when trying to get current id")
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error when trying to save access token in database")
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error when trying to update current resource")
	}
	return nil
}
