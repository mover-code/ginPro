/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-05-24 16:49:07
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-05-25 15:00:56
 * @FileName: log
 * @Desc:
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package logs

import (
    "fmt"
    "os"
    "path"
    "time"

    "github.com/gin-gonic/gin"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
)

var Log = logrus.New()
var baseDir, _ = os.Getwd()

// This function initializes a logrus logger with a JSON formatter, sets a log file output, sets the
// log level to info, and redirects gin framework logs to the same file.
// 
// Args:
//   paths (string): The `paths` parameter is a string that represents the path to the log file. It is
// used to create or open the log file and set it as the default output for the log.
// 
// Returns:
//   an error, which is either nil if the initialization of the log was successful or an error if there
// was a problem creating or opening the log file.
func InitLogrus(paths string) error {
    // 设置为json格式的日志
    Log.Formatter = &logrus.JSONFormatter{}
    fileName := path.Join(baseDir, paths)
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("创建日志文件/打开日志文件失败")
        return err
    }
    // 设置log默认文件输出
    Log.Out = file
    // gin.SetMode(gin.ReleaseMode)
    // gin框架自己记录的日志也会输出
    gin.DefaultWriter = Log.Out
    // 设置日志级别
    Log.Level = logrus.InfoLevel
    return nil
}

// The function creates a middleware for logging HTTP requests and responses in a specified file with
// rotation and hook mechanisms using the logrus package in Go.
// 
// Args:
//   paths (string): The file path where the log file will be created.
func LogMiddleWare(paths string) gin.HandlerFunc {
    // 日志文件
    fileName := path.Join(baseDir, paths)
    // 写入文件
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

    if err != nil {
        fmt.Println("打开/写入文件失败", err)
        return nil
    }
    // 实例化
    logger := logrus.New()
    // 日志级别
    logger.SetLevel(logrus.DebugLevel)
    // 设置输出
    logger.Out = file
    // 设置 rotatelogs,实现文件分割
    logWriter, err := rotatelogs.New(
        // 分割后的文件名称
        fileName+".%Y%m%d.log",
        // 生成软链，指向最新日志文件
        rotatelogs.WithLinkName(fileName),
        // 设置最大保存时间(7天)
        rotatelogs.WithMaxAge(7*24*time.Hour), //以hour为单位的整数
        // 设置日志切割时间间隔(1天)
        rotatelogs.WithRotationTime(1*time.Hour),
    )
    // hook机制的设置
    writerMap := lfshook.WriterMap{
        logrus.InfoLevel:  logWriter,
        logrus.FatalLevel: logWriter,
        logrus.DebugLevel: logWriter,
        logrus.WarnLevel:  logWriter,
        logrus.ErrorLevel: logWriter,
        logrus.PanicLevel: logWriter,
    }
    //给logrus添加hook
    logger.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    }))

    return func(c *gin.Context) {
        // 开始时间
        startTime := time.Now()

        // 处理请求
        c.Next()

        // 结束时间
        endTime := time.Now()

        // 执行时间
        latencyTime := endTime.Sub(startTime)

        // 请求方式
        reqMethod := c.Request.Method

        // 请求路由
        reqUri := c.Request.RequestURI

        // 状态码
        statusCode := c.Writer.Status()

        // 请求IP
        clientIP := c.ClientIP()

        // 日志格式
        logger.WithFields(logrus.Fields{
            "status_code":  statusCode,
            "latency_time": latencyTime,
            "client_ip":    clientIP,
            "req_method":   reqMethod,
            "req_uri":      reqUri,
        }).Info()
    }
}

// func init() {
//     err := initLogrus()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
// }
