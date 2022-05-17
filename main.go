package main

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func main() {
	src := &net.IPNet{
		IP:   net.IPv4(172, 17, 0, 4),
		Mask: net.CIDRMask(int(16), 32),
	}
	dst := &net.IPNet{
		IP:   net.IPv4(172, 17, 0, 3),
		Mask: net.CIDRMask(int(16), 32),
	}
	policy := &netlink.XfrmPolicy{
		Src: src,
		Dst: dst,
		Dir: netlink.Dir(netlink.XFRM_DIR_OUT),
	}

	tunnelLeft := net.IPv4(10, 50, 41, 0)
	tunnelRight := net.IPv4(10, 50, 90, 0)

	tmpl := netlink.XfrmPolicyTmpl{
		Src:   tunnelLeft,
		Dst:   tunnelRight,
		Proto: netlink.XFRM_PROTO_ESP,
		Mode:  netlink.XFRM_MODE_TUNNEL,
		Reqid: 1,
	}

	policy.Tmpls = append(policy.Tmpls, tmpl)

	if err := netlink.XfrmPolicyAdd(policy); err != nil {
		fmt.Printf("error adding policy: %+v err: %v", policy, err)
	}
}
