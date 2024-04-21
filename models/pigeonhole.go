package models

import "pony-goblog/config"

type PigeonholeData struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category         `json:"categorys"`
	Lines     *map[string][]Post `json:"lines"`
}
