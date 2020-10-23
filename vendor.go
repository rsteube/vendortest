package vendortest

type A interface {
        X(c *C) string
}

type B struct {
}

type C struct {

}

func (b *B) X() string {
        return "vendortest"
}

func New() *B {
    return &B{}
}
