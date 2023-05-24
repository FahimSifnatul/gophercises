package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func toModelList(name string) ([]*Redirect, error) {
	body, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	results := make([]*Redirect, 0)
	err = yaml.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func getRedirectUrl(rData []*Redirect, path string) *string {
	for _, r := range rData {
		if r.Path == path {
			return &r.Url
		}
	}
	return nil
}
