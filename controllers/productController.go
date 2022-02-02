package controllers

import (
	"context"
	"fmt"
	"go-gin-mongo-api/configs"
	"go-gin-mongo-api/models"
	"go-gin-mongo-api/responses"
	"log"
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
            Error: true,
            Message: "JSON validation failed", 
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
            Message: "JSON format is incorrect", 
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
            Error: true,
            Message: "failed to create product", 
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

func GetProduct(c *gin.Context) {
    var response responses.ProductResponse;
    var product models.ProductModel
    userId := c.Param("productId")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(userId)

    err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&product)
    if err != nil {
        log.Fatal(err);
        response = responses.ProductResponse{
            Status: http.StatusInternalServerError,
            Error: true, 
            ErrorData: map[string]interface{}{ "data": err.Error() },
            Message: "An error occured while getting product",   
        }
        c.JSON(http.StatusInternalServerError, response)
        return
    }
    response = responses.ProductResponse{
        Status: http.StatusOK, 
        Message: "product gotten successfully", 
        Data: map[string]interface{}{"data": product },
    }

    c.JSON(http.StatusOK, response)
    
}


// find product with id and update product 
func UpdateProduct(c *gin.Context) {
   
    var response responses.ProductResponse;
    var updateProduct models.UpdateProductModel;
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    //validate the request body
    err := c.BindJSON(&updateProduct);
    if  err != nil {
        // NOTE: return ErrorData to client only when in development
        response = responses.ProductResponse{
            Status: http.StatusBadRequest, 
            Error: true,
            Message: "JSON format is incorrect", 
            ErrorData: map[string]interface{}{"data": err.Error()},
        } 
        c.JSON(http.StatusBadRequest, response)
        return
    }

    //validate required fields using validator library
    validationErr := validate.Struct(&updateProduct);
    if  validationErr != nil {
        // NOTE: return ErrorData to client only when in development
        response = responses.ProductResponse{
            Status: http.StatusBadRequest,
            Error: true, 
            Message: "JSON fields validation failed", 
            ErrorData: map[string]interface{}{"data": validationErr.Error()},
        }
        c.JSON(http.StatusBadRequest, response)
        return
    }
    objId, _ := primitive.ObjectIDFromHex(updateProduct.Id)
    filter := bson.M{"_id" : objId}
    update := bson.M{ "$set": bson.M{updateProduct.Key: updateProduct.Value}}

    result, err := productCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        response = responses.ProductResponse{
            Status: http.StatusBadRequest,
            Error: true, 
            Message: "failed to update product", 
            ErrorData: map[string]interface{}{"data": err.Error()},
        } 
        c.JSON(http.StatusBadRequest, response)
        return
    }
    response = responses.ProductResponse{
        Status: http.StatusOK, 
        Message: "product Updated sucessfully", 
        ErrorData: map[string]interface{}{"data": result},
    } 
    fmt.Printf("Documents matched: %v\n", result.MatchedCount)
    fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
    c.JSON(http.StatusOK, response)
}


func DeleteProduct(c *gin.Context) {
    var response responses.ProductResponse;
    productId := c.Param("productId")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(productId)

    result, err := productCollection.DeleteOne(ctx, bson.M{"_id": objId})
    if err != nil {
        response = responses.ProductResponse{
            Status: http.StatusBadRequest, 
            Error: true,
            Message: "failed to delete product", 
            ErrorData: map[string]interface{}{"data": err.Error()},
        }  
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    if result.DeletedCount < 1 {
        response =  responses.ProductResponse{
            Status: http.StatusNotFound,
            Error: true, 
            Message: "product with specified id not found",   
        }
        c.JSON(http.StatusNotFound, response)
        return
    }
    response =  responses.ProductResponse{
        Status: http.StatusOK, 
        Message: "product deleted successfully", 
    }
    
    c.JSON(http.StatusOK, response)
}