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
	"encoding/hex"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	bech32, _ := hex.DecodeString("6ef8df14a7db46ad778b897477467f868b8e2c5c")
	bech32Addr, _ := tw.DecoderV2.AddressEncode(bech32)
	t.Logf("bech32Addr: %s", bech32Addr)

	p2sh, _ := hex.DecodeString("5c0a2143624c9f9cf410976ee24f608e916d69c9")
	p2shAddr, _ := tw.DecoderV2.AddressEncode(p2sh, addressEncoder.BTC_mainnetAddressP2SH)
	t.Logf("p2shAddr: %s", p2shAddr)

	p2pk, _ := hex.DecodeString("0bdb77a2ed8db95bb3af167cfeeb64fb27840bce")
	p2pkAddr, _ := tw.DecoderV2.AddressEncode(p2pk, SYS_mainnetAddressP2PKH)
	t.Logf("p2pkAddr: %s", p2pkAddr)
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	bech32Addr := "sys1qdmud7998mdr26aut3968w3nls69cutzuavnnc7"
	bech32Hash, _ := tw.DecoderV2.AddressDecode(bech32Addr)
	t.Logf("bech32Hash: %s", hex.EncodeToString(bech32Hash))

	p2shAddr := "3A5gE2q2ziDDLVs6UkBY2naSn2e9DhykBt"
	p2shHash, _ := tw.DecoderV2.AddressDecode(p2shAddr, SYS_mainnetAddressP2SH)
	t.Logf("p2shHash: %s", hex.EncodeToString(p2shHash))

	p2pkAddr := "SNNhNotxkqVq5PFV7FRiizrvMLzRECxNTJ"
	p2pkHash, _ := tw.DecoderV2.AddressDecode(p2pkAddr, SYS_mainnetAddressP2PKH)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))
}

