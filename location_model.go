package filemanager

type Location struct {
	Name string

	// the driver is the one that creates the bridge between the client and the destination server (location)
	Driver FileInterface
}
