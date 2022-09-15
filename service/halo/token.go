package halo

import "github.com/duke-git/lancet/v2/netutil"

func (*ServiceHaloGroup) GetToken() {
	netutil.HttpPost("www.baidu.com")
}
