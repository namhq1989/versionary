package domain

import (
	"time"

	"github.com/namhq1989/versionary-server/internal/database"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/internal/utils/manipulation"
)

type WidgetVersion struct {
	ID             string
	WidgetID       string
	Version        string
	ReleaseNoteURL string
	ReleasedAt     time.Time
}

func NewWidgetVersion(widgetID, version, releaseNoteURL string) (*WidgetVersion, error) {
	if !database.IsValidObjectID(widgetID) {
		return nil, apperrors.Widget.InvalidWidgetID
	}

	if version == "" {
		return nil, apperrors.Widget.InvalidWidgetVersion
	}

	if releaseNoteURL == "" {
		return nil, apperrors.Widget.InvalidWidgetReleaseNoteURL
	}

	return &WidgetVersion{
		ID:             database.NewStringID(),
		WidgetID:       widgetID,
		Version:        version,
		ReleaseNoteURL: releaseNoteURL,
		ReleasedAt:     manipulation.NowUTC(),
	}, nil
}
