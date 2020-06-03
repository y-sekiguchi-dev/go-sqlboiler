package person

import (
	"errors"
	"go-sqlboiler/domain/model/shared/core"
)

type (
	Id interface {
		core.Id
		AsPersistForm() uint
	}
	Person interface {
		core.Entity
		core.Versionable
		core.Deletable
		GiveId(id Id) error
		Age() uint
		Birthday() Birthday
		Personality() *Personality
		RevisePersonality(personality *Personality)
		FullName() FullName
		Rename(fullName FullName)
		ChildrenView() []Child
		HasChild() bool
		AddChild(child Child) bool
		RemoveChild(childSubNo uint) (removed Child)
		HasPartner() bool
		SetPartner(hasPartner bool)
	}
)

type person struct {
	core.Entity
	core.Versionable
	core.Deletable
	birthday    Birthday
	personality *Personality
	fullName    FullName
	children    Children
	hasPartner  bool
}

func (p *person) GiveId(id Id) error {
	if core.IsEmpty(p.Id()) {
		p.Entity = core.IdenticalEntity(id)
		return nil
	}
	return errors.New("id is unable to change")
}

func (p *person) Age() uint {
	return p.Birthday().Age()
}

func (p *person) Birthday() Birthday {
	return p.birthday
}

func (p *person) Personality() *Personality {
	return p.personality
}

func (p *person) RevisePersonality(personality *Personality) {
	p.personality = personality
}

func (p *person) FullName() FullName {
	return p.fullName
}

func (p *person) Rename(fullName FullName) {
	p.fullName = fullName
}

func (p *person) ChildrenView() []Child {
	return p.children.view()
}

func (p *person) HasChild() bool {
	return p.children.count() > 0
}

func (p *person) AddChild(child Child) bool {
	return p.children.add(child)
}

func (p *person) RemoveChild(childSubNo uint) (removed Child) {
	return p.children.remove(childSubNo)
}

func (p *person) HasPartner() bool {
	return p.hasPartner
}

func (p *person) SetPartner(hasPartner bool) {
	p.hasPartner = hasPartner
}
