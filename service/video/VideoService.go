package video

type VideoService interface {
	Feed(time int) bool
}
