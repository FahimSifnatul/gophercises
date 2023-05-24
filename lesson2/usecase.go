package main

type UseCase interface {
	GetRedirectUrl(path string) (*string, error)
}

type useCase struct {
	yamlFile string
}

func NewUseCase(fileName string) UseCase {
	return &useCase{
		yamlFile: fileName,
	}
}

func (r *useCase) GetRedirectUrl(path string) (*string, error) {
	rData, err := toModelList(r.yamlFile)
	if err != nil {
		return nil, err
	}

	return getRedirectUrl(rData, path), nil
}
