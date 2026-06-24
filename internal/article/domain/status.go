package domain

type ArticleStatus uint8

const (
	ArticleStatusDraft     ArticleStatus = 1 // 草稿
	ArticleStatusPublished ArticleStatus = 2 // 已发布
	ArticleStatusOffline   ArticleStatus = 3 // 已下线
)

func (s ArticleStatus) String() string {
	switch s {
	case ArticleStatusDraft:
		return "draft"
	case ArticleStatusPublished:
		return "published"
	case ArticleStatusOffline:
		return "offline"
	default:
		return "draft"
	}
}

func (s ArticleStatus) IsValid() bool {
	switch s {
	case ArticleStatusDraft, ArticleStatusPublished, ArticleStatusOffline:
		return true
	default:
		return false
	}
}
