package service

import (
	"fmt"

	"github.com/victorhsb/go-mocker/model"
	"github.com/victorhsb/go-mocker/storage"
)

type MockSvc struct {
	storage storage.Interface
}

func NewService(storage storage.Interface) *MockSvc {
	return &MockSvc{storage: storage}
}

func (s *MockSvc) GetRoutes() ([]*model.Route, error) {
	files, err := s.storage.List("")
	if err != nil {
		return nil, fmt.Errorf("could not list routes; %w", err)
	}

	rts := make([]*model.Route, 0)
	for _, f := range files {
		r, err := s.storage.Read(f)
		if err != nil {
			if err == storage.ErrFormatNotSuported { // skip unsupported files
				fmt.Printf("skipping file %w as it's extension is not supported\n", f)
				continue
			}
			return nil, fmt.Errorf("could not read route %s; %w", f, err)
		}
		rts = append(rts, r)
	}

	return rts, nil
}
