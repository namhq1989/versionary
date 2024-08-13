package apperrors

import "errors"

var Widget = struct {
	InvalidWidgetID             error
	WidgetNotFound              error
	InvalidWidgetVersion        error
	InvalidWidgetHomepageURL    error
	InvalidWidgetReleaseNoteURL error
}{
	InvalidWidgetID:             errors.New("widget_invalid_id"),
	WidgetNotFound:              errors.New("widget_not_found"),
	InvalidWidgetVersion:        errors.New("widget_invalid_version"),
	InvalidWidgetHomepageURL:    errors.New("widget_invalid_homepage_url"),
	InvalidWidgetReleaseNoteURL: errors.New("widget_invalid_release_note_url"),
}
