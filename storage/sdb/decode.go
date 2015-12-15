// Copyright 2015 Eleme Inc. All rights reserved.

package sdb

import (
	"fmt"
	"github.com/eleme/banshee/models"
)

// Decode db value into State.
func (db *DB) decodeValue(value []byte) (*models.State, error) {
	s := &models.State{}
	n, err := fmt.Sscanf(string(value), "%f:%f:%d", &s.Average, &s.StdDev, &s.Count)
	if err != nil {
		return nil, err
	}
	if n != 3 {
		return nil, ErrCorrupted
	}
	return s, nil
}
