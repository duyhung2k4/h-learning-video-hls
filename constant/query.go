package constant

type METHOD string

const (
	GET          METHOD = "get"
	GET_ALL      METHOD = "get_all"
	CREATE       METHOD = "create"
	MULTI_CREATE METHOD = "multi_create"
	UPDATE       METHOD = "update"
	DELETE       METHOD = "delete"
)
