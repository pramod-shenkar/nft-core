package dao

type Daos struct {
	RequestDao *RequestDao
	TokenDao   *TokenDao
}

func NewDaos(
	RequestDao *RequestDao,
	TokenDao *TokenDao,
) *Daos {
	return &Daos{
		RequestDao: RequestDao,
		TokenDao:   TokenDao,
	}
}
