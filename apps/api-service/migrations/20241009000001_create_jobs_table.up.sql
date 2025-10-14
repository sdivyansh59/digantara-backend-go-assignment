-- Create job table
CREATE TABLE IF NOT EXISTS job (
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL,
    interval_time BIGINT,
    scheduled_at BIGINT NOT NULL,
    last_run_at TIMESTAMP,
    successful_runs INTEGER NOT NULL DEFAULT 0,
    attributes JSONB,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create index on scheduled_at for efficient job scheduling queries
CREATE INDEX IF NOT EXISTS idx_job_scheduled_at ON job(scheduled_at);

-- Create index on status for filtering jobs by status
CREATE INDEX IF NOT EXISTS idx_job_status ON job(status);

