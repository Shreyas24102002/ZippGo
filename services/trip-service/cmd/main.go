package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()

	inmemoRepo := repository.NewInmemoRepository()
	svc := service.NewService(inmemoRepo)

	fare := &domain.RideFareModel{
		UserID: "user123",
	}

	trip, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Trip created: %+v", trip)

	// Just for testing. Remove this: -
	for {
		time.Sleep(time.Second)
	}
}
