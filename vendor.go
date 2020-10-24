package vendortest

type A interface {
        X(c *C) string
}

type B struct {
}

type C struct {

}

var CVAR = &C{}

func (b *B) X(c *C) string {
        return "vendortest"
}

func New() *B {
    (&B{}).X(CVAR)
    return &B{}
}
