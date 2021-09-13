package adapter

import "testing"

func TestAdapter(t *testing.T) {
	player := MediaAdapter{}
	player.Player("mp3")
	player.Player("vlc")
	player.Player("mp4")
	player.Player("what")
}
