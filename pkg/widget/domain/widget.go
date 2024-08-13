package domain

import (
	"time"

	"github.com/namhq1989/versionary-server/internal/database"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/internal/utils/manipulation"
)

type Widget struct {
	ID            string
	Name          string
	Code          string
	HomepageURL   string
	Stats         WidgetStats
	LatestVersion string
	ReleasedAt    time.Time
}

func NewWidget(name, code, homepageURL string) (*Widget, error) {
	if name == "" {
		return nil, apperrors.Common.InvalidName
	}

	if code == "" {
		return nil, apperrors.Common.InvalidCode
	}

	if homepageURL == "" {
		return nil, apperrors.Widget.InvalidWidgetHomepageURL
	}

	return &Widget{
		ID:          database.NewStringID(),
		Name:        name,
		Code:        code,
		HomepageURL: homepageURL,
		Stats: WidgetStats{
			NumOfFollowing: 0,
		},
		LatestVersion: "",
		ReleasedAt:    manipulation.NowUTC(),
	}, nil
}

func (w *Widget) IncreaseNumOfFollowing() error {
	w.Stats.NumOfFollowing++
	return nil
}

func (w *Widget) DecreaseNumOfFollowing() error {
	w.Stats.NumOfFollowing--
	if w.Stats.NumOfFollowing < 0 {
		w.Stats.NumOfFollowing = 0
	}
	return nil
}

func (w *Widget) SetLatestVersion(version string) error {
	if version == "" {
		return apperrors.Widget.InvalidWidgetVersion
	}

	w.LatestVersion = version
	return nil
}
