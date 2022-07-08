package repository

type Interface interface {
	Read(path string) (*Route, error)
	List(prefix string) ([]string, error)
}

// New gets a source and returns an Interface based on the type of source
// by now the only available input is file storage (starting with '/')
func New(source string) Interface {
	if source[0] == '/' { // file storage
		return &FileStorage{root: source}
	}
	return nil
}
