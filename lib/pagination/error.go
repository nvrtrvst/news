package pagination

import "errors"

var (
	ErrorMaxPage     = errors.New("choosen page more than total page")
	ErrorPage        = errors.New("choosen have to more than total page")
	ErrorPageEmpty   = errors.New("choosen page more than total page")
	ErrorPageInvalid = errors.New("page invalid, must be nmumberic and more than 0")
)
