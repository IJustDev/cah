package src

import uuid "github.com/nu7hatch/gouuid"

func NewObjectId() string {
	u, _ := uuid.NewV4()

	return u.String()
}
