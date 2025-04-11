package health

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (ctx *HealthService) HealthCheck() error {
	return nil
}
