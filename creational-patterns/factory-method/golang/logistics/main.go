package logistics

/* Interfaces */
type Product interface {
	Use()
}

type ProductCreator interface {
	CreateProduct() Product
}

/* Products */
type Toolbox struct{}

func (t *Toolbox) Use() {}

type Bottle struct{}

func (t *Bottle) Use() {}

/* Creators */

type ToolboxCreator struct {
	info string
}

func (t *ToolboxCreator) CreateProduct() *Toolbox {
	return &Toolbox{}
}

type BottleCreator struct {
	info string
}

func (t *BottleCreator) CreateProduct() *Bottle {
	return &Bottle{}
}

/* Creator Functions*/
