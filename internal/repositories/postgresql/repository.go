package postgresql

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"gitlab.com/protocole/clearkey/internal/core/domain"
	"gitlab.com/protocole/clearkey/internal/core/ports/logger"
	"gitlab.com/protocole/clearkey/pkg/apperrors"
	"time"
)

type postgresql struct {
	db  *bun.DB
	ctx context.Context
}

func NewPostgreSQLRepository(user string, pass string, dbName string, addr string, insecure bool) (*postgresql, error) {
	ctx := context.Background()
	bun.SetLogger(logger.Log)
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithPassword(pass),
		pgdriver.WithUser(user),
		pgdriver.WithAddr(addr),
		pgdriver.WithDatabase(dbName),
		pgdriver.WithInsecure(insecure),
	))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	err := sqldb.Ping()
	if err != nil {
		logger.Log.Errorf("cannot connect %s", err.Error())
		return &postgresql{}, apperrors.DBConnectionFailed
	}

	db.RegisterModel((*domain.ClearKeyDecoded)(nil))
	_, err = db.NewCreateTable().IfNotExists().Model(new(domain.ClearKeyDecoded)).Exec(ctx)
	if err != nil {
		logger.Log.Errorf("cannot migrate %s", err.Error())
		return nil, err
	}

	return &postgresql{
		db:  db,
		ctx: ctx,
	}, nil
}

func (repo *postgresql) Get(id string) (domain.ClearKeyDecoded, error) {
	key := domain.ClearKeyDecoded{}
	ctx, cancelCtx := context.WithTimeout(repo.ctx, time.Second*10)
	defer func() { cancelCtx() }()

	if err := repo.db.NewSelect().Model(&key).Where("id = ?", id).Scan(ctx); err != nil {
		logger.Log.Debugf("GET %s - NOT FOUND", id[:8])
		return domain.ClearKeyDecoded{}, apperrors.NotFound
	}

	logger.Log.Debugf("GET %s - FOUND", id[:8])
	return key, nil
}

func (repo *postgresql) GetAll() (map[string]domain.ClearKeyDecoded, error) {
	keysRaw := make([]domain.ClearKeyDecoded, 0)
	keys := map[string]domain.ClearKeyDecoded{}
	ctx, cancelCtx := context.WithTimeout(repo.ctx, time.Second*10)
	defer func() { cancelCtx() }()

	if err := repo.db.NewSelect().Model(&keysRaw).Scan(ctx); err != nil {
		logger.Log.Debugf("GET ALL - CANNOT BE DONE")
		return map[string]domain.ClearKeyDecoded{}, apperrors.Internal
	}

	for _, keyDecoded := range keysRaw {
		keys[keyDecoded.Id.String()] = keyDecoded
	}

	return keys, nil
}

func (repo *postgresql) Save(key domain.ClearKeyDecoded) error {
	ctx, cancelCtx := context.WithTimeout(repo.ctx, time.Second*10)
	defer func() { cancelCtx() }()

	if _, err := repo.db.NewInsert().Model(&key).Exec(ctx); err != nil {
		logger.Log.Debugf("POST %s - FAILED", err)
		return err
	}

	logger.Log.Debug("POST %s - SAVED")
	return nil
}
