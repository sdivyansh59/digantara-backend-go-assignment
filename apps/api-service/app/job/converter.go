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
		ID:          entity.id.String(),
		Name:        entity.name,
		Description: entity.description,
		Status:      entity.status,
		Interval:    entity.interval,
		ScheduledAt: entity.scheduledAt,
		LastRunAt:   entity.lastRunAt,
		Attributes:  entity.attributes,
		CreatedBy:   entity.createdBy,
		CreatedAt:   entity.createdAt,
		UpdatedAt:   entity.updatedAt,
	}
}

func (c *Converter) ToEntity(dto *CreateJobInput) *Job {
	if dto == nil {
		return nil
	}

	return &Job{
		name:        dto.Name,
		description: dto.Description,
		status:      shared.JobStatusPending, // default status
		interval:    dto.Interval,
		scheduledAt: dto.ScheduledAt,
		attributes:  dto.Attributes,
		createdBy:   dto.CreatedBy,
	}
}
