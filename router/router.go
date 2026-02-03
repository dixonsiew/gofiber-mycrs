package router

import (
    "fmt"
    "mycrs/router/common"
    "mycrs/router/dev"
    "mycrs/router/roaming"

    "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, basePath string) {
    api := app.Group(fmt.Sprintf("/%s", basePath))
    common.SetupRoutes(api)
    dev.SetupRoutes(api)
    roaming.SetupRoutes(api)
}
