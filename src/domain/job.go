package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Job struct {
	ID           string    `valid:"uuid"`
	OutputBucket string    `valid:"notnull"`
	Status       string    `valid:"notnull"`
	Video        *Video    `valid:"-"`
	Error        string    `valid:"-"`
	CreatedDate  time.Time `valid:"-"`
	UpdatedDate  time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucket: output,
		Status: status,
		Video: video,
	}

	job.prepare()

	err := job.Validate()

	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedDate = time.Now()
	job.UpdatedDate = time.Now()
}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)
	if err != nil {
		return err
	}
	return nil
}
