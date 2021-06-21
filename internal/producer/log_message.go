package producer

type ClassroomEventType = int

const (
	Created ClassroomEventType = iota
	Updated
	Removed
)

type ClassroomEvent struct {
	Type ClassroomEventType
	Body map[string]interface{}
}
