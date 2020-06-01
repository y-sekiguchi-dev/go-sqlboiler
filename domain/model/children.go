package model

type Children struct {
	wrapped []Child
}

func NewChildren() Children {
	return Children{make([]Child, 0, 100)}
}

func (c *Children) add(child Child) bool {
	if c.contains(child) {
		return false
	}
	c.wrapped = append(c.wrapped, child)
	return true
}

func (c *Children) findById(subNo uint) Child {
	if index := c.indexOf(subNo); index < 0 {
		return nil
	} else {
		return c.wrapped[index]
	}
}

func (c *Children) indexOf(subNo uint) int {
	for i, child := range c.wrapped {
		if child.SubNo() == subNo {
			return i
		}
	}
	return -1
}

func (c *Children) contains(child Child) bool {
	return c.indexOf(child.SubNo()) > -1
}

func (c *Children) count() int {
	return len(c.wrapped)
}

func (c *Children) remove(subNo uint) Child {
	i := c.indexOf(subNo)
	if i < 0 {
		return nil
	}
	result := c.wrapped[i]
	copy(c.wrapped[i:], c.wrapped[i+1:])
	c.wrapped[len(c.wrapped)-1] = nil
	return result
}