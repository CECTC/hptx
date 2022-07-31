# hptx
[![LICENSE](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/cectc/hptx/blob/main/LICENSE)
![GitHub Workflow Status](https://github.com/cectc/hptx/actions/workflows/main.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/cectc/hptx?status.svg)](https://godoc.org/github.com/cectc/hptx)
[![Go Report Card](https://goreportcard.com/badge/github.com/cectc/hptx)](https://goreportcard.com/report/github.com/cectc/hptx)
[![GitHub Release](https://img.shields.io/github/release-pre/cectc/hptx.svg)](https://github.com/cectc/hptx/releases)

High-performance non-intrusive distributed transaction solution, inspired by kubernetes, only for golang language.

___
## Features
+ Event driven, High-performance and non-intrusive
+ Support AT mode And TCC mode
+ Support check global lock in local transaction

## Requirements
+ Go 1.16 or higher.
+ ETCD(3+)
+ AT mode: Mysql (5.7+), MariaDB

___
## Installation
Simple install the package to your $GOPATH with the go tool from shell:
```shell
$ go get -u github.com/cectc/hptx
```
If you use AT mode to solve distributed transaction problems, you should also install the following package：
```shell
$ go get -u github.com/cectc/mysql
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's PATH.

## Usage
You should have your ETCD ready first. then, you can initialize hptx via `hptx.InitFromFile`:
```go
import (
	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/resource"
	"github.com/cectc/mysql"
)
  
//...

hptx.InitFromFile("${path of your config file}")
// If you use at mode, initial with following code 
mysql.RegisterResource(config.GetATConfig().DSN)
resource.InitATBranchResource(mysql.GetDataSourceManager())
```
It is also possible to set the configuration directly:
```go
import (
	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/resource"
	"github.com/cectc/mysql"
)
  
//...

// Fill in the fields as needed.
hptx.InitWithConf(&config.DistributedTransaction{
    ApplicationID:                    "",
    RetryDeadThreshold:               0,
    RollbackRetryTimeoutUnlockEnable: false,
    EtcdConfig:                       clientv3.Config{},
    ATConfig:                         config.ATConfig{},
    TMConfig:                         config.TMConfig{},
})
// If you use at mode, initial with following code 
mysql.RegisterResource(config.GetATConfig().DSN)
resource.InitATBranchResource(mysql.GetDataSourceManager())
```

Support GlobalLock:
```go
db.ExecContext(context.WithValue(context.Background(), "GlobalLock", true), sql, args...)
```
[Examples are available in our repos](https://github.com/CECTC/hptx-samples)

---
## License
hptx is licensed under the [GNU General Public License v3.0](https://github.com/CECTC/hptx/blob/main/LICENSE).
