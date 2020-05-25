package shared

type Deletable interface {
	Deleted() bool
	Delete()
}

type DeletableImpl struct {
	deleted bool
}

func NewDeletableImpl() *DeletableImpl {
	return StoredDeletableImpl(false)
}

func StoredDeletableImpl(deleted bool) *DeletableImpl {
	return &DeletableImpl{deleted}
}

func (d *DeletableImpl) Deleted() bool {
	return d.deleted
}

func (d *DeletableImpl) Delete() {
	d.deleted = true
}
