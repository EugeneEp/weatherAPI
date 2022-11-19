package weather

import "errors"

var (
	ErrNotFound   = errors.New("city.not_found")
	ErrNotCreated = errors.New("city.not_created")
	ErrNotDeleted = errors.New("city.not_deleted")
	ErrNotWritten = errors.New("temp.not_written")
)
