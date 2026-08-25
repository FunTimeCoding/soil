package log_manager

import "time"

type Log struct {
	FileName                    string      `json:"FileName"`
	Players                     []Player    `json:"Players"`
	EncounterResult             int         `json:"EncounterResult"`
	EncounterMode               int         `json:"EncounterMode"`
	Encounter                   int         `json:"Encounter"`
	MapIdentifier               int         `json:"MapId"`
	PointOfView                 PointOfView `json:"PointOfView"`
	GameBuild                   int         `json:"GameBuild"`
	EvtcVersion                 string      `json:"EvtcVersion"`
	GameLanguage                int         `json:"GameLanguage"`
	MainTargetName              string      `json:"MainTargetName"`
	HealthPercentage            any         `json:"HealthPercentage"`
	EncounterStartTime          time.Time   `json:"EncounterStartTime"`
	EncounterDuration           string      `json:"EncounterDuration"`
	DpsReportEIUpload           DpsReport   `json:"DpsReportEIUpload"`
	ParsingStatus               int         `json:"ParsingStatus"`
	ParseTime                   time.Time   `json:"ParseTime"`
	ParseMilliseconds           int         `json:"ParseMilliseconds"`
	ParsingException            any         `json:"ParsingException"`
	ParsingVersion              string      `json:"ParsingVersion"`
	IsEncounterStartTimePrecise bool        `json:"IsEncounterStartTimePrecise"`
	Tags                        []any       `json:"Tags"`
	IsFavorite                  bool        `json:"IsFavorite"`
	LogExtras                   LogExtras   `json:"LogExtras"`
	MissingEncounterStart       bool        `json:"MissingEncounterStart"`
}
