package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/dialog"
)

func openDialog() {
	log.Println("openDilog called")
	d := dialog.Get()
	files := d.ShowOpenDialog(dialog.OptionOpen{
		Title:       "Renderer Opener",
		DefaultPath: "~/tmp",
		Properties:  []string{dialog.PropOpenFile, dialog.PropShowHiddenFiles, dialog.PropMultiSelections},
		Filters: []dialog.FileFilter{
			dialog.FileFilter{
				Name:       "Go Files",
				Extensions: []string{"go", "js"},
			},
		},
	})
	log.Println("selected files:", files)
}

func main() {
	electron.UseRemote()
	js.Global.Set("openDialog", func() {
		openDialog()
	})
}