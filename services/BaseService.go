package services

// BaseServiceError describe how error object is
type BaseServiceError struct {
	Field   string
	Message string
}

// BaseService is an abstract implementation for Services
type BaseService struct {
	Errors []BaseServiceError
}

// AddError permits add error information to a Service
func (service *BaseService) AddError(field string, message string) *BaseService {
	service.Errors = append(service.Errors, BaseServiceError{field, message})
	return service
}

// GetErrors fetchs list of errors attached to a Service
func (service *BaseService) GetErrors() []BaseServiceError {
	return service.Errors
}

// HasErrors checks if a Service has errors
func (service *BaseService) HasErrors() bool {
	if len(service.Errors) > 0 {
		return true
	}

	return false
}
