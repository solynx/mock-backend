package routers

import (
	// "app/cmd/controllers"

	"app/cmd/handler"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouters(app *fiber.App) {

	admin := app.Group("/admin", logger.New())
	// get all Exam
	admin.Get("/exam.json", handler.GetAllExam)
	admin.Post("/exam.json", handler.CreateExam)
	// //active Exam
	// admin.Post("/exam/actived.json", handler.ActiveExam)
	//C - create

	//R - read
	admin.Get("/exam/details.json", handler.GetExamById)
	// //U - update
	admin.Patch("/exam.json", handler.UpdateExam)
	//D - delete
	admin.Delete("/exam.json", handler.DeleteExam)
}
