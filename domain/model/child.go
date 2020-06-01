package model

type Child interface {
	SubNo() uint
	FullName() FullName
	Age() uint
	Birthday() Birthday
}

type ChildImpl struct {
	subNo uint
	fullName FullName
	birthday Birthday
}

func NewChild(subNo uint, fullName FullName, birthday Birthday) Child {
	return &ChildImpl{
		subNo:    subNo,
		fullName: fullName,
		birthday: birthday,
	}
}

func (c *ChildImpl) SubNo() uint {
	return c.subNo
}

func (c *ChildImpl) FullName() FullName {
	return c.fullName
}

func (c *ChildImpl) Age() uint {
	return c.Birthday().Age()
}

func (c *ChildImpl) Birthday() Birthday {
	return c.birthday
}
