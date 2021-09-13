package adapter

import "fmt"

type MediaPlayer interface {
	PlayMp3()
}

type AudioPlayer struct {
}

func (ap *AudioPlayer) PlayMp3() {
	fmt.Println("audio mp3 player")
}

type AdvanceMediaPlayer interface {
	PlayVlc()
	PlayMp4()
}

type AdvanceAudioPlayer struct {
}

func (aap *AdvanceAudioPlayer) PlayVlc() {
	fmt.Println("advance audio vlc player")
}

func (aap *AdvanceAudioPlayer) PlayMp4() {
	fmt.Println("advance audio mp4 player")
}

type MediaAdapter struct {
	AudioPlayer AudioPlayer
	AdvanceAudioPlayer AdvanceAudioPlayer
}

func (ma *MediaAdapter) Player(fileType string) {
	switch fileType {
	case "mp3":
		ma.AudioPlayer.PlayMp3()
	case "vlc":
		ma.AdvanceAudioPlayer.PlayVlc()
	case "mp4":
		ma.AdvanceAudioPlayer.PlayMp4()
	default:
		fmt.Println("file type err")
	}
}
