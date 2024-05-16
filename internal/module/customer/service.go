package customer

import (
	"context"
	"golang-clean-architecture/domain"
	"time"
)

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{customerRepository: customerRepository}
}

func (s service) All(c context.Context) domain.ApiResponse {
	customer, err := s.customerRepository.FindAll(c)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	var CustomerData []domain.CustomerData
	for _, v := range customer {
		CustomerData = append(CustomerData, domain.CustomerData{
			ID:    v.ID,
			Name:  v.Name,
			Phone: v.Phone,
		})
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "success",
		Data:    CustomerData,
	}
}

func (s service) Save(ctx context.Context, customerData domain.CustomerData) domain.ApiResponse {
	customer := domain.Customer{
		Name:      customerData.Name,
		Phone:     customerData.Phone,
		CreatedAt: time.Now(),
	}

	err := s.customerRepository.Insert(ctx, &customer)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: "Something went wrong",
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "ok",
	}
}
