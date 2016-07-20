package searching

type Prefs struct {
	Width      int
	Echo       string
	Prompt     string
	FileFilter []string
}

var DefaultPrefs = Prefs{
	Width:      30,
	Echo:       "Searching for",
	Prompt:     "%> ",
	FileFilter: []string{"*"},
}
