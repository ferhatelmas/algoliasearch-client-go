package search

import (
	"fmt"
)

type Batch struct {
	Action BatchAction `json:"action"`
	Body   interface{} `json:"body,omitempty"`
}

type BatchAction string

const (
	AddObject                   BatchAction = "addObject"
	UpdateObject                BatchAction = "updateObject"
	PartialUpdateObject         BatchAction = "partialUpdateObject"
	PartialUpdateObjectNoCreate BatchAction = "partialUpdateObjectNoCreate"
	DeleteObject                BatchAction = "deleteObject"
	Delete                      BatchAction = "delete"
	Clear                       BatchAction = "clear"
)

func newOperationBatch(objects []interface{}, action BatchAction) ([]Batch, error) {
	operations := make([]Batch, len(objects))

	for i, o := range objects {
		// If the action is neither an `addObject` nor a `clear`, the
		// `objectID` field is required.
		if action != AddObject && action != Clear && !hasObjectIDField(o) {
			return nil, fmt.Errorf("missing `objectID` field in object %#v (position=%d)", o, i)
		}
		operations[i].Action = action
		operations[i].Body = o
	}

	return operations, nil
}
