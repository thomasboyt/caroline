package resources

type ColorSchemeJson struct {
	BackgroundGradientName string
	TextColor              string
}

type UserProfileJson struct {
	Id   int32
	Name string
	// ColorScheme *ColorSchemeJson
}
