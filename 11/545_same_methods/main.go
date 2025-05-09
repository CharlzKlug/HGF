package main

import "545_same_methods/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	TryOut(player)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	TryOut(player)
}
