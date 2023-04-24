package collection_handler

import (
	"app/database"
	export_format "app/models/export"
	"app/models/menu/collections"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllCollections(c *fiber.Ctx) error {
	db := database.Database
	data := collections.GetMenu(db)
	c.Append("Access-Control-Allow-Origin", "*")
	return c.JSON(fiber.Map{"error": nil, "message": "Success", "status": false, "data": data})
}
func GetACollection(c *fiber.Ctx) error {
	db := database.Database
	collection := new(collections.Collection)
	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	result := collections.GetCollection(db, *collection)
	return c.JSON(fiber.Map{"error": nil, "data": result, "status": true})
}
func FormatCollection(c *fiber.Ctx) error {
	db := database.Database
	collection := new(collections.Collection)
	// res := new(export_format.Response)
	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	result := collections.GetCollection(db, *collection)

	res_result := export_format.Response{}
	// raw_op_req := export_format.RawLang{Language: "json"}
	// option_req := export_format.OptionsOfRequest{Raw: raw_op_req}

	collection_result := export_format.Collection1{Info: export_format.Info{Id: result.ID.String(), Name: result.Name, Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json", ExportId: int32(rand.Intn(99999999))}}
	if result.Variable != "" {
		decoded, err1 := base64.StdEncoding.DecodeString(result.Variable)
		if err1 != nil {
			return c.JSON(fiber.Map{"status": false})
		}

		err2 := json.Unmarshal(decoded, &collection_result.Variable)
		if err2 != nil {
			return c.JSON(fiber.Map{"status": false})
		}
	}
	for _, folder := range result.Folders {

		child_folder := export_format.Folder{Name: folder.Name}
		for _, req := range folder.Requests {

			url_req, _ := url.Parse(req.UriComponent)
			request_child := export_format.RequestOfCollection{}

			request_child.Name = req.Name
			request_child.Request.Method = req.Method

			if req.Header != "" {
				decoded, err1 := base64.StdEncoding.DecodeString(req.Header)
				if err1 != nil {
					return c.JSON(fiber.Map{"status": false})
				}

				err2 := json.Unmarshal(decoded, &request_child.Request.Header)
				if err2 != nil {
					return c.JSON(fiber.Map{"status": false})
				}
			}

			request_child.Request.Url.Raw = req.UriComponent
			request_child.Request.Url.Host = strings.Split(url_req.Hostname(), ".")
			request_child.Request.Url.Port = url_req.Port()
			request_child.Request.Url.Protocol = "http"
			request_child.Request.Url.Path = strings.Split(url_req.Path, "/")

			request_child.ProtocolProfileBehavior.BodyPruning = false
			request_child.Request.Header = [0]int{}
			request_child.Request.Url.Query = [0]int{}
			if req.Body != "" {
				request_child.Request.BodyOfRequest.Mode = "raw"
				request_child.Request.BodyOfRequest.Raw = req.Body
				request_child.Request.BodyOfRequest.Options.Raw.Language = "json"
				request_child.ProtocolProfileBehavior.BodyPruning = true
			}
			if req.Query != "" {
				decoded, err1 := base64.StdEncoding.DecodeString(req.Query)
				if err1 != nil {
					return c.JSON(fiber.Map{"status": false})
				}

				err2 := json.Unmarshal(decoded, &request_child.Request.Url.Query)
				if err2 != nil {
					return c.JSON(fiber.Map{"status": false})
				}
			}
			for _, r := range req.Responses {
				url_res := export_format.Url{}
				originalReq := export_format.OriginalRequest{}

				u, err := url.Parse(r.UriComponent)
				if err != nil {
					panic(err)
				}
				//url
				url_res.Raw = r.UriComponent
				url_res.Host = strings.Split(u.Hostname(), ".")
				url_res.Path = strings.Split(u.Path, "/")
				url_res.Port = u.Port()
				//originalreq
				originalReq.Header = ""
				originalReq.Method = r.Method

				originalReq.Url = url_res

				//response
				res_result.Body = r.Body
				res_result.Cookie = nil
				res_result.Header = nil
				res_result.Name = r.Name
				res_result.PreviewLanguage = "json"
				res_result.OriginalRequest = originalReq

				request_child.Response = append(request_child.Response, res_result)
			}

			child_folder.Item = append(child_folder.Item, request_child)

		}
		collection_result.Item = append(collection_result.Item, child_folder)
	}

	for _, req := range result.Requests {
		url_req, _ := url.Parse(req.UriComponent)
		request_child := export_format.RequestOfCollection{}

		request_child.Name = req.Name
		request_child.Request.Method = req.Method

		request_child.ProtocolProfileBehavior.BodyPruning = false
		request_child.Request.Url.Query = [0]int{}
		request_child.Request.Header = [0]int{}
		if req.Body != "" {
			request_child.Request.BodyOfRequest.Mode = "raw"
			request_child.Request.BodyOfRequest.Raw = req.Body
			request_child.Request.BodyOfRequest.Options.Raw.Language = "json"
			request_child.ProtocolProfileBehavior.BodyPruning = true
		}

		if req.Header != "" {

			decoded, err1 := base64.StdEncoding.DecodeString(req.Header)
			if err1 != nil {
				return c.JSON(fiber.Map{"status": false})
			}

			err2 := json.Unmarshal(decoded, &request_child.Request.Header)

			if err2 != nil {
				return c.JSON(fiber.Map{"status": false})
			}
		}

		request_child.Request.Url.Raw = req.UriComponent
		request_child.Request.Url.Host = strings.Split(url_req.Hostname(), ".")
		request_child.Request.Url.Port = url_req.Port()
		request_child.Request.Url.Protocol = "http"
		request_child.Request.Url.Path = strings.Split(url_req.Path, "/")

		if req.Query != "" {
			decoded, err1 := base64.StdEncoding.DecodeString(req.Query)
			if err1 != nil {
				return c.JSON(fiber.Map{"status": false})
			}

			err2 := json.Unmarshal(decoded, &request_child.Request.Url.Query)
			if err2 != nil {
				return c.JSON(fiber.Map{"status": false})
			}
		}
		for _, r := range req.Responses {
			url_res := export_format.Url{}
			originalReq := export_format.OriginalRequest{}

			u, err := url.Parse(r.UriComponent)
			if err != nil {
				panic(err)
			}
			//url
			url_res.Raw = r.UriComponent
			url_res.Host = strings.Split(u.Hostname(), ".")
			url_res.Path = strings.Split(u.Path, "/")
			url_res.Port = u.Port()
			//originalreq
			originalReq.Header = ""
			originalReq.Method = r.Method

			originalReq.Url = url_res

			//response
			res_result.Body = r.Body
			res_result.Cookie = nil
			res_result.Header = nil
			res_result.Name = r.Name
			res_result.PreviewLanguage = "json"
			res_result.OriginalRequest = originalReq

			request_child.Response = append(request_child.Response, res_result)
		}
		collection_result.Item = append(collection_result.Item, request_child)
	}
	// Giải mã mảng byte thành map[string][]responses.Response

	_, err := os.Stat("./storage/export")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./storage/export", 0755)
		if errDir != nil {
			return c.SendStatus(500)
		}

	}
	file, _ := json.MarshalIndent(collection_result, "", " ")
	file_path := "./storage/export/" + collection.ID.String() + ".json"
	_ = ioutil.WriteFile(file_path, file, 0644)

	return c.Download(file_path)
}
func CreateNewCollection(c *fiber.Ctx) error {
	db := database.Database
	collection := new(collections.Collection)
	c.Append("Access-Control-Allow-Origin", "*")
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}

	collection.CreatedAt = time.Now()
	collection.UpdatedAt = time.Now()
	collection.UserId = 1

	result := collections.CreateACollection(db, *collection)

	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The collection created successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The collection created don't successfully!", "status": false})

}
func UpdateCollection(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	collection := new(collections.Collection)
	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	collection.UpdatedAt = time.Now()

	db := database.Database
	result := collections.UpdateCollectionById(db, *collection)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data updated successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": "updated failed!", "message": "The data updated don't successfully!", "status": false})

}
func RemoveCollection(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	collection := new(collections.Collection)

	if err := c.BodyParser(&collection); err != nil {
		return c.JSON(fiber.Map{"error": err, "message": "The data not valid", "status": false})
	}
	db := database.Database
	data := collections.Collection{}
	data.ID = collection.ID
	result := collections.RemoveCollection(db, data)
	if result {
		return c.JSON(fiber.Map{"error": nil, "message": "The data deleted successfully!", "status": true})
	}
	return c.JSON(fiber.Map{"error": nil, "message": "The data don't successfully!", "status": false})
}
