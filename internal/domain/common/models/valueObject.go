package models

import "reflect"

type ValueObjet[OT any] struct{}

func (v *ValueObjet[OT]) Compare(c OT) bool {

	return reflect.DeepEqual(v, c)
}
