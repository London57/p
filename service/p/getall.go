package pservice 

import (
	"context"

	db "github.com/London57/gsqlc/datagen"
)

type GetAll struct {
	db.Queries
}

func (GetAll) New(q db.Queries) GetAll {
	return GetAll{q}
}

func (s *GetAll) Exec(ctx context.Context) ([]db.P, error) {
	res, err := s.GetAllP(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

