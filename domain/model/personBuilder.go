package model

import (
	"go-sqlboiler/domain/model/shared"
)

type PersonBuilder struct {
	entity *shared.EntityImpl
	version *shared.VersionableImpl
	deleted *shared.DeletableImpl
	birthday Birthday
	personality *Personality
	fullName FullName
	children Children
	hasPartner bool
}

func (pb *PersonBuilder) Personality(personality *Personality) *PersonBuilder {
	pb.personality = personality
	return pb
}

func (pb *PersonBuilder) AddChild(child Child) *PersonBuilder {
	pb.children.add(child)
	return pb
}

func AsNew(birthday Birthday, fullName FullName) *PersonBuilder {
	return &PersonBuilder{
		shared.NewEntityImpl(),
		shared.NewVersionableImpl(),
		shared.NewDeletableImpl(),
		birthday,
		nil,
		fullName,
		NewChildren(),
		false,
	}
}

func AsStored(id PersonId, version uint, deleted bool, birthday Birthday, fullName FullName) *PersonBuilder {
	return &PersonBuilder{
		shared.IdenticalEntityImpl(id),
		shared.StoredVersionableImpl(version),
		shared.StoredDeletableImpl(deleted),
		birthday,
		nil,
		fullName,
		NewChildren(),
		false,
	}
}

func (pb *PersonBuilder) Build() Person {
	return & personImpl{
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