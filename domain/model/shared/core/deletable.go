package core

type Deletable interface {
	Deleted() bool
	Delete()
}

type deletable struct {
	deleted bool
}

func NewDeletable() Deletable {
	return StoredDeletable(false)
}

func StoredDeletable(deleted bool) Deletable {
	return &deletable{deleted}
}

func (d *deletable) Deleted() bool {
	return d.deleted
}

func (d *deletable) Delete() {
	d.deleted = true
}
