package pgx

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"projectname/internal/project/domain/configuration"
)

const (
	ServiceName = `DatabasePgxDriver`
)

type Driver struct {
	name string
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewDriver(ctx context.Context, cfg *viper.Viper) (*Driver, error) {
	connString := cfg.GetString(configuration.DatabaseConn)

	pool, err := pgxpool.Connect(ctx, connString)

	if err != nil {
		return nil, err
	}

	return &Driver{name: ServiceName, pool: pool, ctx: ctx}, nil
}

func (d Driver) Get(dst interface{}, query string, args ...interface{}) error {
	if err := pgxscan.Get(d.ctx, d.pool, dst, query, args...); err != nil {
		return err
	}
	return nil
}

func (d Driver) Select(dst interface{}, query string, args ...interface{}) error {
	if err := pgxscan.Select(d.ctx, d.pool, dst, query, args...); err != nil {
		return err
	}
	return nil
}

func (d Driver) Query(query string, args ...interface{}) error {
	rows, err := d.pool.Query(context.Background(), query, args...)

	if err != nil {
		return err
	}

	rows.Close()

	return nil
}

func (d Driver) Close() {
	d.pool.Close()
}

func (d Driver) Name() string {
	return d.name
}
