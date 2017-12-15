package banner


type Banner struct {
	origin	string
	adjoin	bool
	ascii	Ascii
	theme	Theme
}

func NewBanner(origin string, adjoin bool) *Banner {
	b := &Banner{
		origin: origin,
		adjoin: adjoin,
		theme: DefaultTheme{},
	}
	ascii := b.theme.Convert(origin[0])
	for i := 1; i < len(origin); i++ {
		ascii.Append(b.theme.Convert(origin[i]), adjoin)
	}
	b.ascii = ascii
	return b
}

func (b *Banner) Show() {
	b.ascii.Show()
}

