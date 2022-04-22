package common

import "time"

type ReleaseList struct {
	ReleaseMeta *ReleaseMeta   `json:"releaseMeta"`
	ReleaseNote []*ReleaseNote `json:"releaseNote"`
}

type ReleaseMeta struct {
	ImageDigest string    `json:"imageDigest"`
	Release     int       `json:"release"`
	Version     int       `json:"version"`
	CreatedOn   time.Time `json:"createdOn"`
}

type ReleaseNote struct {
	Note string `json:"note"`
	Body string `json:"body"`
}
