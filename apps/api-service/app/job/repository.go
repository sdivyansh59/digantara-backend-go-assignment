package job

import (
	"context"
	"fmt"
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup/dbconfig"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/shared"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/database/crud"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/database/query"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"
	"github.com/uptrace/bun"
)

type IRepository interface {
	Filter(ctx context.Context, option ...query.SearchOption) ([]Job, error)
	Create(ctx context.Context, job *Job) error
	Update(ctx context.Context, job *Job) error
	GetByID(ctx context.Context, id snowflake.ID) (*Job, error)
	DeleteByID(ctx context.Context, job *Job) error
	GetNextJobToRun(ctx context.Context) (*Job, error)
	GetNextJobScheduledTime(ctx context.Context) (*int64, error)
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

func (r *Repository) Filter(ctx context.Context, option ...query.SearchOption) ([]Job, error) {
	return r.handler.Search(ctx, option...)
}

func (r *Repository) Create(ctx context.Context, job *Job) error {
	job.Id = r.snowflakeGenerator.Next()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

	return r.handler.Create(ctx, job)
}

func (r *Repository) Update(ctx context.Context, job *Job) error {
	job.UpdatedAt = time.Now()
	return r.handler.Update(ctx, job)
}

func (r *Repository) GetByID(ctx context.Context, id snowflake.ID) (*Job, error) {
	return r.handler.GetByID(ctx, id)
}

func (r *Repository) DeleteByID(ctx context.Context, job *Job) error {
	return r.handler.Delete(ctx, job)
}

func (r *Repository) GetNextJobScheduledTime(ctx context.Context) (*int64, error) {
	var nextRunAt int64
	sql := fmt.Sprintf("SELECT id, scheduled_at FROM job WHERE status = %s ORDER BY scheduled_at ASC LIMIT 1", shared.JobStatusScheduled)

	err := r.db.NewRaw(sql).Scan(ctx, &nextRunAt)
	if err != nil {
		return nil, err
	}

	return &nextRunAt, nil
}

func (r *Repository) GetNextJobToRun(ctx context.Context) (*Job, error) {
	getNextJob := fmt.Sprintf(`
  UPDATE job 
  SET status = '%s'
  WHERE id = (
   SELECT id FROM jobs
   WHERE status = '%s' AND run_at <= NOW()
   ORDER BY run_at ASC
   LIMIT 1
   FOR UPDATE SKIP LOCKED
  )
  RETURNING *`, shared.JobStatusRunning, shared.JobStatusScheduled)

	_, err := r.db.NewRaw(getNextJob).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
