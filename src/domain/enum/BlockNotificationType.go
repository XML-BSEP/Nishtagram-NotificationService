package enum

type BlockNotificationType int

const (
	COMMENT	BlockNotificationType = iota + 1
	POST
	STORY
)

