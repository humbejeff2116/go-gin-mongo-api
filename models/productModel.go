



package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ProductModel struct {
	Id primitive.ObjectID `json:"id,omitempty"`
    ProductName string `json:"productName,omitempty" validate:"required"`
    ProductCategory string `json:"productCategory,omitempty" validate:"required"`
    ProductPrice string `json:"productPrice,omitempty" validate:"required"`
    SellHistory []SellHistory `json:"sellHistory,omitempty"`
    ProductImages interface{} `json:"productImages,omitempty"`
}

type SellHistory struct {
    BuyerEmail, BuyerName, BuyerId string
    Buantity  int64
}

type UpdateProductModel struct {
    Id string `json:"id,omitempty" validate:"required"`
    Key string `json:"key,omitempty" validate:"required"`
    Value string `json:"value,omitempty" validate:"required"`
}