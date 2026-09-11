package models

import "errors"

var (
	ErrWrongAttr       = errors.New("Wrong hero attribute")
	ErrWrongAttackType = errors.New("Wrong hero attack type")
	ErrWrongRole       = errors.New("Wrong hero role")
	ErrBadRank         = errors.New("Bad rank number")
	ErrZeroMatches     = errors.New("No matches played on this hero")
	ErrHeroNotFound    = errors.New("Unable to found hero")
)
