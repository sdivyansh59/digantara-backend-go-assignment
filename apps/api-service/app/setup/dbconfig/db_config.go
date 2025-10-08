package dbconfig

import (
	"github.com/rs/zerolog"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/utils"
	"github.com/uptrace/bun"
)

type JobSchedulerDB struct {
	*bun.DB
}

func ProvideJobSchedulerDB(logger, config *utils.DefaultConfig) *JobSchedulerDB {

}
