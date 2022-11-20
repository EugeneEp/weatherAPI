package database

import (
	stdLog "log"
	"projectname/internal/project"
	"projectname/internal/project/infrastructure/data"
)

var ctx data.Context

func init() {
	app := project.New()
	ctn, err := newCtn(app)
	if err != nil {
		stdLog.Fatalf("Filling test data error: %s", err)
	}

	ctx, err = data.New(ctn)
	if err != nil {
		stdLog.Fatalf("Filling test data error: %s", err)
	}
}
