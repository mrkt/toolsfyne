package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
)

func main() {
	a := app.New()
	w := a.NewWindow("CSV format TEXT tool v1.0")
	//s := ""
	//data1 := binding.BindString(&s)
	//data2 := binding.BindString(&s)
	//entry1 := widget.NewEntryWithData(data1)
	//entry2 := widget.NewEntryWithData(data2)
	title := canvas.NewText("CSV format TEXT tool v1.0", color.White)
	line1 := canvas.NewLine(color.White)
	line2 := canvas.NewLine(color.White)
	line3 := canvas.NewLine(color.White)
	line4 := canvas.NewLine(color.White)
	text1 := canvas.NewText("Entry csv:", color.White)
	text2 := canvas.NewText("Entry format text:", color.White)
	text3 := canvas.NewText("Text:", color.White)
	title.TextStyle = fyne.TextStyle{Italic: true}
	lineEntry1 := widget.NewMultiLineEntry()
	lineEntry1.SetMinRowsVisible(8)
	lineEntry1.Text =
		`Lumbergh, Bill, Initech
Waddams, Milton, Initech
Gibbons, Peter, Initech
Bolton, Michael, Initech`
	lineEntry2 := widget.NewMultiLineEntry()
	lineEntry2.SetMinRowsVisible(5)
	lineEntry2.Text =
		`To: $1.$0@$2.com
Hello $1 $0,
I'm sorry to inform you of a terrible accident at $2.
---`
	label := widget.NewMultiLineEntry()
	label.SetMinRowsVisible(8)
	author := canvas.NewText("https://github.com/mrkt", color.White)
	author.Alignment = fyne.TextAlignTrailing
	author.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(container.NewVBox(
		title,
		line1,
		text1,
		lineEntry1,
		line2,
		text2,
		lineEntry2,
		line3,
		widget.NewButton("Run", func() {
			countSplit := make([][]string, 0)
			lineEntryArr := strings.Split(lineEntry1.Text, "\n")
			for _, v := range lineEntryArr {
				countSplit = append(countSplit, strings.Split(v, ","))
			}
			var runText string
			for k, v := range countSplit {
				formatText := lineEntry2.Text
				for kk, vv := range v {
					formatText = strings.Replace(formatText, fmt.Sprintf("$%d", kk), vv, -1)
				}
				runText += formatText
				if k < len(countSplit)-1 {
					runText += "\n"
				}
			}
			label.SetText(runText)

		}),
		line4,
		text3,
		label,
		author,
	))
	w.Resize(fyne.NewSize(800, 800))
	w.Show()
	a.Run()
}
