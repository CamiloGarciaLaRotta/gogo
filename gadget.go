package gogo

// GadgetID represents a gadget identifier
type GadgetID string

// GadgetFunc represents a
type GadgetFunc func() error

// Gadget represents a gadget stored in the server
type Gadget struct {
	ID   GadgetID   `json:"id"`   // ID is the identifier of the Gadget
	Name string     `json:"name"` // Name is the name of the Gadget
	Do   GadgetFunc `json:"-"`    // Do is the function the gadget performs
}

// GadgetService represents a service to manage Gadget management operations
type GadgetService interface {
	CreateGadget(name string, fn GadgetFunc) error
	Gadget(id GadgetID) (*Gadget, error)
	UpdateGadgetAction(id GadgetID, fn GadgetFunc) error
	UpdateGadgetName(id GadgetID, name string) error
	DeleteGadget(id GadgetID) error
}
