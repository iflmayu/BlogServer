package domain

type CommentStatus uint8

const (
	CommentStatusNormal  CommentStatus = 1
	CommentStatusDeleted CommentStatus = 2
)

func (s CommentStatus) String() string {
	switch s {
	case CommentStatusNormal:
		return "normal"
	case CommentStatusDeleted:
		return "deleted"
	default:
		return "normal"
	}
}
