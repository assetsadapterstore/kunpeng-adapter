package kunpeng

import (
	"encoding/hex"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {

	addrdec := tw.GetAddressDecoderV2()
	p2pk, _ := hex.DecodeString("48093514227a3482a658601d92fd3e3f320653ca")
	p2pkAddr, _ := addrdec.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)

	p2sh, _ := hex.DecodeString("1406b6c5e35c62b425c627369edcc615c5089ccc")
	p2shAddr, _ := addrdec.AddressEncode(p2sh, KPG_mainnetAddressP2SH)
	t.Logf("p2shAddr: %s", p2shAddr)
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	addrdec := tw.GetAddressDecoderV2()
	p2pkAddr := "KDn2zSFQehXrK9vnU56PvZp7cHif9p9wTo"
	p2pkHash, _ := addrdec.AddressDecode(p2pkAddr)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))

	p2shAddr := "PZkfhQs1RUuPW7Rngrk3TM87earb9e2qxx"

	p2shHash, _ := addrdec.AddressDecode(p2shAddr, KPG_mainnetAddressP2SH)
	t.Logf("p2shHash: %s", hex.EncodeToString(p2shHash))
}
