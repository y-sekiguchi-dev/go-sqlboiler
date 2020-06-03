package person

import (
	"go-sqlboiler/domain/model/shared/core"
)

type Builder struct {
	entity      core.Entity
	version     core.Versionable
	deleted     core.Deletable
	birthday    Birthday
	personality *Personality
	fullName    FullName
	children    Children
	hasPartner  bool
}

func (pb *Builder) Personality(personality *Personality) *Builder {
	pb.personality = personality
	return pb
}

func (pb *Builder) AddChild(child Child) *Builder {
	pb.children.add(child)
	return pb
}

func AsNew(birthday Birthday, fullName FullName) *Builder {
	return &Builder{
		core.NewEntity(),
		core.NewVersionable(),
		core.NewDeletable(),
		birthday,
		nil,
		fullName,
		NewChildren(),
		false,
	}
}

func AsStored(id Id, version uint, deleted bool, birthday Birthday, fullName FullName) *Builder {
	return &Builder{
		core.IdenticalEntity(id),
		core.StoredVersionable(version),
		core.StoredDeletable(deleted),
		birthday,
		nil,
		fullName,
		NewChildren(),
		false,
	}
}

func (pb *Builder) Build() Person {
	return &person{
		pb.entity,
		pb.version,
		pb.deleted,
		pb.birthday,
		pb.personality,
		pb.fullName,
		pb.children,
		pb.hasPartner,
	}
}
