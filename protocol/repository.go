package protocol

type Repositories string

const (
	MemRepo Repositories = "memMap"
)

type Repo interface {
	IsPrivateBoard(id string) (bool, string, error)
	AddPublicBoard(id string)
	AddPrivateBoard(id string, pw string)
	Subscribe(board string, subscription Subscribing) error
	GetSubscribers(board string) []Subscribing
	Unsubscribe(board string, user string)
}
