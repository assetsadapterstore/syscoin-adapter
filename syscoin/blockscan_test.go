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

package syscoin

import (
	"github.com/pborman/uuid"
	"testing"
)

func TestGetBTCBlockHeight(t *testing.T) {
	height, err := tw.GetBlockHeight()
	if err != nil {
		t.Errorf("GetBlockHeight failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlockHeight height = %d \n", height)
}


func TestBTCBlockScanner_GetCurrentBlockHeight(t *testing.T) {
	bs := tw.Blockscanner
	header, _ := bs.GetCurrentBlockHeader()
	t.Logf("GetCurrentBlockHeight height = %d \n", header.Height)
	t.Logf("GetCurrentBlockHeight hash = %v \n", header.Hash)
}

func TestGetBlockHeight(t *testing.T) {
	height, _ := tw.GetBlockHeight()
	t.Logf("GetBlockHeight height = %d \n", height)
}

func TestGetBlockHash(t *testing.T) {
	//height := GetLocalBlockHeight()
	hash, err := tw.GetBlockHash(1354918)
	if err != nil {
		t.Errorf("GetBlockHash failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlockHash hash = %s \n", hash)
}

func TestGetBlock(t *testing.T) {
	raw, err := tw.GetBlock("0000000000000000b4fa805f7110b3df79aa808413642d1a8ab7fcbbbc7af769")
	if err != nil {
		t.Errorf("GetBlock failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlock = %v \n", raw)
}

func TestGetTransaction(t *testing.T) {
	//f1ce2117e20fc64bfd314941ccde86a7d86203b53889550defb9c5a5035b55ed
	raw, err := tw.GetTransaction("1a11160cdbddf1e7596c2f159da584122b82f3a374020b26aedf37a8b1393c55")
	if err != nil {
		t.Errorf("GetTransaction failed unexpected error: %v\n", err)
		return
	}
	t.Logf("BlockHash = %v \n", raw.BlockHash)
	t.Logf("BlockHeight = %v \n", raw.BlockHeight)
	t.Logf("Blocktime = %v \n", raw.Blocktime)
	t.Logf("Fees = %v \n", raw.Fees)

	t.Logf("========= vins ========= \n")

	for i, vin := range raw.Vins {
		t.Logf("TxID[%d] = %v \n", i, vin.TxID)
		t.Logf("Vout[%d] = %v \n", i, vin.Vout)
		t.Logf("Addr[%d] = %v \n", i, vin.Addr)
		t.Logf("Value[%d] = %v \n", i, vin.Value)
	}

	t.Logf("========= vouts ========= \n")

	for i, out := range raw.Vouts {
		t.Logf("ScriptPubKey[%d] = %v \n", i, out.ScriptPubKey)
		t.Logf("Addr[%d] = %v \n", i, out.Addr)
		t.Logf("Value[%d] = %v \n", i, out.Value)
	}
}

func TestGetTxIDsInMemPool(t *testing.T) {
	txids, err := tw.GetTxIDsInMemPool()
	if err != nil {
		t.Errorf("GetTxIDsInMemPool failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetTxIDsInMemPool = %v \n", txids)
}

func TestBTCBlockScanner_scanning(t *testing.T) {

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "miphUAzHHeM1VXGSFnw6owopsQW3jAQZAk"

	//wallet, err := tw.GetWalletInfo(accountID)
	//if err != nil {
	//	t.Errorf("BTCBlockScanner_scanning failed unexpected error: %v\n", err)
	//	return
	//}

	bs := tw.Blockscanner

	//bs.DropRechargeRecords(accountID)

	bs.SetRescanBlockHeight(1384586)
	//tw.SaveLocalNewBlock(1355030, "00000000000000125b86abb80b1f94af13a5d9b07340076092eda92dade27686")

	//bs.AddAddress(address, accountID)

	bs.ScanBlockTask()
}

func TestBTCBlockScanner_Run(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
	)

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "mpkUFiXonEZriywHUhig6PTDQXKzT6S5in"

	//wallet, err := tw.GetWalletInfo(accountID)
	//if err != nil {
	//	t.Errorf("BTCBlockScanner_Run failed unexpected error: %v\n", err)
	//	return
	//}

	bs := tw.Blockscanner

	//bs.DropRechargeRecords(accountID)

	bs.SetRescanBlockHeight(1384586)

	//bs.AddAddress(address, accountID)

	bs.Run()

	<- endRunning

}

func TestWallet_GetRecharges(t *testing.T) {
	accountID := "WFvvr5q83WxWp1neUMiTaNuH7ZbaxJFpWu"
	wallet, err := tw.GetWalletInfo(accountID)
	if err != nil {
		t.Errorf("GetRecharges failed unexpected error: %v\n", err)
		return
	}

	recharges, err := wallet.GetRecharges(false)
	if err != nil {
		t.Errorf("GetRecharges failed unexpected error: %v\n", err)
		return
	}

	t.Logf("recharges.count = %v", len(recharges))
	//for _, r := range recharges {
	//	t.Logf("rechanges.count = %v", len(r))
	//}
}

func TestBTCBlockScanner_RescanFailedRecord(t *testing.T) {
	bs := tw.Blockscanner
	bs.RescanFailedRecord()
}

func TestFullAddress (t *testing.T) {

	dic := make(map[string]string)
	for i := 0;i<20000000;i++ {
		dic[uuid.NewUUID().String()] = uuid.NewUUID().String()
	}
}
