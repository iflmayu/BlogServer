package domain

type RoleType uint8

const (
	RoleAdmin RoleType = 1 // 管理员
	RoleUser  RoleType = 2 // 普通用户
	RoleGuest RoleType = 3 // 游客
)

func (r RoleType) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleUser:
		return "user"
	case RoleGuest:
		return "guest"
	default:
		return "user"
	}
}
