package model

import (
	"go-sqlboiler/domain/model/shared"
	"strconv"
)

type (
	PersonId shared.Id
	Person interface {
		shared.Entity
		shared.Versionable
		shared.Deletable

		Age() uint
		Personality() Personality
		RevisePersonality(personality Personality)
		FullName() FullName
		Rename(fullName FullName)
		HasChild() bool
		AddChild(child Child) bool
		RemoveChild(childSubNo uint) (removed *Child)
		HasPartner() bool
		SetPartner(hasPartner bool)
	}

)

type personIdImpl struct {
	id uint
}

func (pid *personIdImpl) String() string {
	return strconv.Itoa(int(pid.id))
}

func newPersonId(id uint) PersonId {
	return &personIdImpl{id}
}

type personImpl struct {
	*shared.EntityImpl
	*shared.VersionableImpl
	*shared.DeletableImpl
	birthday Birthday
	personality Personality
	fullName FullName
	children Children
	hasPartner bool
}

func (p *personImpl) Age() uint {
	return p.birthday.Age()
}

func (p *personImpl) Personality() Personality {
	return p.personality
}

func (p *personImpl) RevisePersonality(personality Personality) {
	p.personality = personality
}

func (p *personImpl) FullName() FullName {
	return p.fullName
}

func (p *personImpl) Rename(fullName FullName) {
	p.fullName = fullName
}

func (p *personImpl) HasChild() bool {
	return p.children.count() > 0
}

func (p *personImpl) AddChild(child Child) bool {
	return p.children.add(child)
}

func (p *personImpl) RemoveChild(childSubNo uint) (removed *Child) {
	return p.children.remove(childSubNo)
}

func (p *personImpl) HasPartner() bool {
	return p.hasPartner
}

func (p *personImpl) SetPartner(hasPartner bool) {
	p.hasPartner = hasPartner
}
