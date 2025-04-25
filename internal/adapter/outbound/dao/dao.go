package dao

type Daos struct {
	RequestDao *RequestDao
}

func NewDaos(
	RequestDao *RequestDao,
) *Daos {
	return &Daos{
		RequestDao: RequestDao,
	}
}
