package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/project104/controller"
)

func Setup(app *fiber.App) {

	app.Post("/api/register/", controller.Register)
	app.Post("/api/Login/", controller.Login)
	app.Post("/api/Login/Create/", controller.Create)
	app.Post("/api/Login/Show/", controller.Show)

	app.Post("/api/Login/Update", controller.Update)
	//	app.Post("/api/Login/Delete/", controller.Delete)
	app.Get("/api/User/", controller.User)
	app.Post("/api/Logut/", controller.Logout)

}
