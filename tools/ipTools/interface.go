package ipTools

const (
	IPClASS_A IPCLASS = iota
	IPCLASS_B
	IPCLASS_C
	IPCLASS_D
	IPCLASS_E
)

type IPCLASS uint

type IpTool interface {
	IsIPv4() (bool, error)
	IsIPv6() (bool, error)
	WhichClass() (IPCLASS, error)
	IsReserved() (bool, error)
	IsUsed() (bool, error)
}
