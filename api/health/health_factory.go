package health

func MakeHealthController() *HealthController {
	healthService := NewHealthService()
	healthController := NewHealthController(healthService)

	return healthController
}
