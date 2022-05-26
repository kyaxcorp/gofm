package fs

// filesystem
type FS struct {
	// Physical path of the file on the disk
	Path string `gorm:"size:1000"`
}

func (f *FS) Save() {

}
