package service

import (
	"app/config"
	"app/dto/request"

	"gorm.io/gorm"
)

type queryRawService[T any] struct {
	psql *gorm.DB
}

type QueryRawService[T any] interface {
	Query(payload request.QueryRawReq[T]) (*T, error)
	QueryAll(payload request.QueryRawReq[T]) ([]T, error)
}

func (s *queryRawService[T]) Query(payload request.QueryRawReq[T]) (*T, error) {
	var newData T

	condition := []interface{}{}
	condition = append(condition, payload.Data...)
	condition = append(condition, payload.Args...)
	err := s.psql.Raw(
		payload.Sql,
		condition...,
	).Scan(&newData).Error
	if err != nil {
		return nil, err
	}

	return &newData, nil
}

func (s *queryRawService[T]) QueryAll(payload request.QueryRawReq[T]) ([]T, error) {
	var newData []T

	condition := []interface{}{}
	condition = append(condition, payload.Data...)
	condition = append(condition, payload.Args...)
	err := s.psql.Raw(
		payload.Sql,
		condition...,
	).Scan(&newData).Error
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func NewQueryRawService[T any]() QueryRawService[T] {
	return &queryRawService[T]{
		psql: config.GetPsql(),
	}
}
