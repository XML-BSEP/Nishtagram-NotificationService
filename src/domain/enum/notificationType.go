package enum

type NotificationType int

const (
	Like	= iota
	Dislike
	Comment
	Post
	Follow
	Story
)

func (g NotificationType) String() string {
	return [...]string{"Like", "Dislike", "Comment", "Post", "Follow", "Story"}[g]
}


