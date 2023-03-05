package usecase

import "strings"

type sampleService struct {
	sampleRepository SampleRepository
}

func NewSampleService(sr SampleRepository) sampleService {
	return sampleService{sampleRepository: sr}
}

func (ss *sampleService) Exec(text string) {
	r := strings.NewReader(text)
	ss.sampleRepository.Read(r)
}
