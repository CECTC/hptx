/*
 * Copyright 2022 CECTC, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/cectc/dbpack/pkg/log"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/cectc/hptx/pkg/misc"
)

var distributedTransaction *DistributedTransaction

type DistributedTransaction struct {
	ApplicationID                    string `yaml:"appid" json:"appid"`
	RetryDeadThreshold               int64  `yaml:"retry_dead_threshold" json:"retry_dead_threshold"`
	RollbackRetryTimeoutUnlockEnable bool   `yaml:"rollback_retry_timeout_unlock_enable" json:"rollback_retry_timeout_unlock_enable"`

	EtcdConfig clientv3.Config `yaml:"etcd_config" json:"etcd_config"`

	ATConfig ATConfig `yaml:"at" json:"at,omitempty"`
	TMConfig TMConfig `yaml:"tm" json:"tm,omitempty"`
}

type ATConfig struct {
	DSN                 string        `yaml:"dsn" json:"dsn,omitempty"`
	ReportRetryCount    int           `default:"5" yaml:"reportRetryCount" json:"reportRetryCount,omitempty"`
	ReportSuccessEnable bool          `default:"false" yaml:"reportSuccessEnable" json:"reportSuccessEnable,omitempty"`
	LockRetryInterval   time.Duration `default:"10ms" yaml:"lockRetryInterval" json:"lockRetryInterval,omitempty"`
	LockRetryTimes      int           `default:"30" yaml:"lockRetryTimes" json:"lockRetryTimes,omitempty"`
}

type TMConfig struct {
	CommitRetryCount   int32 `default:"5" yaml:"commitRetryCount" json:"commitRetryCount,omitempty"`
	RollbackRetryCount int32 `default:"5" yaml:"rollbackRetryCount" json:"rollbackRetryCount,omitempty"`
}

// GetTMConfig return TMConfig
func GetTMConfig() TMConfig {
	if distributedTransaction == nil {
		panic("please init client config")
	}
	return distributedTransaction.TMConfig
}

// GetATConfig return ATConfig
func GetATConfig() ATConfig {
	if distributedTransaction == nil {
		panic("please init client config")
	}
	return distributedTransaction.ATConfig
}

// parseClientConfig parses an input configuration yaml document into a ClientConfig struct
//
// Environment variables may be used to override configuration parameters other than version,
// following the scheme below:
// DistributedTransaction.Abc may be replaced by the value of HPTX_ABC,
// DistributedTransaction.Abc.Xyz may be replaced by the value of HPTX_ABC_XYZ, and so forth
func parseClientConfig(rd io.Reader) (*DistributedTransaction, error) {
	in, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	p := misc.NewParser("HPTX")

	config := new(DistributedTransaction)
	err = p.Parse(in, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// InitDistributedTransaction init configuration from a file path
func InitDistributedTransaction(configurationPath string) *DistributedTransaction {
	fp, err := os.Open(configurationPath)
	if err != nil {
		log.Fatalf("open configuration file fail, %v", err)
	}

	config, err := parseClientConfig(fp)
	if err != nil {
		log.Fatalf("error parsing %s: %v", configurationPath, err)
	}

	defer func() {
		err = fp.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	distributedTransaction = config
	return distributedTransaction
}

func SetClientConfig(config *DistributedTransaction) {
	distributedTransaction = config
}
