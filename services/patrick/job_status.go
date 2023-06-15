package patrick

// TODO: Shouldn't JobStatus just be typed string?
type JobStatus int

const (
	JobStatusFailed JobStatus = iota - 1
	JobStatusOpen
	JobStatusLocked
	JobStatusSuccess
	JobStatusCancelled
)

func (s JobStatus) String() string {
	switch s {
	case JobStatusFailed:
		return "Failed"
	case JobStatusOpen:
		return "Open"
	case JobStatusLocked:
		return "Locked"
	case JobStatusSuccess:
		return "Success"
	case JobStatusCancelled:
		return "Canceled"
	default:
		return "Unknown"
	}
}
