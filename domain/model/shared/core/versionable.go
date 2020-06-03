package core

type Versionable interface {
	GetVersion() uint
	SetVersion(version uint)
	IncrementVersion()
}

type versionable struct {
	version uint
}

func NewVersionable() *versionable {
	return StoredVersionable(0)
}

func StoredVersionable(version uint) *versionable {
	return &versionable{version}
}

func (v versionable) GetVersion() uint {
	return v.version
}

func (v versionable) SetVersion(version uint) {
	v.version = version
}

func (v versionable) IncrementVersion() {
	v.version++
}
