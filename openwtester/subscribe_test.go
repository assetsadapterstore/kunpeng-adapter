/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/common/file"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
	"path/filepath"
	"testing"
)

////////////////////////// 测试单个扫描器 //////////////////////////

type subscriberSingle struct {
}

//BlockScanNotify 新区块扫描完成通知
func (sub *subscriberSingle) BlockScanNotify(header *openwallet.BlockHeader) error {
	log.Notice("header:", header)
	return nil
}

//BlockTxExtractDataNotify 区块提取结果通知
func (sub *subscriberSingle) BlockExtractDataNotify(sourceKey string, data *openwallet.TxExtractData) error {
	log.Notice("account:", sourceKey)

	for i, input := range data.TxInputs {
		log.Std.Notice("data.TxInputs[%d]: %+v", i, input)
	}

	for i, output := range data.TxOutputs {
		log.Std.Notice("data.TxOutputs[%d]: %+v", i, output)
	}

	log.Std.Notice("data.Transaction: %+v", data.Transaction)

	return nil
}

//BlockExtractSmartContractDataNotify 区块提取智能合约交易结果通知
func (sub *subscriberSingle) BlockExtractSmartContractDataNotify(sourceKey string, data *openwallet.SmartContractReceipt) error {

	return nil
}

func TestSubscribeAddress(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "KPG"
		addrs      = map[string]string{
			"KCbW2aBmLmeqN4Aj1v7zuHffKtbBgMvTnh": "sender",
			"KGCFT74pa5pKHQaTuMB7ZgWjTyFaUMPYo4": "receiver",
		}
	)

	scanTargetFunc := func(target openwallet.ScanTargetParam) openwallet.ScanTargetResult {
		sourceKey, ok := addrs[target.ScanTarget]
		return openwallet.ScanTargetResult{SourceKey: sourceKey, Exist: ok, TargetInfo: nil,}
	}

	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	if scanner.SupportBlockchainDAI() {
		dbFilePath := filepath.Join("data", "db")
		dbFileName := "blockchain.db"
		file.MkdirAll(dbFilePath)
		dai, err := openwallet.NewBlockchainLocal(filepath.Join(dbFilePath, dbFileName), false)
		if err != nil {
			log.Error("NewBlockchainLocal err: %v", err)
			return
		}

		scanner.SetBlockchainDAI(dai)
	}
	scanner.SetRescanBlockHeight(162500)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanTargetFuncV2(scanTargetFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}
