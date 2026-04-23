package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"ride-sharing/services/trip-service/internal/domain"
)

type inmemoRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInmemoRepository() domain.TripRepository {
	return &inmemoRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemoRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	if trip == nil {
		return nil, errors.New("trip is nil")
	}

	trip.ID = primitive.NewObjectID()
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
