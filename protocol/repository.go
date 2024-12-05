package protocol

type Repositories string

const (
	MemRepo Repositories = "memMap"
)

type Repo interface {
	IsPrivateBoard(id string) (bool, string, error)
	AddPublicBoard(id string) Board
	AddPrivateBoard(id string, pw string) Board
	Subscribe(board string, subscription Subscribing) error
	GetSubscribers(board string) []Subscribing
	Unsubscribe(board string, user string) Unsubscribe
}

type Board struct {
	Name          string
	Subscriptions map[string][]NotifyService // userId -> []notifiers
	Private       bool
	Password      string
}

type Unsubscribe struct {
	Board     string
	User      string
	Notifiers []NotifyService
}
