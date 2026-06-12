package response

type Code uint16

const (
	SuccessCode     Code = 0
	FailValidCode   Code = 1001
	FailServiceCode Code = 1002
)

func (c Code) String() (s string) {
	switch c {
	case SuccessCode:
		return "响应成功"
	case FailValidCode:
		return "校验失败"
	case FailServiceCode:
		return "服务异常"
	default:
		return ""
	}
}
