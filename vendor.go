package vendortest

type A interface {
        X() string
}

type B struct {
}

func (b *B) X() string {
        return "vendortest"
}
