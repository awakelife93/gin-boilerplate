package structures

type RequestItem struct {
	Item map[string]interface{}
}

type Api struct {
	Method   string
	Path     string
	Function func(RequestItem) (interface{}, *ErrorResult)
}

type Apis = []Api

type ErrorResult struct {
	Message *string
	Status  *int
	Data    interface{}
}

type Result struct {
	Message string
	Status  int
	Data    interface{}
}
