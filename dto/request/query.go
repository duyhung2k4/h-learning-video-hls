package request

import (
	"app/constant"
)

type QueryReq[T any] struct {
	Data        T                   `json:"data"`
	Datas       []T                 `json:"datas"`
	Args        []interface{}       `json:"args"`
	Condition   string              `json:"condition"`
	Preload     map[string]*string  `json:"preload"`
	Omit        map[string][]string `json:"omit"`
	Limit       int                 `json:"limit"`
	Joins       []string            `json:"joins"`
	Method      constant.METHOD     `json:"method"`
	Order       string              `json:"order"`
	Unscoped    bool                `json:"unscoped"`
	PreloadNull bool                `json:"preloadNull"`
}

type QueryRawReq[T any] struct {
	Data []interface{} `json:"data"`
	Args []interface{} `json:"args"`
	Sql  string        `json:"sql"`
}

type FindPayload struct {
	Condition string
	Preload   map[string]*string
	Omit      map[string][]string
	Order     string
	Agrs      []interface{}
}

type FirstPayload struct {
	Condition string
	Preload   map[string]*string
	Omit      map[string][]string
	Agrs      []interface{}
}
