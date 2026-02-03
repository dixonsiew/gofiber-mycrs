package dev

import (
    "mycrs/controller/dev"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/dev")
    api.Get("/enrollRoamingCert", dev.GetEnrollRoamingCert)
    api.Get("/generateResetPinCode", dev.GetGenerateResetPinCode)
    api.Get("/resetPin/:resetCode", dev.GetResetPin)
    api.Get("/requestRevoke/:requestCode", dev.GetRequestRevoke)
    api.Get("/certificateStatusIIH", dev.GetCertificateStatusIIH)
}
