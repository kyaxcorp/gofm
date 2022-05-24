package filemanager

type Location struct {
	Name string

	// the driver is the one that creates the bridge between the client and the destination server (location)
	Driver DriverFileInterface

	// TODO: we should add here driver options, the user should make reference to the package driver/DRIVE_NAME
	// 		 and retrieve the Options Model from there!
}
