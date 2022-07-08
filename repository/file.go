package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	ErrFormatNotSuported = errors.New("this file extension is not supported")
)

type FileStorage struct {
	root string
}

func (f *FileStorage) Read(filename string) (*Route, error) {
	if filename[len(filename)-5:] != ".yaml" {
		return nil, ErrFormatNotSuported
	}

	file, err := ioutil.ReadFile(path.Join(f.root, filename))
	if err != nil {
		return nil, fmt.Errorf("could not read file %s; %w", filename, err)
	}

	filedir, filepath := path.Split(filename)
	filenameParts := strings.Split(path.Base(filepath), ".")
	route := &Route{
		Method: "GET",
		Path:   path.Join(filenameParts[:len(filenameParts)-1]...),
	}
	if len(filenameParts) > 2 {
		route = &Route{
			Method: strings.ToUpper(filenameParts[len(filenameParts)-2]), // before the last one
			Path:   path.Join(filenameParts[:len(filenameParts)-2]...),
		}
	}
	route.Path = path.Join(filedir, route.Path) // join with the file directory path

	err = yaml.Unmarshal(file, route)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config file; %w", err)
	}

	return route, nil
}

// list will return an array with all of the endpoints excluding the root
// it works like a tree lister
func (f *FileStorage) List(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(path.Join(f.root, dir))
	if err != nil {
		return nil, fmt.Errorf("could not read directory; %w", err)
	}

	dirs := make([]string, 0)
	for _, file := range files {
		p := path.Join(dir, file.Name())
		if file.IsDir() {
			list, err := f.List(p)
			if err != nil {
				return nil, fmt.Errorf("could not list directory %s; %w", p, err)
			}
			dirs = append(dirs, list...)
			continue
		}

		dirs = append(dirs, p)
	}

	return dirs, nil
}
