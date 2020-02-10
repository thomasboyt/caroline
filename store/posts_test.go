package store

import (
	"testing"

	"github.com/thomasboyt/caroline/models"
)

func TestGetAggregatedPublicPosts(t *testing.T) {
	store := CreateTestStore()

	// returns empty list when table empty
	items := store.GetAggregatedPublicPosts(nil, nil, 20)

	if len(items) != 0 {
		t.Error("did not return empty list when table empty")
	}

	// let's add some items i guess
	user := models.User{}
	err := store.db.Get(&user, "insert into users (name, email, show_in_public_feed) values ('jeff', 'jeff@jambuds.club', true) returning *")
	if err != nil {
		t.Fatal(err)
	}

	song := models.Song{}
	err = store.db.Get(&song, `insert into songs (title, artists) values ('video games', '{"lana del rey"}') returning *`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = store.db.NamedExec("insert into posts (user_id, song_id) values (:userId, :songId)", map[string]interface{}{
		"userId": user.Id,
		"songId": song.Id,
	})
	if err != nil {
		t.Fatal(err)
	}

	// returns public posts
	items = store.GetAggregatedPublicPosts(nil, nil, 20)
	if len(items) != 1 {
		t.Errorf("returned incorrect number of posts %v", len(items))
	}
	if items[0].SongId.Int32 != song.Id {
		t.Errorf("returned incorrect post with song id %v", items[0].SongId.Int32)
	}
}
