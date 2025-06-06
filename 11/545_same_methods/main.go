package main

import (
	"545_same_methods/gadget"
	"fmt"
)

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
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	// mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	// var player Player = gadget.TapePlayer{}
	// playList(player, mixtape)
	// TryOut(player)
	// player = gadget.TapeRecorder{}
	// playList(player, mixtape)
	// TryOut(player)
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
