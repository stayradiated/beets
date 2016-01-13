package beets

type ByTrackNumber []Item

func (s ByTrackNumber) Len() int {
	return len(s)
}
func (s ByTrackNumber) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByTrackNumber) Less(i, j int) bool {
	return s[i].Track < s[j].Track
}
