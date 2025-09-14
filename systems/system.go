package systems

type BaseSystem interface {
	Init()
	Render()
	Update()
	Done()
}
