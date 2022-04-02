package domain

import "github.com/asaskevich/govalidator"

type Video struct {
	id string `valid:"uuid"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}

	return nil
}
