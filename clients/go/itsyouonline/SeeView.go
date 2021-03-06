package itsyouonline

import (
	"gopkg.in/validator.v2"
)

// See view object
type SeeView struct {
	Category                   string `json:"category" validate:"nonzero"`
	Content_type               string `json:"content_type" validate:"nonzero"`
	Creation_date              string `json:"creation_date" validate:"nonzero"`
	End_date                   string `json:"end_date" validate:"nonzero"`
	Globalid                   string `json:"globalid" validate:"nonzero"`
	Keystore_label             string `json:"keystore_label" validate:"nonzero"`
	Link                       string `json:"link" validate:"nonzero"`
	Markdown_full_description  string `json:"markdown_full_description" validate:"nonzero"`
	Markdown_short_description string `json:"markdown_short_description" validate:"nonzero"`
	Signature                  string `json:"signature" validate:"nonzero"`
	Start_date                 string `json:"start_date" validate:"nonzero"`
	Uniqueid                   string `json:"uniqueid" validate:"nonzero"`
	Username                   string `json:"username" validate:"nonzero"`
	Version                    int    `json:"version" validate:"nonzero"`
}

func (s SeeView) Validate() error {

	return validator.Validate(s)
}
