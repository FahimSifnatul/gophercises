package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func parseStoryScenarioByOption(name, option string) (*StoryScenario, error) {
	body, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var stories Story
	err = json.Unmarshal(body, &stories)
	if err != nil {
		return nil, err
	}

	if scenario, ok := stories[option]; ok {
		return &scenario, nil
	}
	return nil, nil
}

func parseHtmlTemplateFromScenario(sc *StoryScenario) (*string, error) {
	if sc == nil {
		return nil, nil
	}

	mn, err := os.ReadFile("templates/main.html")
	if err != nil {
		return nil, err
	}

	template := fmt.Sprintf(string(mn), sc.Title, strings.Join(sc.Story, "</p><p>"))
	if len(sc.Options) > 0 {
		op, err := os.ReadFile("templates/option.html")
		if err != nil {
			return nil, err
		}
		template += string(op)

		subOp, err := os.ReadFile("templates/sub_option.html")
		if err != nil {
			return nil, err
		}
		for _, o := range sc.Options {
			template += fmt.Sprintf(string(subOp), o.Arc, o.Text)
		}
	}
	return &template, nil
}
