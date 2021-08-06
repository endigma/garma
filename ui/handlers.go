package ui

import (
	"github.com/endigma/garma/common"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rs/zerolog/log"
)

func presetToggle(t *gtk.CellRendererToggle, iterString string) {
	iter, err := presetlist.TreeModel.GetIterFromString(iterString)
	common.CheckErr(err, "Unable to get Iter")

	presetlist.Set(iter, []int{0}, []interface{}{!t.GetActive()})

	m, _ := presetlist.GetValue(iter, 1)
	presetName, _ := m.GetString()

	log.Info().Str("preset", presetName).Bool("toggle", t.GetActive()).Msg("toggled preset")
}

func modToggle(t *gtk.CellRendererToggle, iterString string) {
	iter, err := modlist.TreeModel.GetIterFromString(iterString)
	common.CheckErr(err, "Unable to get Iter")

	modlist.Set(iter, []int{0}, []interface{}{!t.GetActive()})

	m, _ := modlist.GetValue(iter, 1)
	modName, _ := m.GetString()

	log.Info().Str("mod", modName).Bool("toggle", t.GetActive()).Msg("toggled mod")
}

func toggleMods(b *gtk.CheckButton) {
	log.Info().Bool("toggle", b.GetActive()).Msg("Toggling mods")
}

func start(b *gtk.Button) {
	log.Info().Msg("Mock starting Arma 3")
}

func force_kill(b *gtk.Button) {
	log.Info().Msg("Mock killing Arma 3")
}
