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
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
	"testing"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer_KPG(t *testing.T) {

	addrs := []string{
		"KGSFtQRbKUktn8CcwhFPMhTLqnRLeWzJwP",
		"KPFuKwuz2rxqs8ho2qsfRiT43kc4NSVtDL",
		"KPs6uAUqgTYm7H1TVpg5jSy8C3XhfC7jPE",
		"KHMyCcVigqiVuYa4r1Qkm2e5sR3kCXvzuw",
		"KCbW2aBmLmeqN4Aj1v7zuHffKtbBgMvTnh",
		"KL4psSu1727pribMx4q6i6y6CX21CDFhCi",

		//"KNbe8jVjripJ3cT4Dr6zMfQfwMRMLUoQSA",
		//"KSce8umEfJZLkpU4FxXiqQ794JjxMTQ7WA",
	}

	tm := testInitWalletManager()
	walletID := "WC8JhWE2UG4TTgzj1LdoVUCLhhrQ4XANqz"
	accountID := "HQmpz49FJw2uPPX5QYzQno8XvRWYF8g1JW67JTkYKGqw"
	//accountID := "9fj4QzdhJKj8YBxuFnZUG68MB8Xr4qaRcuYYrJ4zqxAf"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	for _, to := range addrs {

		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "3.65", "", nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

	}
}

func TestTransfer_QRC20(t *testing.T) {

	addrs := []string{
		"QRedeaK8D7qPQb6WvNnX1MV4enKKwFqmHk",
	}

	tm := testInitWalletManager()
	walletID := "WEqcj8FDLvf3uAS44ChEutM6oUbmgN23bf"
	accountID := "GVK6daCGmqKHfe2zEbpixarAJ9HEqawyAm9jFvmqU59Q"

	contract := openwallet.SmartContract{
		Address:  "0xf2033ede578e17fa6231047265010445bca8cf1c",
		Symbol:   "KPG",
		Name:     "QCASH",
		Token:    "QC",
		Decimals: 8,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.1", "", &contract)
		if err != nil {
			return
		}

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestSummary_KPG(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WC8JhWE2UG4TTgzj1LdoVUCLhhrQ4XANqz"
	accountID := "9fj4QzdhJKj8YBxuFnZUG68MB8Xr4qaRcuYYrJ4zqxAf"
	summaryAddress := "KSce8umEfJZLkpU4FxXiqQ794JjxMTQ7WA"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_QRC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WEqcj8FDLvf3uAS44ChEutM6oUbmgN23bf"
	accountID := "4gc8Ff4tiKFtPj6JzbAZqofyUHpbVHk3CpR21hTovS8T"
	summaryAddress := "QYV6cA236fyVKpM9fCFHBp8GCATW6sUF5a"

	contract := openwallet.SmartContract{
		Address:  "f2033ede578e17fa6231047265010445bca8cf1c",
		Symbol:   "KPG",
		Name:     "QCASH",
		Token:    "QC",
		Decimals: 8,
	}

	feesSupport := openwallet.FeesSupportAccount{
		AccountID: "GVK6daCGmqKHfe2zEbpixarAJ9HEqawyAm9jFvmqU59Q",
		//FixSupportAmount: "0.01",
		FeesSupportScale: "1.3",
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, &contract, &feesSupport)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
