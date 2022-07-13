package main

import "design-patterns/uiapi"

/*
	CONTEXT:
 The Abstract factory is used to generate objects that are related. The relation is defined by a context.

 In this example, we're goint to try and tie up a service that lets you create any kind of UI, depending on the OS that you select.
*/

/*
 SOLUTION:
 Go to file > uiapi/main.go for further explanation
*/

func main() {

	// WindowsGUI will return us a Windows GUI Factory
	WindowsGUI := uiapi.CreateOSFactory("Windows")

	// OSXGUI will return us an OSX GUI Factory
	OSXGUI := uiapi.CreateOSFactory("OSX")

}
