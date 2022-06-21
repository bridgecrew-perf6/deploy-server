package returned

//
import (
	"encoding/json"
	"oneclick/server/Signature"
	"oneclick/server/pkg"
	"strconv"
)

var sign string

func SuccessWg(ret int, msg, peerInfo, FrpcConfig string) (res pkg.BindingReturned) {
	res = pkg.BindingReturned{
		Returned: pkg.Returned{
			Ret: ret,
			Msg: msg,
		},
		Request: pkg.Request{
			Sign: sign,
		},
		//Networklaundry: networklaundry,
		Data: struct {
			pkg.BindingReturnedConfig
		}{
			pkg.BindingReturnedConfig{
				PeerInfo:   peerInfo,
				FrpcConfig: FrpcConfig,
			},
		},
	}
	m, _ := json.Marshal(&res.Data)
	sign = Signature.SignatureMethod{}.RequestSucceeded(strconv.Itoa(ret), msg, string(m))
	res.Sign = sign
	return res
}

func ReturnUdp2raw(ret int, msg string, Udp2rawEndpoint int, Udp2rawServerIp string, Udp2rawPasswd string, peerInfo string, FrpcConfig string) (res pkg.BindingReturned) {
	res = pkg.BindingReturned{
		Returned: pkg.Returned{
			Ret: ret,
			Msg: msg,
		},
		Request: pkg.Request{
			Sign: sign,
		},
		//Networklaundry: networklaundry,
		Data: struct {
			pkg.BindingReturnedConfig
		}{
			pkg.BindingReturnedConfig{
				PeerInfo:        peerInfo,
				FrpcConfig:      FrpcConfig,
				Udp2rawEndpoint: Udp2rawEndpoint,
				Udp2rawServerIp: Udp2rawServerIp,
				Udp2rawPasswd:   Udp2rawPasswd,
			},
		},
	}
	m, _ := json.Marshal(&res.Data)
	sign = Signature.SignatureMethod{}.RequestSucceeded(strconv.Itoa(ret), msg, string(m))
	res.Sign = sign
	return res
}

func ReturnErrors(ret int, msg string) (res pkg.Errors) {
	res = pkg.Errors{
		Returned: pkg.Returned{
			Ret: ret,
			Msg: msg,
		},
		Request: pkg.Request{
			Sign: sign,
		},
	}
	sign = Signature.SignatureMethod{}.RequestSucceeded(strconv.Itoa(ret), msg, "")
	res.Sign = sign
	return res
}
