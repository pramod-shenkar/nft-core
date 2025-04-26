package grpc

type Handlers struct {
	Request *RequestHandler
}

func NewHandlers(Request *RequestHandler) *Handlers {
	return &Handlers{
		Request: Request,
	}
}
