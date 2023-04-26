package mockapi

import (
	"app/database"
	"app/models/menu/collections"
	mock_server "app/models/mock-server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewMockApi(c *fiber.Ctx) error {

	c.Append("Access-Control-Allow-Origin", "*")
	db := database.Database
	collection := new(collections.Collection)
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	collection.CreatedAt = time.Now()
	collection.UpdatedAt = time.Now()
	collection.UserId = 1
	collection.IsServer = true
	result := collections.CreateACollection(db, *collection)

	if result {
		data := mock_server.InitJsonFile{}
		file, _ := json.MarshalIndent(data, "", " ")
		file_path := "./storage/" + collection.ID.String() + ".json"
		_ = ioutil.WriteFile(file_path, file, 0644)
		return c.JSON(fiber.Map{"error": nil, "message": "success", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "failed", "status": false})
}
func RemoveMockApi(c *fiber.Ctx) error {
	data := new(mock_server.MockResponse)
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db := database.Database
	collection := collections.Collection{ID: data.CollectionId}
	is_server := collections.CheckMockApi(db, collection)
	if is_server {
		method := c.Method()
		path := data.Path

		mockResponses := loadMockResponses(data.CollectionId.String())
		mockResponse, _ := mockResponses[method]
		flag := false
		for idx := range mockResponse {

			if mockResponse[idx].Path == path {
				mockResponses[method] = append(mockResponse[:idx], mockResponse[idx+1:]...)
				flag = true
				break
			}
		}

		return c.JSON(fiber.Map{"error": nil, "message": "Delete a api status", "status": flag})

	}
	return c.JSON(fiber.Map{"error": nil, "message": "Delete a api status", "status": false})
}
func CheckMock(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	id := c.Params("uuid")

	// Xử lý request ở đây
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Mock response not found")
	}

	db := database.Database
	collection := collections.Collection{ID: uuid}

	root_path := c.Path()
	tem_path := strings.Split(root_path, "/")
	path := "/" + strings.Join(tem_path[2:], "/")
	fmt.Println(path)
	method := c.Method()
	result := collections.CheckMockApi(db, collection)
	if result {
		mockResponses := loadMockResponses(uuid.String())
		mockResponse, _ := mockResponses[method]

		for idx := range mockResponse {
			if mockResponse[idx].Path == path {

				return c.Status(mockResponse[idx].StatusCode).JSON(mockResponse[idx].ResponseBody)

			}
		}

		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Mock response not found"})
	}
	return c.Status(http.StatusNotFound).SendString("Mock response not found")

}
func CreateOrUpdateApi(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Append("Access-Control-Allow-Origin", "*")
	data := new(mock_server.MockResponse)
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	db := database.Database
	collection := collections.Collection{ID: data.CollectionId}
	is_server := collections.CheckMockApi(db, collection)
	if is_server {
		method := c.Method()
		if method != data.Method {
			//fix bug url get  method
			method = data.Method
		}
		path := data.Path

		// Use type assertion to convert data to MockResponse

		var data1 interface{}
		err := json.Unmarshal([]byte(data.ResponseBody.(string)), &data1)
		if err != nil {
			c.SendStatus(404)
		}
		switch v := data1.(type) {
		case map[string]interface{}:
			data.ResponseBody = data1
		case []interface{}:
			data.ResponseBody = data1
		default:
			fmt.Println("It is something else:", v)
		}

		// Unmarshal the string into the slice
		// err := json.Unmarshal([]byte(data.ResponseBody.(string)), &arr)
		// if err != nil {
		// 	var arr map[string]interface{}
		// 	err := json.Unmarshal([]byte(data.ResponseBody.(string)), &arr)
		// }
		responseBody := data.ResponseBody

		// return c.JSON(arr)
		statusCode := data.StatusCode
		mockResponses := loadMockResponses(data.CollectionId.String())
		mockResponse, _ := mockResponses[method]
		flag := false
		for idx := range mockResponse {

			if mockResponse[idx].Path == path {
				mockResponse[idx].ResponseBody = responseBody
				mockResponse[idx].StatusCode = statusCode

				flag = true
				break
			}
		}
		if !flag {

			api := mock_server.MockResponse{Path: path, ResponseBody: responseBody, StatusCode: statusCode, CollectionId: data.CollectionId, Method: data.Method}

			mockResponse = append(mockResponse, api)
			mockResponses[method] = mockResponse
		}
		newFile, err := json.MarshalIndent(mockResponses, "", "    ")
		if err != nil {
			return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
		}
		file_path := "./storage/" + data.CollectionId.String() + ".json"
		err = ioutil.WriteFile(file_path, newFile, 0644)
		if err != nil {
			return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
		}
		return c.JSON(fiber.Map{"error": err, "message": "Update body by path successfully", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "Collection is not mock api server", "status": false})
}

func loadMockResponses(id string) map[string][]mock_server.MockResponse {
	file_name := "./storage/" + id + ".json"
	file, err := os.Open(file_name)

	if err != nil {

		log.Fatal(err)
	}

	mockResponses := make(map[string][]mock_server.MockResponse)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&mockResponses)

	if err != nil {
		log.Fatal(err)
	}

	return mockResponses
}
