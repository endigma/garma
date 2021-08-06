package arma

type Mod struct {
	Name       string
	Enabled    bool
	WorkshopID string
}

type Preset struct {
	Name         string
	Filename     string
	CollectionID string
	Mods         []Mod
}
