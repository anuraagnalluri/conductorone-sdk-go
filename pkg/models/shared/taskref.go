// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskRef -  This object references a task by ID.
type TaskRef struct {
	//  The ID of the TaskRef
	//
	ID *string `json:"id,omitempty"`
}

func (o *TaskRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
