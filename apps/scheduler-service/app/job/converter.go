package job

import "github.com/sdivyansh59/digantara-backend-golang-assignment/app/shared"

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ToDTO(entity *Job) *JobDTO {
	if entity == nil {
		return nil
	}

	return &JobDTO{
		ID:          entity.Id.String(),
		Name:        entity.Name,
		Description: entity.Description,
		Status:      entity.Status,
		ScheduledAt: entity.ScheduledAt,
		LastRunAt:   entity.LastRunAt,
		Attributes:  entity.Attributes,
		CreatedBy:   entity.CreatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (c *Converter) ToEntity(dto *CreateJobInput) *Job {
	if dto == nil {
		return nil
	}

	return &Job{
		Name:        dto.Name,
		Description: dto.Description,
		Status:      shared.JobStatusScheduled, // default status
		ScheduledAt: dto.ScheduledAt,
		Attributes:  dto.Attributes,
		CreatedBy:   dto.CreatedBy,
	}
}
