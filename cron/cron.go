package cron

import (
    "fmt"
    "mycrs/utils"

    "github.com/go-co-op/gocron/v2"
)

var (
    CronStatus string
    CronTask   gocron.Scheduler
    port       string
)

func CatchPanic(funcName string) {
    if err := recover(); err != nil {
        utils.LogInfo(fmt.Sprintf("recovered from panic -%s", funcName))
    }
}

func SetupCron(port string) {
    InitCron(port)
}

func InitCron(p string) {
    port = p
    s, _ := gocron.NewScheduler()
    /* _, _ = s.NewJob(
        gocron.DurationJob(5*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoGenerateLocalFile")
            DoGenerateLocalFile()
        }),
    )
    _, _ = s.NewJob(
        gocron.DurationJob(2*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoScanFile")
            DoScanFile()
        }),
    )
    _, _ = s.NewJob(
        gocron.DurationJob(2*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoVerifyRoamingCert")
            DoVerifyRoamingCert()
        }),
    )
    _, _ = s.NewJob(
        gocron.DurationJob(2*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoVerifyRoamingPin")
            DoVerifyRoamingPin()
        }),
    )
    _, _ = s.NewJob(
        gocron.DurationJob(2*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoSignFile")
            DoSignFile()
        }),
    )
    _, _ = s.NewJob(
        gocron.DurationJob(2*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoDownloadFile")
            DoDownloadFile()
        }),
    )
    _, _ = s.NewJob(
        gocron.DurationJob(2*time.Second),
        gocron.NewTask(func() {
            defer CatchPanic("DoBase64FromLocalFile")
            DoBase64FromLocalFile()
        }),
    ) */

    CronTask = s
    StartCron()
}

func StartCron() {
    CronStatus = "RUNNING"
    CronTask.Start()
}

func StopCron() {
    if CronTask != nil {
        CronTask.StopJobs()
    }

    CronStatus = "STOP"
}

func ShutdownCron() {
    if CronTask != nil {
        CronTask.Shutdown()
    }
}
