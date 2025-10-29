package gk

type Element interface {
	Update(delta uint64)
	Render()
	Destroy()
}

func UpdateAll(elements []Element, delta uint64) {
	for _, element := range elements {
		element.Update(delta)
	}
}

func RenderAll(elements []Element) {
	for _, element := range elements {
		element.Render()
	}
}

func DestroyAll(elements []Element) {
	for _, element := range elements {
		element.Destroy()
	}
}
