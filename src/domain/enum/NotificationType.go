package enum

type NotificationType int

const (
	Like	NotificationType = iota + 1
	Dislike
	Comment
	Post
	Follow
	Story
)


