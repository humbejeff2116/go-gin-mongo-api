package controllers;

import (
	"context"
    "log"
	"go-gin-mongo-api/configs"
	"go-gin-mongo-api/models"
	"go-gin-mongo-api/responses"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type ProductController interface {
	GetProducts()
	CreateProduct()
	RemoveProduct()
	UpdateProduct()
}


var productCollection *mongo.Collection = configs.GetCollection(configs.MongodbClient, "golang", "products");
var validate = validator.New()

func CreateProduct(c *gin.Context){
    var response responses.ProductResponse;
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var product models.ProductModel
    defer cancel()

    //validate the request body
    err := c.BindJSON(&product);
    if  err != nil {
        // NOTE: return ErrorData to client only when in development
        response = responses.ProductResponse{
            Status: http.StatusBadRequest, 
            Message: "error", 
            ErrorData: map[string]interface{}{"data": err.Error()},
        } 
        c.JSON(http.StatusBadRequest, response)
        return
    }

    //validate required fields using validator library
    validationErr := validate.Struct(&product);
    if  validationErr != nil {
        // NOTE: return ErrorData to client only when in development
        response = responses.ProductResponse{
            Status: http.StatusBadRequest,
            Error: true, 
            Message: "error", 
            ErrorData: map[string]interface{}{"data": validationErr.Error()},
        }
        c.JSON(http.StatusBadRequest, response)
        return
    }
    // TODO... update newProduct struct properties
    newProduct := models.ProductModel{
        Id:       primitive.NewObjectID(),
        UserName:     product.UserName,
    }
  
    result, err := productCollection.InsertOne(ctx, newProduct)
    // NOTE: return ErrorData to client only when in development
    // TODO... log error and return a custom error to client
    if err != nil {
        response = responses.ProductResponse{
            Status: http.StatusInternalServerError, 
            Message: "error", 
            ErrorData: map[string]interface{}{ "data": err.Error() },
        }
        log.Fatal(err);
        c.JSON(http.StatusInternalServerError, response)
        return
    }
    response = responses.ProductResponse{
        Status: http.StatusCreated, 
        Message: "product created successfully", 
        Data: map[string]interface{}{"data": result },
    }
    c.JSON(http.StatusCreated, response)
}

// TODO... get products from database and send to client
func GetProducts(c *gin.Context) {
    var response responses.ProductResponse;
    var products []bson.D;
    // create a custom mongoDB context 
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
    defer cancel();
    // access products using a cursor(allows us to iterate over db while holding only a subset of them in memory at a given time)
    cursor, err := productCollection.Find(ctx, bson.D{});
   
    if err != nil {
        response = responses.ProductResponse{
            Status: http.StatusInternalServerError, 
            Error: true,
            ErrorData: map[string]interface{}{ "data": err.Error() },
            Message: "error occured while getting products from database",
        }
        log.Fatal(err);
        c.JSON(http.StatusInternalServerError, response);
        return
        // panic(err)
    }
    // close cursor to free resources it consumes in both the client application and the MongoDB server
    defer cursor.Close(ctx);
    //populate products array with all products query results
    err = cursor.All(ctx, &products);
    if  err != nil {
        response = responses.ProductResponse{
            Status: http.StatusInternalServerError,
            Error: true, 
            ErrorData: map[string]interface{}{ "data": err.Error() },
            Message: "error",   
        }
        log.Fatal(err);
        c.JSON(http.StatusInternalServerError, response);
        return
        // panic(err)
    }

    response = responses.ProductResponse{
        Status: http.StatusOK, 
        Message: "products gotten successfully", 
        Data: map[string]interface{}{"data": products },
    }
    c.JSON(http.StatusOK, response);
}