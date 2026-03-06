package main

import (
    "errors"
    "fmt"
    "html/template"
    "mycrs/cron"
    "mycrs/database"
    _ "mycrs/docs"
    "mycrs/router"
    "mycrs/utils"
    "os"

    "github.com/gofiber/contrib/fiberzerolog"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/compress"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/swagger"
    redoc "github.com/natebwangsut/fiber-redoc"
)

// @title Swagger MYCRS API
// @version 1.0
// @description This is a mycrs server.
// @BasePath /mycrs
func main() {
    defer utils.CatchPanic("main")
    utils.SetValidator()
    utils.SetSetting()
    runLogFile, _ := os.OpenFile(
        "app.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
    dLogFile, _ := os.OpenFile(
        "debug.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
    defer dLogFile.Close()
    defer runLogFile.Close()
    utils.SetDebugLogger(dLogFile)
    utils.SetLogger(runLogFile)
    port := utils.Setting.PORT
    cron.CronStatus = "STOP"

    if !fiber.IsChild() {
        cron.SetupCron(port)
        defer cron.ShutdownCron()
    }

    app := fiber.New(fiber.Config{
        Prefork: false,
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            var e *fiber.Error
            if errors.As(err, &e) {
                code = e.Code
            }

            return c.Status(code).JSON(fiber.Map{
                "statusCode": code,
                "message":    err.Error(),
            })
        },
    })
    app.Use(recover.New(recover.Config{
        EnableStackTrace: true,
    }))
    app.Use(compress.New())
    app.Use(fiberzerolog.New(fiberzerolog.Config{
        Logger: &utils.Logger,
    }))
    database.ConnectDB()
    defer database.CloseDB()

    basePath := "mycrs"
    initSwagger(app, basePath)
    router.SetupRoutes(app, basePath)

    err := app.Listen(fmt.Sprintf(":%s", port))

    if err != nil {
        utils.Logger.Fatal().Err(err).Msg("Fiber app error")
    }
}

func initSwagger(app *fiber.App, basePath string) {
    b, _ := os.ReadFile("./public/css/theme-feeling-blue.css")
    css := string(b)

    cfg := swagger.Config{
        URL:          "doc.json",
        DeepLinking:  true,
        DocExpansion: "list",
        Title:        "Swagger MYCRS API",
        SyntaxHighlight: &swagger.SyntaxHighlightConfig{
            Activate: true,
            Theme:    "arta",
        },
        CustomStyle: template.CSS(css),
    }

    app.Get(fmt.Sprintf("/%s/docs/*", basePath), swagger.New(cfg))
    app.Get(fmt.Sprintf("/%s/redocs/*", basePath), redoc.Handler)

    app.Static(fmt.Sprintf("/%s/static", basePath), "./public")
}
