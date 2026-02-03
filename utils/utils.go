package utils

import (
    "fmt"
    "mycrs/config"
    "mycrs/model"
    "strconv"
    "strings"

    "os"

    "github.com/go-playground/validator/v10"
    "github.com/go-resty/resty/v2"
    "github.com/gofiber/fiber/v2"
    "github.com/rs/zerolog"
    "github.com/ztrue/tracerr"
)

type (
    ErrorResponse struct {
        Error       bool
        FailedField string
        Tag         string
        Value       any
    }

    XValidator struct {
        validator *validator.Validate
    }
)

var (
    validate     = validator.New()
    Setting      model.Setting
    Logger       zerolog.Logger
    iLogger      zerolog.Logger
    dLogger      zerolog.Logger
    appValidator *XValidator
    client       *resty.Client
)

func SetSetting() {
    client = resty.New()
    Setting = model.Setting{
        DB_URL:          config.Config("DB_URL"),
        DB_PORT:         config.Config("DB_PORT"),
        DB_SERVER:       config.Config("DB_SERVER"),
        DB_SERVICE:      config.Config("DB_SERVICE"),
        DB_USER:         config.Config("DB_USER"),
        DB_PASSWORD:     config.Config("DB_PASSWORD"),
        PORT:            config.Config("PORT"),
        BASE_URL:        config.Config("BASE_URL"),
        MAIL_HOST:       config.Config("MAIL_HOST"),
        MAIL_PORT:       config.Config("MAIL_PORT"),
        MAIL_USERNAME:   config.Config("MAIL_USERNAME"),
        MAIL_PASSWORD:   config.Config("MAIL_PASSWORD"),
        MAIL_REQUIRETLS: config.Config("MAIL_REQUIRETLS"),
        MAIL_SENDER:     config.Config("MAIL_SENDER"),
        MAIL_FROM_NAME:  config.Config("MAIL_FROM_NAME"),
        MAIL_APP_NAME:   config.Config("MAIL_APP_NAME"),
        MAIL_TO_ADDRESS: config.Config("MAIL_TO_ADDRESS"),
        PUSH_NTFY:       config.Config("PUSH_NTFY"),
    }
}

func SetLogger(runLogFile *os.File) {
    multi := zerolog.MultiLevelWriter(os.Stdout, zerolog.ConsoleWriter{Out: runLogFile})
    Logger = zerolog.New(multi).Level(zerolog.ErrorLevel).With().Timestamp().Caller().Logger()

    iLogger = zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func SetDebugLogger(dLogFile *os.File) {
    multi := zerolog.MultiLevelWriter(os.Stdout, zerolog.ConsoleWriter{Out: dLogFile})
    dLogger = zerolog.New(multi).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func SetValidator() {
    v := &XValidator{
        validator: validate,
    }
    appValidator = v
}

func GetValidator() *XValidator {
    return appValidator
}

func ValidatePayload(data any, c *fiber.Ctx) error {
    errs := GetValidator().Validate(data)
    if len(errs) > 0 && errs[0].Error {
        errMsgs := make([]string, 0)
        for _, err := range errs {
            errMsgs = append(errMsgs, fmt.Sprintf(
                "[%s]: '%v' | Needs to implement '%s'",
                err.FailedField,
                err.Value,
                err.Tag,
            ))
        }

        return &fiber.Error{
            Code:    fiber.ErrBadRequest.Code,
            Message: strings.Join(errMsgs, " and "),
        }
    }

    return nil
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
    validationErrors := []ErrorResponse{}

    errs := validate.Struct(data)
    if errs != nil {
        for _, err := range errs.(validator.ValidationErrors) {
            // In this case data object is actually holding the User struct
            var elem ErrorResponse

            elem.FailedField = err.Field() // Export struct field name
            elem.Tag = err.Tag()           // Export struct tag
            elem.Value = err.Value()       // Export field value
            elem.Error = true

            validationErrors = append(validationErrors, elem)
        }
    }

    return validationErrors
}

func ToString(s any) string {
    r := fmt.Sprintf("%v", s)
    switch v := s.(type) {
    case string:
        r = v
    case int:
        r = strconv.Itoa(v)
    default:
        r = fmt.Sprintf("%v", s)
    }

    return r
}

func SendNtfy() {
    pt := "Hospital"
    s := fmt.Sprintf("%s SFTP Failed (%s)", Setting.MAIL_APP_NAME, pt)
    url := fmt.Sprintf("https://ntfy.sh/%s", Setting.PUSH_NTFY)
    _, _ = client.R().SetHeader("Content-Type", "text/plain").SetBody(s).Post(url)
}

func GetErrors(errs []error) string {
    ls := []string{}
    for _, err := range errs {
        ls = append(ls, err.Error())
    }

    return strings.Join(ls, "|")
}

func GetR(action string) *resty.Request {
    return client.R().
        SetHeader("Content-Type", "text/xml; charset=utf-8").
        SetHeader("SOAPAction", fmt.Sprintf("http://tempuri.org/IRoamingCertificateEnrollNoLimit/%s", action))
}

func GetXmlResult(body string) []byte {
    i := strings.Index(body, "<s:Body>")
    j := strings.Index(body, "</s:Body>")
    content := body[8+i : j]
    content = strings.ReplaceAll(content, "&lt;", "<")
    content = strings.ReplaceAll(content, "&gt;", ">")
    bx := []byte(content)
    return bx
}

func CatchPanic(funcName string) {
    if err := recover(); err != nil {
        LogError(fmt.Errorf("recovered from panic -%s:%v", funcName, err))
    }
}

func LogError(err error) {
    if strings.Contains(err.Error(), "The process cannot access the file because it is being used by another process") ||
        strings.Contains(err.Error(), "The system cannot find the file specified") ||
        strings.Contains(err.Error(), "timeout") ||
        strings.Contains(err.Error(), "deadline"){
        return
    }

    ex := tracerr.Wrap(err)
    Logger.Err(err).Msg(tracerr.Sprint(ex))
}

func LogInfo(s string) {
    iLogger.Info().Msg(s)
}

func LogDebug(s string) {
    dLogger.Debug().Msg(s)
}
