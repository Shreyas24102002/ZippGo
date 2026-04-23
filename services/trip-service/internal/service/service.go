package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TripRepository
}

func NewService (repo domain.TripRepository) domain.TripService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	trip := &domain.TripModel{
		ID : primitive.NewObjectID(),
		UserID:  fare.UserID,
		Status: "Pending",
		RideFare: *fare, 
	}
	return s.repo.CreateTrip(ctx, trip)
}