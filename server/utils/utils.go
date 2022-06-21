package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const WireguardConfig = string(`[Interface]
PrivateKey = qILlkph/Z/l6iCefr08QnnvU2h9CGmWpsQBQlVOSb3Y=
Address = 10.188.10.10/24
DNS = 10.188.10.1
MTU = 1360

[Peer]
PublicKey = KPX513rslJ3mXeDE/B2EaF3FIBbAZXlVNZ++0gGBlhk=
Endpoint = 8.210.40.190:12786
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
`)

const FrpcConfig = string(`[common]
server_addr = 154.207.81.130
server_port = 7000
token = Hitosea@123

[ssh2]
type = tcp
local_ip = 127.0.0.1
local_port = 22
remote_port = 6000

[web2]
type = http
local_port = 8095
custom_domains = lts9527.com

log_file = /root/oneclickdeployment/log/frpc.log
log_level = info
log_max_days = 3
`)

// RandomString 随机字符串
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}