package hero

import "errors"

var (
	ErrWrongAttr       = errors.New("Wrong hero attribute")
	ErrWrongAttackType = errors.New("Wrong hero attack type")
	ErrWrongRole       = errors.New("Wrong hero role")
)
