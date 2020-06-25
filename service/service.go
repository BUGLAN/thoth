package service

type service struct {
}

type Service interface {
}

// 注入
func NewService() Service {
	return service{

	}
}

func (svc *service) XXX() {

}
