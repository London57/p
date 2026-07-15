package pservice

import (
	"context"

	db "github.com/London57/gsqlc/datagen"
)

	type InsertOne struct {
		db.Queries
	}

	func (InsertOne) New(q db.Queries) InsertOne {
		return InsertOne{q}
	}

	func (s *InsertOne) Exec(ctx context.Context, args db.InsertOnePParams) (db.P, error) {
		res, err := s.InsertOneP(ctx, args)
		if err != nil {
			return db.P{}, err
		}
		return res, nil
	}
