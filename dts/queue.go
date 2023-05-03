package dts

// Simple Queue implementation
type Queue interface {
	Enqueue(string)
	Dequeue() []string
	NextValue() string
	IsNotEmpty() bool
	Size() int
}

type DefaultQueue struct {
	queue []string
}

func (d *DefaultQueue) Enqueue(s string) {
	d.queue = append(d.queue, s)
}

func (d *DefaultQueue) Dequeue() []string {
	d.queue = d.queue[1:]
	return d.queue
}

func (d *DefaultQueue) IsNotEmpty() bool {
	return d.Size() > 0
}

func (d *DefaultQueue) Size() int {
	return len(d.queue)
}

func (d *DefaultQueue) NextValue() string {
	return d.queue[0]
}

func NewQueue() Queue {
	return &DefaultQueue{}
}
