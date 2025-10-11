package job

type Controller struct {
	repository IRepository
}

func NewController(repository IRepository) *Controller {
	return &Controller{
		repository: repository,
	}
}
