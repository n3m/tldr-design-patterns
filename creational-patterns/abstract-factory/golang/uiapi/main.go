package uiapi

/* Components */

// We're defining an abstract interface that will be used to define the components
// that can be rendered.
type component interface {
	Render()
}

// We'll have 4 types of compónents, two per OS.
// We're defining them separete because we want separete logic for each of them
type oSXButton struct{}
type oSXTextBox struct{}
type winButton struct{}
type winTextBox struct{}

func (b *oSXButton) Render()  {}
func (t *oSXTextBox) Render() {}
func (b *winButton) Render()  {}
func (t *winTextBox) Render() {}

/* Factory */

// The factories will be used to create components and will be separated by the OS
type guiFactory interface {
	CreateButton() component
	CreateTextBox() component
}

// We've two types of OS: OSX and Windows.
type OSXFactory struct{}
type WindowsFactory struct{}

// Here we're using the abstract factory to create the components depending on the OS
func (f *OSXFactory) CreateButton() component {
	return &oSXButton{}
}

func (f *OSXFactory) CreateTextBox() component {
	return &oSXTextBox{}
}

func (f *WindowsFactory) CreateButton() component {
	return &winButton{}
}

func (f *WindowsFactory) CreateTextBox() component {
	return &winTextBox{}
}

/* Final Creator */

// Here we have the final function that we'll need +
func CreateOSFactory(osType string) guiFactory {
	switch osType {
	case "OSX":
		return &OSXFactory{}
	case "Windows":
		return &WindowsFactory{}
	default:
		return nil
	}

}
