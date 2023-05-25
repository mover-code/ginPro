/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-05-24 14:54:09
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-05-25 15:55:54
 * @FileName: main
 * @Desc:  just do it
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package main

import (
    "context"
    "flag"
    "fmt"
    "ginProjectTemplate/api/internal/config"
    "ginProjectTemplate/api/model"
    "ginProjectTemplate/api/internal/router"

    cors "ginProjectTemplate/tools/cros"
    "ginProjectTemplate/tools/logs"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gin-gonic/gin"
)

var (
    configFile = flag.String("f", "etc/dev.yaml", "the config file")
    log        = logs.Log
    conf       *config.Conf
)

func init() {
    c, err := config.NewConf(*configFile)
    if err != nil {
        log.Fatalf("in file %v: %v", *configFile, err)
    }
    conf = c
    logs.InitLogrus(conf.LogsPath)
    conf.Log = log
    model.Init(conf)
}

func main() {
    gin.SetMode(conf.Mod)

    r := gin.Default()
    r.Use(logs.LogMiddleWare(conf.LogsPath), cors.Cors())
    router.Register(r)
    srv := &http.Server{
        Addr:    fmt.Sprintf(":%v", conf.Port),
        Handler: r,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Printf("listen: %s\n", err)
        }
    }()
    log.Printf("server start listen: http://127.0.0.1:%v/admin\n", conf.Port)

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }
    log.Println("Server exiting")

}
