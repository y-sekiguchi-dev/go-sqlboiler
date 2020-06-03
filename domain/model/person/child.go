package person

type Child interface {
	SubNo() uint
	FullName() FullName
	Age() uint
	Birthday() Birthday
}

type child struct {
	subNo    uint
	fullName FullName
	birthday Birthday
}

func StoredChild(subNo uint, fullName FullName, birthday Birthday) Child {
	return &child{
		subNo:    subNo,
		fullName: fullName,
		birthday: birthday,
	}
}

func (c *child) SubNo() uint {
	return c.subNo
}

func (c *child) FullName() FullName {
	return c.fullName
}

func (c *child) Age() uint {
	return c.Birthday().Age()
}

func (c *child) Birthday() Birthday {
	return c.birthday
}
