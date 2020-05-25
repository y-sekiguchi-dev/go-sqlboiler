package model

type Child struct {
	subNo uint
	fullName FullName
	birthday Birthday
}

func newChild(subNo uint, fullName FullName, birthday Birthday) Child {
	return Child{
		subNo:    subNo,
		fullName: fullName,
		birthday: birthday,
	}
}

func (c *Child) SubNo() uint {
	return c.subNo
}

func (c *Child) FullName() FullName {
	return c.fullName
}

func (c *Child) Age() uint {
	return c.birthday.Age()
}
