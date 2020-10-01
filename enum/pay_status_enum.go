package enum

type PayStatus int

const (
	UnPay PayStatus = 0
	Payed PayStatus = 1
)

func (p PayStatus) String() string {
	switch p {
	case UnPay:
		return "未支付"
	case Payed:
		return "已支付"
	default:
		return "UNKNOWN"
	}
}
