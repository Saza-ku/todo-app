package domain

type Todo struct {
	ID          int
	Name        string
	Description string
	Status      Status
}

type Status string

const (
	NEW  Status = "new"
	WIP  Status = "wip"
	DONE Status = "done"
)

func (s Status) Validate() error {
	if s == NEW || s == WIP || s == DONE {
		return nil
	}
	return NewInvalidRequestError("Status should be new, wip or done")
}
