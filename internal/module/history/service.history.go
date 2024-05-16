package history

import (
	"context"
	"golang-clean-architecture/domain"
)

type service struct {
	historyRepository domain.HistoryRepository
}

func NewService(historyRepository domain.HistoryRepository) domain.HistoryService {
	return &service{historyRepository: historyRepository}
}

// FindByID implements domain.HistoryService.
func (s *service) FindByID(ctx context.Context, history domain.History) domain.ApiResponse {
	history, err := s.historyRepository.FindByID(ctx, 1)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "success",
		Data:    history,
	}
}
