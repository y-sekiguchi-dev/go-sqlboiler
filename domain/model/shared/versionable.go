package shared

type Versionable interface {
	GetVersion() uint
	SetVersion(version uint)
	IncrementVersion()
}

type VersionableImpl struct{
	version uint
}

func NewVersionableImpl() *VersionableImpl {
	return StoredVersionableImpl(0)
}

func StoredVersionableImpl(version uint) *VersionableImpl {
	return &VersionableImpl{version}
}

func (v VersionableImpl) GetVersion() uint {
	return v.version
}

func (v VersionableImpl) SetVersion(version uint) {
	v.version = version
}

func (v VersionableImpl) IncrementVersion() {
	v.version++
}