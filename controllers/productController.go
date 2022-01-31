




package controllers;

import (
    "context"
    "go-gin-mongo-api/configs"
    "go-gin-mongo-api/models"
    "go-gin-mongo-api/responses"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)



type ProductController interface {
	GetProducts()
	CreateProduct()
	RemoveProduct()
	UpdateProduct()
}


var productCollection *mongo.Collection = configs.GetCollection(configs.MongodbClient, configs.SetServerConfigurations().DatabaseURI, "products")
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
func GetProducts (c *gin.Context) {

    var products = []models.ProductModel{
        
    }
    response := responses.ProductResponse{
        Status: http.StatusCreated, 
        Message: "product created successfully", 
        Data: map[string]interface{}{"data": products },
    }
    c.JSON(http.StatusOK, response);
}











// func CreateProduct() gin.HandlerFunc {
//     return func(c *gin.Context) {
// 		var response responses.ProductResponse;
//         ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//         var product models.ProductModel
//         defer cancel()

//         //validate the request body
//         err := c.BindJSON(&product);
//         if  err != nil {
//             response = responses.ProductResponse{
//                 Status: http.StatusBadRequest, 
//                 Message: "error", 
//                 Data: map[string]interface{}{"data": err.Error()},
//             } 
//             c.JSON(http.StatusBadRequest, response)
//             return
//         }

//         //use the validator library to validate required fields
//         validationErr := validate.Struct(&product);
//         if  validationErr != nil {
//             response = responses.ProductResponse{
//                 Status: http.StatusBadRequest,
//                 Error: true, 
//                 Message: "error", 
//                 Data: map[string]interface{}{"data": validationErr.Error()},
//             }
//             c.JSON(http.StatusBadRequest, response)
//             return
//         }

//         newProduct := models.ProductModel{
//             Id:       primitive.NewObjectID(),
//             UserName:     product.UserName,
//         }
      
//         result, err := productCollection.InsertOne(ctx, newProduct)
//         if err != nil {
// 			response = responses.ProductResponse{
// 				Status: http.StatusInternalServerError, 
// 				Message: "error", 
// 				Data: map[string]interface{}{ "data": err.Error() },
// 			}
//             c.JSON(http.StatusInternalServerError, response)
//             return
//         }
// 		response = responses.ProductResponse{
// 			Status: http.StatusCreated, 
// 			Message: "product created successfully", 
// 			Data: map[string]interface{}{ "data": result },
// 		}
//         c.JSON(http.StatusCreated, response)
//     }
// }