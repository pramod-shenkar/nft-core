package api

type Services struct {
	RequestService *RequestService
	OnboardService *OnboardService
}

func NewServices(
	RequestService *RequestService,
	OnboardService *OnboardService,
) *Services {
	return &Services{
		RequestService: RequestService,
		OnboardService: OnboardService,
	}
}
