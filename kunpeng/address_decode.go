package kunpeng

import (
	"github.com/blocktree/go-owaddress"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/blocktree/qtum-adapter/qtum"
)

var (
	alphabet = addressEncoder.BTCAlphabet
)

var (

	//KPG stuff
	KPG_mainnetAddressP2PKH         = addressEncoder.AddressType{"base58", alphabet, "doubleSHA256", "h160", 20, []byte{0x2d}, nil}
	KPG_mainnetAddressP2SH          = addressEncoder.AddressType{"base58", alphabet, "doubleSHA256", "h160", 20, []byte{0x38}, nil}
	KPG_mainnetPrivateWIF           = addressEncoder.AddressType{"base58", alphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	KPG_mainnetPrivateWIFCompressed = addressEncoder.AddressType{"base58", alphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}
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

	cfg := KPG_mainnetAddressP2PKH

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

	cfg := KPG_mainnetAddressP2PKH

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)

	if dec.wm.Config.RPCServerType == qtum.RPCServerCore {
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
	valid, err := owaddress.Verify("kpg", address)
	if err != nil {
		return false
	}
	return valid
}
