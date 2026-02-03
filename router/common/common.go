package common

import (
    "mycrs/controller/common"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    a := router.Group("/common")
    a.Get("/cron/start", common.StartCron)
    a.Get("/cron/stop", common.StopCron)
    a.Get("/cron/stat", common.CronStat)
}
