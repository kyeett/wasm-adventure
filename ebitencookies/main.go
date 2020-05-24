package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/colornames"
	"log"
	"strconv"

	"github.com/kyeett/wasm-adventure/preferences"
	"github.com/hajimehoshi/ebiten"
)

var pref *preferences.Preferences

func init() {
	var err error
	pref, err = preferences.New("ebitengame")
	if err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyR):
		pref.SetItem("color", "red")
	case inpututil.IsKeyJustPressed(ebiten.KeyG):
		pref.SetItem("color", "green")
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		pref.SetItem("color", "white")
	}

	for i := ebiten.Key0; i <= ebiten.Key9; i++ {

		if inpututil.IsKeyJustPressed(i) {
			fmt.Println(i)
			pref.SetItem("number", int64(i))
		}
	}


	clr, _ := pref.GetString("color")
	switch clr {
	case "red":
		screen.Fill(colornames.Darkred)
	case "green":
		screen.Fill(colornames.Green)
	default:
		screen.Fill(colornames.White)
	}

	number, err := pref.GetInt("number")
	if err != nil {
		fmt.Println(err)
	}
	ebitenutil.DebugPrint(screen, strconv.FormatInt(number, 10))

	return nil
}

func main() {
	if err := ebiten.Run(update, 640,480,1, "ebitencookie"); err != nil {
		log.Fatal(err)
	}
}
