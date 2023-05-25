/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-05-24 15:25:07
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-05-25 15:58:24
 * @FileName: config
 * @Desc:
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package config

import (
    "fmt"
    "io/ioutil"
    "time"

    "github.com/sirupsen/logrus"
    "gopkg.in/yaml.v2"
)

type Conf struct {
    // The above code defines a struct type Conf with various fields and their corresponding data types,
    // some of which have default values and optional tags.
    // @property {string} Host - The Host property specifies the IP address or hostname on which the server
    // should listen for incoming connections. The default value is "0.0.0.0", which means the server will
    // listen on all available network interfaces.
    // @property {int} Port - The Port property is an integer that specifies the port number on which the
    // server will listen for incoming connections.
    // @property {string} CertFile - The path to the certificate file for TLS encryption. It is marked as
    // optional in the JSON tag.
    // @property {string} KeyFile - The KeyFile property is a string that specifies the path to the private
    // key file for TLS encryption. It is marked as optional in the JSON format.
    // @property {bool} Verbose - Verbose is a boolean property that is marked as optional in both JSON and
    // YAML formats. It can be used to enable or disable verbose logging in the application. If set to
    // true, the application will log more detailed information about its operations. If set to false or
    // not provided, the application will log only
    // @property {int} MaxConns - MaxConns is a property of the Conf struct that specifies the maximum
    // number of connections allowed to the server. It has a default value of 10000.
    // @property {int64} MaxBytes - MaxBytes is a property of the Conf struct and is of type int64. It is
    // used to specify the maximum number of bytes that can be received in a single request. The default
    // value for this property is 1048576 (1MB).
    // @property {int64} Timeout - Timeout is a property that specifies the maximum time in milliseconds
    // that a request can take before it times out. If the request takes longer than the specified timeout,
    // it will be cancelled. The default value for Timeout is 3000 milliseconds.
    // @property {int64} CpuThreshold - CpuThreshold is a property of type int64 that represents the CPU
    // usage threshold in milliseconds. It has a default value of 900 and a range of 0 to 1000. This
    // property is used to set the maximum amount of CPU time that can be used by the application before it
    // is
    // @property {string} Mod - Mod is a string property that specifies the mode of the application. It has
    // a default value of "dev".
    // @property {string} LogsPath - The path where log files will be stored.
    // @property Chain - The `Chain` property is a nested struct that contains two string fields: `RPC` and
    // `Contract`. These fields likely represent endpoints or addresses for interacting with a blockchain
    // network or smart contract.
    // @property Mysql - This property is a struct that contains a single field called "DataSource". It is
    // used to specify the data source for a MySQL database connection.
    Host     string `json:",default=0.0.0.0" yaml:"Host"`
    Port     int    `yaml:"Port"`
    CertFile string `json:",optional"`
    KeyFile  string `json:",optional"`
    Verbose  bool   `json:",optional"`
    MaxConns int    `json:",default=10000" yaml:"MaxConns"`
    MaxBytes int64  `json:",default=1048576" yaml:"MaxBytes"`
    // milliseconds
    Timeout      int64  `json:",default=3000" yaml:"Timeout"`
    CpuThreshold int64  `json:",default=900,range=[0:1000]"`
    Mod          string `json:",default=dev" yaml:"Mod"`
    LogsPath     string `yaml:"LogsPath"`

    Chain struct {
        RPC      string `yaml:"RPC"`
        Contract string `yaml:"Contract"`
    }
    Mysql struct {
        DataSource   string        `yaml:"DataSource"`
        MaxIdleCount int           `yaml:"MaxIdleCount"`
        MaxOpenCount int           `yaml:"MaxOpenCount"`
        MaxLifeTime  time.Duration `yaml:"MaxLifeTime"`
    }

    Log *logrus.Logger
}

func NewConf(configFile string) (*Conf, error) {
    buf, err := ioutil.ReadFile(configFile)
    if err != nil {
        return nil, err
    }
    var conf Conf
    err = yaml.Unmarshal(buf, &conf)
    if err != nil {
        return nil, fmt.Errorf("in file %q: %v", configFile, err)
    }
    return &conf, nil
}
