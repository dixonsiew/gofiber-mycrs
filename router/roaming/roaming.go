package roaming

import (
    "mycrs/controller/roaming"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/roaming")
    api.Get("/enrollRoamingCert", roaming.GetEnrollRoamingCert)
    api.Get("/generateResetPinCode", roaming.GetGenerateResetPinCode)
    api.Get("/resetPin", roaming.GetResetPin)
    api.Get("/requestRevoke", roaming.GetRequestRevoke)
    api.Get("/certificateStatusIIH", roaming.GetCertificateStatusIIH)
}
