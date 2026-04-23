package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // ex: We have a van, sedan or luxury car. This is the package slug that the user has selected for the trip
	TotalPriceInCents float64
}