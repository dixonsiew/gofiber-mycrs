package roaming

import (
    "context"
    "encoding/xml"
    "mycrs/model"
    "mycrs/utils"
    "time"

    "github.com/flosch/pongo2/v6"
    "github.com/gofiber/fiber/v2"
)

// GetEnrollRoamingCert
//
// @Tags Roaming
// @Produce json
// @Success 200
// @Router /roaming/enrollRoamingCert [get]
func GetEnrollRoamingCert(c *fiber.Ctx) error {
    var result *model.EnrollRoamingCertResponse = new(model.EnrollRoamingCertResponse)
    url := utils.Setting.BASE_URL
    r := utils.GetR("EnrollRoamingCert")
    envelope :=
    `
    <soapenv:Envelope 
        xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:tem="http://tempuri.org/"
        xmlns:onec="http://schemas.datacontract.org/2004/07/OneCRSWcf"
        xmlns:arr="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
        <soapenv:Header/>
        <soapenv:Body>
            <tem:EnrollRoamingCert>
                <!--Optional:-->
                <tem:projectCodesetProfile>{{ projectCodesetProfile }}</tem:projectCodesetProfile>
                <!--Optional:-->
                <tem:productCode>{{ productCode }}</tem:productCode>
                <!--Optional:-->
                <tem:userProfile>
                    <!--Optional:-->
                    <onec:Add>{{ userProfile.Add }}</onec:Add>
                    <!--Optional:-->
                    <onec:BussinessType>{{ userProfile.BussinessType }}</onec:BussinessType>
                    <!--Optional:-->
                    <onec:Citizenship>{{ userProfile.Citizenship }}</onec:Citizenship>
                    <!--Optional:-->
                    <onec:City>{{ userProfile.City }}</onec:City>
                    <!--Optional:-->
                    <onec:CompAdd>{{ userProfile.CompAdd }}</onec:CompAdd>
                    <!--Optional:-->
                    <onec:CompCity>{{ userProfile.CompCity }}</onec:CompCity>
                    <!--Optional:-->
                    <onec:CompCountry>{{ userProfile.CompCountry }}</onec:CompCountry>
                    <!--Optional:-->
                    <onec:CompFaxno>{{ userProfile.CompFaxno }}</onec:CompFaxno>
                    <!--Optional:-->
                    <onec:CompName>{{ userProfile.CompName }}</onec:CompName>
                    <!--Optional:-->
                    <onec:CompPostcode>{{ userProfile.CompPostcode }}</onec:CompPostcode>
                    <!--Optional:-->
                    <onec:CompRegNo>{{ userProfile.CompRegNo }}</onec:CompRegNo>
                    <!--Optional:-->
                    <onec:CompState>{{ userProfile.CompState }}</onec:CompState>
                    <!--Optional:-->
                    <onec:Comptel>{{ userProfile.Comptel }}</onec:Comptel>
                    <!--Optional:-->
                    <onec:Country>{{ userProfile.Country }}</onec:Country>
                    <!--Optional:-->
                    <onec:Dob>{{ userProfile.Dob }}</onec:Dob>
                    <!--Optional:-->
                    <onec:Email>{{ userProfile.Email }}</onec:Email>
                    <!--Optional:-->
                    <onec:Faxno>{{ userProfile.Faxno }}</onec:Faxno>
                    <!--Optional:-->
                    <onec:Idno>{{ userProfile.Idno }}</onec:Idno>
                    <!--Optional:-->
                    <onec:Idtype>{{ userProfile.Idtype }}</onec:Idtype>
                    <!--Optional:-->
                    <onec:Name>{{ userProfile.Name }}</onec:Name>
                    <!--Optional:-->
                    <onec:Postcode>{{ userProfile.Postcode }}</onec:Postcode>
                    <!--Optional:-->
                    <onec:State>{{ userProfile.State }}</onec:State>
                    <!--Optional:-->
                    <onec:Telhome>{{ userProfile.Telhome }}</onec:Telhome>
                    <!--Optional:-->
                    <onec:Telmobile>{{ userProfile.Telmobile }}</onec:Telmobile>
                </tem:userProfile>
                <!--Optional:-->
                <tem:document>
                    <!--Zero or more repetitions:-->
                </tem:document>
                <!--Optional:-->
                <tem:PIN>{{ PIN }}</tem:PIN>
                <!--Optional:-->
                <tem:sendNotice>{{ sendNotice }}</tem:sendNotice>
            </tem:EnrollRoamingCert>
        </soapenv:Body>
    </soapenv:Envelope>
    `
    tpl, _ := pongo2.FromString(envelope)
    v, _ := tpl.Execute(pongo2.Context{
        "projectCodesetProfile": "Beta-MYCRS_AL8W6FGZ",
        "productCode":           "IHP0001-NEW",
        "userProfile": fiber.Map{
            "Name":          "DOCTOR IHP TEN",
            "Idtype":        "PP",
            "Idno":          "IHPDOCTOR10",
            "Email":         "uatuserone01@gmail.com",
            "Dob":           "22-09-1990",
            "Citizenship":   "MY",
            "Telmobile":     "0192222222",
            "Telhome":       "0444422222",
            "Faxno":         "",
            "Add":           "JALAN MACALISTER",
            "Postcode":      "10450",
            "City":          "GEORGE TOWN",
            "State":         "09",
            "Country":       "MY",
            "BussinessType": "C",
            "CompName":      "ISLAND HOSPITAL SDN. BHD.",
            "CompRegNo":     "199401038023",
            "Comptel":       "0444422222",
            "CompFaxno":     "",
            "CompAdd":       "JALAN MACALISTER",
            "CompCity":      "GEORGE TOWN",
            "CompState":     "09",
            "CompCountry":   "MY",
            "CompPostcode":  "10450",
        },
        "PIN":        "Pin123456@",
        "sendNotice": "1",
    })
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    res, err := r.SetContext(ctx).SetBody(v).Post(url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "err": err.Error(),
        })
    }

    bx := utils.GetXmlResult(res.String())
    if res.StatusCode() == 500 {
        var fx *model.Fault = new(model.Fault)
        err := xml.Unmarshal(bx, fx)
        if err != nil {
            return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
                "message": "No response data!",
            })
        }

        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": res.StatusCode(),
            "fault":  fx,
        })
    }

    err = xml.Unmarshal(bx, result)
    if err != nil {
        return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
            "message": "No response data!",
        })
    }

    return c.JSON(fiber.Map{
        "status": res.StatusCode(),
        "data":   result.EnrollRoamingCertResult,
    })
}

// GetGenerateResetPinCode
//
// @Tags Roaming
// @Produce json
// @Success 200
// @Router /roaming/generateResetPinCode [get]
func GetGenerateResetPinCode(c *fiber.Ctx) error {
    var result *model.GenerateResetPinCodeResponse = new(model.GenerateResetPinCodeResponse)
    url := utils.Setting.BASE_URL
    r := utils.GetR("GenerateResetPinCode")
    envelope :=
    `
    <soapenv:Envelope 
        xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:tem="http://tempuri.org/">
        <soapenv:Header/>
        <soapenv:Body>
            <tem:GenerateResetPinCode>
                <!--Optional:-->
                <tem:projectCodesetProfile>{{ projectCodesetProfile }}</tem:projectCodesetProfile>
                <!--Optional:-->
                <tem:userId>{{ userId }}</tem:userId>
                <!--Optional:-->
                <tem:email>{{ email }}</tem:email>
            </tem:GenerateResetPinCode>
        </soapenv:Body>
    </soapenv:Envelope>
    `
    tpl, _ := pongo2.FromString(envelope)
    v, _ := tpl.Execute(pongo2.Context{
        "projectCodesetProfile": "Beta-MYCRS_AL8W6FGZ",
        "userId":                "IHPDOCTOR10",
        "email":                 "uatuserone01@gmail.com",
    })
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    res, err := r.SetContext(ctx).SetBody(v).Post(url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "err": err.Error(),
        })
    }

    bx := utils.GetXmlResult(res.String())
    if res.StatusCode() == 500 {
        var fx *model.Fault = new(model.Fault)
        err := xml.Unmarshal(bx, fx)
        if err != nil {
            return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
                "message": "No response data!",
            })
        }

        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": res.StatusCode(),
            "fault":  fx,
        })
    }

    err = xml.Unmarshal(bx, result)
    if err != nil {
        return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
            "message": "No response data!",
        })
    }

    return c.JSON(fiber.Map{
        "status": res.StatusCode(),
        "data":   result.GenerateResetPinCodeResult,
    })
}

// GetResetPin
//
// @Tags Roaming
// @Produce json
// @Success 200
// @Router /roaming/resetPin [get]
func GetResetPin(c *fiber.Ctx) error {
    var result *model.ResetPinResponse = new(model.ResetPinResponse)
    url := utils.Setting.BASE_URL
    r := utils.GetR("ResetPin")
    envelope :=
    `
    <soapenv:Envelope 
        xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:tem="http://tempuri.org/">
        <soapenv:Header/>
        <soapenv:Body>
            <tem:ResetPin>
                <!--Optional:-->
                <tem:projectCodesetProfile>{{ projectCodesetProfile }}</tem:projectCodesetProfile>
                <!--Optional:-->
                <tem:userId>{{ userId }}</tem:userId>
                <!--Optional:-->
                <tem:resetCode>{{ resetCode }}</tem:resetCode>
                <!--Optional:-->
                <tem:newPIN>{{ newPIN }}</tem:newPIN>
            </tem:ResetPin>
        </soapenv:Body>
    </soapenv:Envelope>
    `
    tpl, _ := pongo2.FromString(envelope)
    v, _ := tpl.Execute(pongo2.Context{
        "projectCodesetProfile": "Beta-MYCRS_AL8W6FGZ",
        "userId":                "IHPDOCTOR10",
        "resetCode":             "2bce5b2d-44e4-4338-aa99-86cb3dbb860b",
        "newPIN":                "Pin123456@",
    })
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    res, err := r.SetContext(ctx).SetBody(v).Post(url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "err": err.Error(),
        })
    }

    bx := utils.GetXmlResult(res.String())
    if res.StatusCode() == 500 {
        var fx *model.Fault = new(model.Fault)
        err := xml.Unmarshal(bx, fx)
        if err != nil {
            return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
                "message": "No response data!",
            })
        }

        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": res.StatusCode(),
            "fault":  fx,
        })
    }

    err = xml.Unmarshal(bx, result)
    if err != nil {
        return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
            "message": "No response data!",
        })
    }

    return c.JSON(fiber.Map{
        "status": res.StatusCode(),
        "data":   result.ResetPinResult,
    })
}

// GetRequestRevoke
//
// @Tags Roaming
// @Produce json
// @Success 200
// @Router /roaming/requestRevoke [get]
func GetRequestRevoke(c *fiber.Ctx) error {
    var result *model.RequestRevokeResponse = new(model.RequestRevokeResponse)
    url := utils.Setting.BASE_URL
    r := utils.GetR("RequestRevoke")
    envelope :=
    `
    <soapenv:Envelope 
        xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:tem="http://tempuri.org/">
        <soapenv:Header/>
        <soapenv:Body>
            <tem:RequestRevoke>
                <!--Optional:-->
                <tem:projectCode>{{ projectCode }}</tem:projectCode>
                <!--Optional:-->
                <tem:userId>{{ userId }}</tem:userId>
                <!--Optional:-->
                <tem:requestnCode>{{ requestnCode }}</tem:requestnCode>
                <!--Optional:-->
                <tem:reasonCode>{{ reasonCode }}</tem:reasonCode>
                <!--Optional:-->
                <tem:reasonDesc>{{ reasonDesc }}</tem:reasonDesc>
                <!--Optional:-->
                <tem:docType>{{ docType }}</tem:docType>
                <!--Optional:-->
                <tem:authLetter>{{ authLetter }}</tem:authLetter>
            </tem:RequestRevoke>
        </soapenv:Body>
    </soapenv:Envelope>
    `
    tpl, _ := pongo2.FromString(envelope)
    v, _ := tpl.Execute(pongo2.Context{
        "projectCode":  "Beta-MYCRS_AL8W6FGZ",
        "userId":       "IHPDOCTOR10",
        "requestnCode": "IHP000048",
        "reasonCode":   "-2",
        "reasonDesc":   "test",
        "docType":      "pdf",
        "authLetter":   "",
    })
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    res, err := r.SetContext(ctx).SetBody(v).Post(url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "err": err.Error(),
        })
    }

    bx := utils.GetXmlResult(res.String())
    if res.StatusCode() == 500 {
        var fx *model.Fault = new(model.Fault)
        err := xml.Unmarshal(bx, fx)
        if err != nil {
            return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
                "message": "No response data!",
            })
        }

        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": res.StatusCode(),
            "fault":  fx,
        })
    }

    err = xml.Unmarshal(bx, result)
    if err != nil {
        return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
            "message": "No response data!",
        })
    }

    return c.JSON(fiber.Map{
        "status": res.StatusCode(),
        "data":   result.RequestRevokeResult,
    })
}

// GetCertificateStatusIIH
//
// @Tags Roaming
// @Produce json
// @Success 200
// @Router /roaming/certificateStatusIIH [get]
func GetCertificateStatusIIH(c *fiber.Ctx) error {
    var result *model.CertificateStatusIIHResponse = new(model.CertificateStatusIIHResponse)
    url := utils.Setting.BASE_URL
    r := utils.GetR("certificateStatusIIH")
    envelope :=
    `
    <soapenv:Envelope 
        xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:tem="http://tempuri.org/">
        <soapenv:Header/>
        <soapenv:Body>
            <tem:certificateStatusIIH>
                <!--Optional:-->
                <tem:sender>{{ sender }}</tem:sender>
                <!--Optional:-->
                <tem:idNo>{{ idNo }}</tem:idNo>
            </tem:certificateStatusIIH>
        </soapenv:Body>
    </soapenv:Envelope>
    `
    tpl, _ := pongo2.FromString(envelope)
    v, _ := tpl.Execute(pongo2.Context{
        "sender": "Beta-MYCRS_AL8W6FGZ",
        "idNo":   "IHPDOCTOR10",
    })
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    res, err := r.SetContext(ctx).SetBody(v).Post(url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "err": err.Error(),
        })
    }

    bx := utils.GetXmlResult(res.String())
    if res.StatusCode() == 500 {
        var fx *model.Fault = new(model.Fault)
        err := xml.Unmarshal(bx, fx)
        if err != nil {
            return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
                "message": "No response data!",
            })
        }

        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": res.StatusCode(),
            "fault":  fx,
        })
    }

    err = xml.Unmarshal(bx, result)
    if err != nil {
        return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
            "message": "No response data!",
        })
    }

    return c.JSON(fiber.Map{
        "status": res.StatusCode(),
        "data":   result.CertificateStatusIIHResult,
    })
}
