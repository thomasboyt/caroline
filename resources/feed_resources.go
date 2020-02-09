package resources

import (
	"time"

	"github.com/thomasboyt/jam-buds-golang/models"
)

type SongJson struct {
	models.SongWithMeta
}

type MixtapePreviewJson struct {
	models.MixtapePreview
}

type FeedItemJson struct {
	Timestamp time.Time
	UserNames []string
	Song      *SongJson
	Mixtape   *MixtapePreviewJson
}

type PlaylistItemJson struct {
	Timestamp time.Time
	Song      *SongJson
	Mixtape   *MixtapePreviewJson
}

type PlaylistJson struct {
	Items       []PlaylistItemJson
	UserProfile UserProfileJson
}
