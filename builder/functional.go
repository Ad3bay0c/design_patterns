package main

type Person1 struct {
	name, position string
}

type personMod func(*Person1)

type PersonBuilder1 struct {
	actions []personMod
}

func (b *PersonBuilder1) Called(name string) *PersonBuilder1 {
	b.actions = append(b.actions, func(p *Person1) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder1) Build() *Person1 {
	p := &Person1{}
	for _, action := range b.actions {
		action(p)
	}
	return p
}