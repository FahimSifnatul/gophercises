package main

type UseCase interface {
	GetStory(option string) (*string, error)
}

type useCase struct {
	fileName string
}

func NewUseCase(name string) UseCase {
	return &useCase{
		fileName: name,
	}
}

func (s *useCase) GetStory(option string) (*string, error) {
	story, err := parseStoryScenarioByOption(s.fileName, option)
	if err != nil {
		return nil, err
	}
	return parseHtmlTemplateFromScenario(story)
}
