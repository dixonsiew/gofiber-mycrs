package common

import (
	"mycrs/cron"

	"github.com/gofiber/fiber/v2"
)

// StartCron
//
// @Tags Common
// @Produce json
// @Success 200
// @Router /common/cron/start [get]
func StartCron(c *fiber.Ctx) error {
    cron.StartCron()
    return c.JSON(fiber.Map{
        "status": cron.CronStatus,
    })
}

// StopCron
//
// @Tags Common
// @Produce json
// @Success 200
// @Router /common/cron/stop [get]
func StopCron(c *fiber.Ctx) error {
    cron.StopCron()
    return c.JSON(fiber.Map{
        "status": cron.CronStatus,
    })
}

// CronStat
//
// @Tags Common
// @Produce json
// @Success 200
// @Router /common/cron/stat [get]
func CronStat(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "status": cron.CronStatus,
    })
}
