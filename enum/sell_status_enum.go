package enum

type SellStatus int

const (
	Selling  SellStatus = 0
	StopSell SellStatus = 1
)

func (p SellStatus) String() string {
	switch p {
	case Selling:
		return "销售中"
	case StopSell:
		return "停止销售"
	default:
		return "UNKNOWN"
	}
}
