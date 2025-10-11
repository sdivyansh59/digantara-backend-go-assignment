package job

import (
	"context"
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/database/crud"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/snowflake"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup/dbconfig"
	"github.com/uptrace/bun"
)

type JobFilter struct {
	Name        *string
	Status      *string
	ScheduledAt *int64
	CreatedBy   *string
}

type IRepository interface {
	Create(ctx context.Context, job *Job) error
	Update(ctx context.Context, job *Job) error
	GetByID(ctx context.Context, id snowflake.ID) (*Job, error)
	DeleteByID(ctx context.Context, job *Job) error
}

type Repository struct {
	db                 *bun.DB
	snowflakeGenerator *snowflake.Generator
	handler            *crud.Handler[Job, snowflake.ID]
}

func NewRepository(snowflakeGenerator *snowflake.Generator, jobSchedulerDB *dbconfig.JobSchedulerDB) IRepository {
	return &Repository{
		snowflakeGenerator: snowflakeGenerator,
		handler:            crud.NewHandler[Job, snowflake.ID](jobSchedulerDB.DB),
	}
}

func (r *Repository) Create(ctx context.Context, job *Job) error {
	job.id = r.snowflakeGenerator.Next()
	job.createdAt = time.Now()
	job.updatedAt = time.Now()

	return r.handler.Create(ctx, job)
}

func (r *Repository) Update(ctx context.Context, job *Job) error {
	job.updatedAt = time.Now()
	return r.handler.Update(ctx, job)
}

func (r *Repository) GetByID(ctx context.Context, id snowflake.ID) (*Job, error) {
	return r.handler.GetByID(ctx, id)
}

func (r *Repository) DeleteByID(ctx context.Context, job *Job) error {
	return r.handler.Delete(ctx, job)
}
