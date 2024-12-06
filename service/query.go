package service

import (
	"app/config"
	"app/constant"
	"app/dto/request"
	"encoding/json"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type queryService[T any] struct {
	psql *gorm.DB
}

type QueryService[T any] interface {
	First(payload request.QueryReq[T]) (*T, error)
	Find(payload request.QueryReq[T]) ([]T, error)
	Create(data T) (*T, error)
	MultiCreate(datas []T) ([]T, error)
	Update(payload request.QueryReq[T]) (*T, error)
	Delete(payload request.QueryReq[T]) error
}

func (s *queryService[T]) First(payload request.QueryReq[T]) (*T, error) {
	var item *T
	var personOmit []string

	query := s.psql

	for key, omitChild := range payload.Omit {
		if len(omitChild) == 0 {
			personOmit = append(personOmit, key)
		}
	}

	for _, j := range payload.Joins {
		query = query.Joins(j)
	}

	for p, c := range payload.Preload {
		if c != nil {
			query.Preload(p, gorm.Expr(*c), func(tx *gorm.DB) *gorm.DB {
				return tx.Omit(payload.Omit[p]...)
			})
		} else {
			query.Preload(p, func(tx *gorm.DB) *gorm.DB {
				return tx.Omit(payload.Omit[p]...)
			})
		}
	}

	query = query.Where(payload.Condition, payload.Args...).Omit(personOmit...)

	err := query.First(&item).Error
	if err != nil {
		return nil, err
	}

	if payload.PreloadNull == constant.TRUE {
		return item, nil
	}

	jsonItem, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	var mapItem map[string]interface{}
	err = json.Unmarshal(jsonItem, &mapItem)
	if err != nil {
		return nil, err
	}

	for p, c := range payload.Preload {
		fields := strings.Split(p, ".")
		if c == nil {
			continue
		}

		var result map[string]interface{} = mapItem
		for _, f := range fields {
			f = strings.ToLower(string(f[0])) + f[1:]

			if result[f] == nil {
				return nil, nil
			}

			jsonData, err := json.Marshal(result[f])
			if err != nil {
				return nil, err
			}

			var converData map[string]interface{}
			err = json.Unmarshal(jsonData, &converData)
			if err != nil {
				return nil, err
			}

			result = converData
		}
	}

	return item, nil
}

func (s *queryService[T]) Find(payload request.QueryReq[T]) ([]T, error) {
	var list []T
	var personOmit []string

	query := s.psql

	for _, j := range payload.Joins {
		query = query.Joins(j)
	}

	for p, c := range payload.Preload {
		if c != nil {
			query = query.Preload(p, gorm.Expr(*c), func(tx *gorm.DB) *gorm.DB {
				return tx.Omit(payload.Omit[p]...)
			})
		} else {
			query = query.Preload(p, func(tx *gorm.DB) *gorm.DB {
				return tx.Omit(payload.Omit[p]...)
			})
		}
	}

	for key, omitChild := range payload.Omit {
		if len(omitChild) == 0 {
			personOmit = append(personOmit, key)
		}
	}
	query = query.Where(payload.Condition, payload.Args...).Omit(personOmit...)

	if payload.Order != "" {
		query = query.Order(payload.Order)
	}
	if payload.Limit != 0 {
		query = query.Limit(payload.Limit)
	}

	err := query.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *queryService[T]) Create(data T) (*T, error) {
	newData := data
	if err := s.psql.Create(&newData).Error; err != nil {
		return nil, err
	}
	return &newData, nil
}

func (s *queryService[T]) Update(payload request.QueryReq[T]) (*T, error) {
	newData := payload.Data

	err := s.psql.Where(payload.Condition, payload.Args...).Clauses(clause.Returning{}).Updates(&newData).Error
	if err != nil {
		return nil, err
	}

	return &newData, nil
}

func (s *queryService[T]) MultiCreate(datas []T) ([]T, error) {
	newDatas := datas
	if err := s.psql.Create(&newDatas).Error; err != nil {
		return []T{}, err
	}

	return newDatas, nil
}

func (s *queryService[T]) Delete(payload request.QueryReq[T]) error {
	var del T

	query := s.psql.Where(payload.Condition, payload.Args...)

	if payload.Unscoped {
		query = query.Unscoped()
	}

	if err := query.Delete(&del).Error; err != nil {
		return err
	}
	return nil
}

func NewQueryService[T any]() QueryService[T] {
	return &queryService[T]{
		psql: config.GetPsql(),
	}
}
