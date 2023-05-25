/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-05-25 14:41:37
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-05-25 16:06:13
 * @FileName: types
 * @Desc:
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package model

import (
    "fmt"
    "ginProjectTemplate/api/internal/config"
    "time"

    "github.com/sirupsen/logrus"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
)

var Orm *gorm.DB

type (
    // This line of code is defining a struct named `User` which will be used to represent a user in the
    // application. The struct contains fields for `Addr`, `RefererId`, `TeamNum`, and `InviteAmount`,
    // which are all attributes of a user. The `gorm.Model` field is embedded in the struct and provides
    // some default fields like `ID`, `CreatedAt`, `UpdatedAt`, and `DeletedAt` which are commonly used in
    // database models.
    User struct {
        gorm.Model
        Addr         string `json:"addr"`
        RefererId    int64  `json:"referer_id"`
        TeamNum      int64  `json:"team_num"`
        InviteAmount int64  `json:"invite_amount"`
    }

    // The `WithDrawLog` struct is defining a model for a withdrawal log in the application. It contains
    // fields for `Addr`, `Block`, and `Amount`, which are all attributes of a withdrawal log. The
    // `gorm.Model` field is embedded in the struct and provides some default fields like `ID`,
    // `CreatedAt`, `UpdatedAt`, and `DeletedAt` which are commonly used in database models.
    WithDrawLog struct {
        gorm.Model
        Addr   string `json:"addr"`
        Block  int64  `json:"block"`
        Amount int64  `json:"amount"`
    }
)

// This code is using the GORM library to open a connection to a MySQL database with the
// given credentials and configuration options. The `mysql.Open` function is used to specify
// the MySQL dialect and connection string, which includes the username, password, host,
// port, database name, and additional options such as character set and time zone. The
// `gorm.Config` struct is used to configure various options for the GORM library, such as
// the naming strategy for database tables (in this case, using singular table names) and the
// logger settings (in this case, logging errors only). The resulting `*gorm.DB` object is
// returned by the `OInit` function and can be used to perform database operations using
// GORM.

func Init(c *config.Conf) {

    slowLogger := logger.New(
        //设置Logger
        &MyWriter{mlog: c.Log},
        logger.Config{
            //慢SQL阈值
            SlowThreshold: time.Millisecond * 200,
            //设置日志级别，只有Warn以上才会打印sql
            LogLevel: logger.Warn,
        },
    )

    connect, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true, // 使用单数表名
        },
        Logger:          slowLogger,
        CreateBatchSize: 100,
    })
    if err != nil {
        panic("models: failed to connect database " + err.Error())
    } else {
        Orm = connect
    }
    connect.AutoMigrate(&User{}, &WithDrawLog{})
    sqlDB, _ := Orm.DB()
    sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleCount)
    sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenCount)
    sqlDB.SetConnMaxLifetime(c.Mysql.MaxLifeTime * time.Second)
}

type MyWriter struct {
    mlog *logrus.Logger
}

//实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
    logstr := fmt.Sprintf(format, v...)
    //利用loggus记录日志
    m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
    log := logrus.New()
    //配置logrus
    log.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })

    return &MyWriter{mlog: log}
}
