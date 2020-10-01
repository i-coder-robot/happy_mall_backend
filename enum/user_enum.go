package enum

type ResponseType int

const (
	OperateOk   ResponseType = 200
	OperateFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case OperateOk:
		return "Ok"
	case OperateFail:
		return "Fail"
	default:
		return "UNKNOWN"
	}
}
