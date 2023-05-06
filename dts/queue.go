package dts

// Simple Queue implementation
type Queue interface {
	Enqueue(string)
	DequeueAll() []string
	IsNotEmpty() bool
	Size() int
}

type DefaultQueue struct {
	queue []string
}

func (d *DefaultQueue) Enqueue(s string) {
	d.queue = append(d.queue, s)
}

func (d *DefaultQueue) DequeueAll() []string {
	q := d.queue
	d.queue = []string{}

	return q
}

func (d *DefaultQueue) IsNotEmpty() bool {
	return d.Size() > 0
}

func (d *DefaultQueue) Size() int {
	return len(d.queue)
}

func NewQueue() Queue {
	return &DefaultQueue{}
}
