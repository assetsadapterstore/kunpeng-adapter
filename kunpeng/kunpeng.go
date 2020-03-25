/*
 * Copyright 2019 The openwallet Authors
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

package kunpeng

import (
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/qtum-adapter/qtum"
	"github.com/blocktree/qtum-adapter/qtum/btcLikeTxDriver"
)

const (
	Symbol = "KPG"
)

var (
	KPGMainnetAddressPrefix = btcLikeTxDriver.AddressPrefix{[]byte{0x2d}, []byte{0x38}, "kp"}
	KPGTestnetAddressPrefix = btcLikeTxDriver.AddressPrefix{[]byte{0x78}, []byte{0x6E}, "tb"}
)

type WalletManager struct {
	*qtum.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = qtum.NewWalletManager()
	wm.Config = qtum.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.Decoder = NewAddressDecoder(&wm)
	wm.Config.MainNetAddressPrefix = KPGMainnetAddressPrefix
	wm.Config.TestNetAddressPrefix = KPGTestnetAddressPrefix
	return &wm
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "kunpeng"
}
