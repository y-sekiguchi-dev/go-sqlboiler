package model

import (
	"errors"
	"time"
)

type Birthday struct {
	birthday time.Time
}

func NewBirthday(year int, month time.Month, day int) (Birthday, error) {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	if t.After(today()) {
		return Birthday{nil}, errors.New("未来日付です。")
	}
	return Birthday{t}, nil
}

func NewBirthdayFromTime(time time.Time) (Birthday, error) {
	return NewBirthday(time.Year(), time.Month(), time.Day())
}

func (b Birthday) Age() uint {
	today := today()
	age := today.Year() - today.Year()
	return uint(age)
}

func (b Birthday) AsTime() time.Time {
	return b.birthday
}
func today() time.Time {
	now := time.Now().Local()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return today
}