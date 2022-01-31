



package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ProductModel struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	UserName string `json:"userName,omitempty" validate:"required"`
    UserId string `json:"userId,omitempty" validate:"required"`
    UserProfileImage string  `json:"userProfileImage,omitempty"`
    UserEmail string `json:"userEmail,omitempty" validate:"required"`
    ProductName string `json:"productName,omitempty" validate:"required"`
    ProductCategory string `json:"productCategory,omitempty" validate:"required"`
    ProductCountry string `json:"productCountry,omitempty" validate:"required"`
    ProductState string `json:"productState,omitempty" validate:"required"`
    ProductUsage string `json:"productUsage,omitempty"`
    ProductCurrency string `json:"productCurrency,omitempty" validate:"required"`
    ProductPrice string `json:"productPrice,omitempty" validate:"required"`
    ProductContactNumber string `json:"productContactNumber,omitempty" validate:"required"`
    // QuantitySold [{}]
    // productImages: [{}],
    // stars: [{}],
    // unstars: [{}],
    // comments: [{}],
    // interests: [{}],
    // createdAt: { type: Date, default: Date.now }
}