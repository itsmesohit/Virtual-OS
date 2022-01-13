package main

import (
	"log"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"github.com/notnil/chess"
)

func chessApp() {

	w := myApp.NewWindow("Chess")

	game := chess.NewGame()
	u := newUI(w, game)
	u.eng = loadOpponent()
	if u.eng != nil {
		defer u.eng.Close()
	} else {
		log.Println("Cound not find stockfish executable, using random player")
		rand.Seed(time.Now().Unix()) // random seed for random responses
	}

	loadGameFromPreference(game, myApp.Preferences())
	myApp.Preferences().AddChangeListener(func() {
		loadGameFromPreference(game, myApp.Preferences())
		u.refreshGrid()
	})

	w.SetContent(u.makeUI())
	w.Resize(fyne.NewSize(480, 480+theme.IconInlineSize()*2+theme.Padding()))

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Chess",
			fyne.NewMenuItem("New Game", func() {
				u.game = chess.NewGame()
				u.refreshGrid()

				_ = u.blackTurn.Set(false)
				_ = u.outcome.Set(string(chess.NoOutcome))
				myApp.Preferences().SetString(preferenceKeyCurrent, game.FEN())
			})),
	))

	w.Show()
}

func move(m *chess.Move, game *chess.Game, white bool, u *ui) {
	off := squareToOffset(m.S1())
	cell := u.grid.objects[off].(*fyne.Container)
	img := cell.Objects[2].(*piece)

	u.over.Resource = resourceForPiece(game.Position().Board().Piece(m.S1()))
	u.over.Resize(img.Size())
	u.over.Refresh() // clear our old resource before showing

	u.over.Show()
	img.Resource = nil
	img.Refresh()

	off = squareToOffset(m.S2())
	cell = u.grid.objects[off].(*fyne.Container)
	pos2 := cell.Position()

	a := canvas.NewPositionAnimation(u.over.Position(), pos2, time.Millisecond*500, func(p fyne.Position) {
		u.over.Move(p)
		u.over.Refresh()
	})
	a.Start()
	time.Sleep(time.Millisecond * 550)

	game.Move(m)
	u.refreshGrid()
	u.over.Hide()

	_ = u.blackTurn.Set(white)
	fyne.CurrentApp().Preferences().SetString(preferenceKeyCurrent, game.FEN())

	_ = u.outcome.Set(string(game.Outcome()))
	if game.Outcome() != chess.NoOutcome {
		result := "draw"
		switch game.Outcome().String() {
		case "1-0":
			result = "won"
		case "0-1":
			result = "lost"
		}

		fyne.CurrentApp().Preferences().SetString(preferenceKeyCurrent, "")
		dialog.ShowInformation("Game ended",
			"Game "+result+" because "+game.Method().String(), u.win)
	}
}
