package routers

import (
	// "app/cmd/controllers"

	"app/cmd/handler"
	collection_handler "app/cmd/handler/collection"
	folder_handler "app/cmd/handler/folder"
	request_handler "app/cmd/handler/request"
	response_handler "app/cmd/handler/response"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouters(app *fiber.App) {

	admin := app.Group("/admin", logger.New())
	// get all Exam
	admin.Get("/exam.json", handler.GetAllExam)
	admin.Post("/exam.json", handler.CreateExam)
	//R - read
	admin.Get("/exam/details.json", handler.GetExamById)
	// //U - update
	admin.Patch("/exam.json", handler.UpdateExam)
	//D - delete
	admin.Delete("/exam.json", handler.DeleteExam)

	// ----------- COLLECTION
	admin.Get("/collection.json", collection_handler.GetAllCollections)
	admin.Post("/collection.json", collection_handler.CreateNewCollection)
	admin.Patch("/collection.json", collection_handler.UpdateCollection)
	admin.Delete("/collection.json", collection_handler.RemoveCollection)

	//------------ Folder
	admin.Post("/folder.json", folder_handler.CreateNewFolder)
	admin.Patch("/folder.json", folder_handler.UpdateFolder)
	admin.Delete("/folder.json", folder_handler.RemoveFolder)
	//----------- REQUEST
	admin.Post("/request.json", request_handler.CreateNewRequest)
	admin.Patch("/request.json", request_handler.UpdateRequest)
	admin.Delete("/request.json", request_handler.RemoveRequest)
	// ----------- RESPONSE
	admin.Post("/response.json", response_handler.CreateNewResponse)
	admin.Patch("/response.json", response_handler.Updateresponse)
	admin.Delete("/response.json", response_handler.RemoveResponse)
}
