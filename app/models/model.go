package models

import "github.com/runningape/goblog/pkg/types"

type BaseModel struct {
	ID uint64
}

func (b BaseModel) GetStringID() string {
	return types.Uint64ToString(b.ID)
}
