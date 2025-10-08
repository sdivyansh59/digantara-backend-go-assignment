package scheduler

//UPDATE jobs
//SET status = 'running'
//WHERE id = (
//SELECT id FROM jobs
//WHERE status = 'pending' AND run_at <= NOW()
//ORDER BY run_at ASC
//LIMIT 1
//FOR UPDATE SKIP LOCKED
//)
//RETURNING *;
