package ui

import (
	"errors"
	"os"

	"github.com/endigma/garma/arma"
	"github.com/endigma/garma/common"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rs/zerolog/log"
)

var (
	window     *gtk.Window
	builder    *gtk.Builder
	modlist    *gtk.ListStore
	presetlist *gtk.ListStore
)

const (
	appID = "com.gitcat.endigma.gao"
)

func Run() {
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	common.CheckErr(err, "Unable to create Application")

	app.Connect("startup", func() {
		log.Info().Msg("Application started")
	})

	app.Connect("activate", Activate)

	os.Exit(app.Run(os.Args))
}

func Activate(app *gtk.Application) {
	var err error
	builder, err = gtk.BuilderNewFromFile("ui/app.glade")
	common.CheckErr(err, "Unable to activate Application")

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	builder.ConnectSignals(map[string]interface{}{
		"destroy": func() {
			log.Info().Msg("Application quitting")
		},
		"select_preset": presetToggle,
		"mod_toggle":    modToggle,
		"toggle_mods":   toggleMods,
		"start":         start,
		"force_kill":    force_kill,
	})

	// Setup objects
	getObjects()
	initializeModlist()

	// Show window
	window.Show()
	app.AddWindow(window)
}

func initializeModlist() {
	mods := arma.GetMods()

	for _, m := range mods {
		modlist.Set(modlist.Append(), []int{0, 1, 2}, []interface{}{m.Enabled, m.Name, m.WorkshopID})
	}
}

func getObjects() {
	var err error
	window, err = func() (*gtk.Window, error) {
		obj, err := builder.GetObject("main_window")
		common.CheckErr(err, "Cannot find main_window")

		if win, ok := obj.(*gtk.Window); ok {
			return win, nil
		}
		return nil, errors.New("not a *gtk.Window")
	}()
	common.CheckErr(err, "Unable to create Window")

	modlist, err = func() (*gtk.ListStore, error) {
		obj, err := builder.GetObject("mod_list")
		common.CheckErr(err, "Cannot find mod_list")

		if modlist, ok := obj.(*gtk.ListStore); ok {
			return modlist, nil
		}
		return nil, errors.New("not a *gtk.ListStore")
	}()
	common.CheckErr(err, "Unable to create ListStore")

	presetlist, err = func() (*gtk.ListStore, error) {
		obj, err := builder.GetObject("preset_lis")
		common.CheckErr(err, "Cannot find preset_list")

		if modlist, ok := obj.(*gtk.ListStore); ok {
			return modlist, nil
		}
		return nil, errors.New("not a *gtk.ListStore")
	}()
	common.CheckErr(err, "Unable to create ListStore")
}

// ls.Set(ls.Append(), []int{0, 1, 2}, []interface{}{m.Enabled, m.Name, m.WorkshopID})
// ls.Set(ls.Append(), []int{0, 1}, []interface{}{p.Name, p.Filename})
