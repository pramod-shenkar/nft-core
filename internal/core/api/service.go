package api

type Services struct {
	RequestService *RequestService
}

func NewServices(
	RequestService *RequestService,

) *Services {
	return &Services{
		RequestService: RequestService,
	}
}
