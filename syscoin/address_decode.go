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
	"fmt"
	"github.com/blocktree/bitcoin-adapter/bitcoin"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	alphabet       = addressEncoder.BTCAlphabet
	bech32Alphabet = addressEncoder.BTCBech32Alphabet

	SYS_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x3f}, Suffix: nil}
	SYS_testnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x41}, Suffix: nil}
	SYS_mainnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0x80}, Suffix: []byte{0x01}}
	SYS_testnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0xef}, Suffix: []byte{0x01}}
	SYS_mainnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x05}, Suffix: nil}
	SYS_testnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0xc4}, Suffix: nil}
	SYS_mainnetAddressBech32V0      = addressEncoder.AddressType{EncodeType: "bech32", Alphabet: bech32Alphabet, ChecksumType: "sys", HashType: "h160", HashLen: 20, Prefix: []byte{0}, Suffix: nil}
	SYS_testnetAddressBech32V0      = addressEncoder.AddressType{EncodeType: "bech32", Alphabet: bech32Alphabet, ChecksumType: "tsys", HashType: "h160", HashLen: 20, Prefix: []byte{0}, Suffix: nil}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
	wm *WalletManager
}

//NewAddressDecoder 地址解析器
func NewAddressDecoder(wm *WalletManager) *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	decoder.wm = wm
	return &decoder
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := SYS_mainnetAddressBech32V0

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := SYS_mainnetAddressBech32V0

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	if len(hash) != cfg.HashLen {
		hash = owcrypt.Hash(hash, 0, owcrypt.HASH_ALG_HASH160)
	}

	address := addressEncoder.AddressEncode(hash, cfg)

	if dec.wm.Config.RPCServerType == bitcoin.RPCServerCore {
		//如果使用core钱包作为全节点，需要导入地址到core，这样才能查询地址余额和utxo
		err := dec.wm.ImportAddress(address, "")
		if err != nil {
			return "", err
		}
	}

	return address, nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	_, err := dec.AddressDecode(address, SYS_mainnetAddressP2PKH)
	if err == nil {
		return true
	}
	_, err = dec.AddressDecode(address, SYS_mainnetAddressP2SH)
	if err == nil {
		return true
	}
	_, err = dec.AddressDecode(address, SYS_mainnetAddressBech32V0)
	if err == nil {
		return true
	}
	return false
}

//ScriptPubKeyToBech32Address scriptPubKey转Bech32地址
func (dec *AddressDecoderV2) ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error) {

	return "", fmt.Errorf("ScriptPubKeyToBech32Address is not supported")

}
