package main

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryScenario struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]StoryScenario
