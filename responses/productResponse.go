


package responses
import (
    "go-gin-mongo-api/models"
)


type ProductResponse struct {
    Status  int `json:"status"`
    Error bool  `json:"error"`
    ErrorData map[string]interface{} `json:"erroData"`
    Message string `json:"message"`
    // Data []models.ProductModel `json:"data"`
    Data map[string]interface{} `json:"data"`
}



type Res struct {
    Status  int `json:"status"`
    Error bool  `json:"error"`
    Message string `json:"message"`
    Data []models.ProductModel `json:"data"`
}
