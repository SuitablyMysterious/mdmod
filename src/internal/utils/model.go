package utils


type element struct {
	name string
	kind string
	
}
type manifest struct {
	name string
	description string
	version float32
	elements element
}