package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func games() {

	w := myApp.NewWindow("Games")

	r1, _ := fyne.LoadResourceFromPath("./icon/8.png")
	button1 := widget.NewButtonWithIcon("Chess", r1, func() {
		go chessApp()
	})
	button1.Resize(fyne.NewSize(100, 40))
	button1.Move(fyne.NewPos(30, 70))

	r2, _ := fyne.LoadResourceFromPath("./icon/9.png")
	button2 := widget.NewButtonWithIcon("Weather App", r2, func() {
		go WeatherFunc()
	})
	button2.Resize(fyne.NewSize(150, 40))
	button2.Move(fyne.NewPos(140, 70))

	w.Resize(fyne.NewSize(450, 400))
	w.CenterOnScreen()
	w.SetContent(container.NewMax(img,
		container.NewWithoutLayout(
			button1, button2,
		),
	))
	w.Show()
}

// paint app code
// func ColorPicker(env gui.Env, pick chan<- color.Color, r image.Rectangle, clr color.Color) {
// 	env.Draw() <- func(drw draw.Image) image.Rectangle {
// 		draw.Draw(drw, r, &image.Uniform{clr}, r.Min, draw.Src)
// 		return r
// 	}

// 	for event := range env.Events() {
// 		switch event := event.(type) {
// 		case win.MoDown:
// 			if event.Point.In(r) {
// 				pick <- clr
// 			}
// 		}
// 	}

// 	close(env.Draw())
// }

// func Canvas(env gui.Env, pick <-chan color.Color, r image.Rectangle) {
// 	canvas := image.NewRGBA(r)
// 	draw.Draw(canvas, r, image.White, r.Min, draw.Src)
// 	dc := gg.NewContextForRGBA(canvas)

// 	env.Draw() <- func(drw draw.Image) image.Rectangle {
// 		draw.Draw(drw, r, canvas, image.ZP, draw.Src)
// 		return r
// 	}

// 	var (
// 		clr     = color.Color(color.Black)
// 		pressed = false
// 		prev    image.Point
// 	)

// 	for {
// 		select {
// 		case clr = <-pick:

// 		case event, ok := <-env.Events():
// 			if !ok {
// 				close(env.Draw())
// 				return
// 			}

// 			switch event := event.(type) {
// 			case win.MoDown:
// 				if event.Point.In(r) {
// 					pressed = true
// 					prev = event.Point
// 				}

// 			case win.MoUp:
// 				pressed = false

// 			case win.MoMove:
// 				if pressed {
// 					x0, y0, x1, y1 := prev.X, prev.Y, event.X, event.Y
// 					prev = event.Point

// 					env.Draw() <- func(drw draw.Image) image.Rectangle {
// 						dc.SetColor(clr)
// 						dc.SetLineCapRound()
// 						dc.SetLineWidth(5)
// 						dc.DrawLine(float64(x0), float64(y0), float64(x1), float64(y1))
// 						dc.Stroke()
// 						rect := image.Rect(x0, y0, x1, y1)
// 						rect.Min.X -= 5
// 						rect.Min.Y -= 5
// 						rect.Max.X += 5
// 						rect.Max.Y += 5
// 						draw.Draw(drw, rect, canvas, rect.Min, draw.Src)
// 						return rect
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func run() {
// 	w, err := win.New(win.Title("Paint"), win.Size(800, 600))
// 	if err != nil {
// 		panic(err)
// 	}

// 	mux, env := gui.NewMux(w)

// 	pick := make(chan color.Color)

// 	for i, clr := range []color.Color{
// 		color.RGBA{255, 0, 0, 255},
// 		color.RGBA{255, 255, 0, 255},
// 		color.RGBA{0, 255, 0, 255},
// 		color.RGBA{0, 255, 255, 255},
// 		color.RGBA{0, 0, 255, 255},
// 		color.RGBA{255, 0, 255, 255},
// 		color.RGBA{255, 255, 255, 255},
// 		color.RGBA{0, 0, 0, 255},
// 	} {
// 		go ColorPicker(mux.MakeEnv(), pick, image.Rect(750, i*75, 800, (i+1)*75), clr)
// 	}

// 	go Canvas(mux.MakeEnv(), pick, image.Rect(0, 0, 750, 600))

// 	for event := range env.Events() {
// 		switch event.(type) {
// 		case win.WiClose:
// 			close(env.Draw())
// 		}
// 	}
// }

// func main() {
// 	mainthread.Run(run)
// }
