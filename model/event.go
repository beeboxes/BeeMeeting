package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/manyminds/api2go/jsonapi"
)

// Event is a generic database Event
type Event struct {
	ID            int64     `json:"-"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Type          int       `json:"type"`
	SignInStartAt time.Time `json:"sign-in-start-at"`
	SignInEndAt   time.Time `json:"sign-in-end-at"`
	Form          *ApplicationForm
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Event) GetID() string {
	return strconv.FormatInt(u.ID, 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Event) SetID(id string) error {
	u.ID, _ = strconv.ParseInt(id, 10, 64)
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Event) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "chocolates",
			Name: "sweets",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Event) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, chocolate := range u.Chocolates {
		result = append(result, jsonapi.ReferenceID{
			ID:   chocolate.GetID(),
			Type: "chocolates",
			Name: "sweets",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Event) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.Chocolates {
		result = append(result, u.Chocolates[key])
	}

	return result
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *Event) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "sweets" {
		u.ChocolatesIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new sweets that a Events loves so much
func (u *Event) AddToManyIDs(name string, IDs []string) error {
	if name == "sweets" {
		u.ChocolatesIDs = append(u.ChocolatesIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some sweets from a Events because they made him very sick
func (u *Event) DeleteToManyIDs(name string, IDs []string) error {
	if name == "sweets" {
		for _, ID := range IDs {
			for pos, oldID := range u.ChocolatesIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.ChocolatesIDs = append(u.ChocolatesIDs[:pos], u.ChocolatesIDs[pos+1:]...)
				}
			}
		}
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
