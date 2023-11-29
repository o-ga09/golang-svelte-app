package counter

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Get() (Clicks, error)
	Put(counter Clicks) error
}

type ICounter interface {
	GetCount() (int, error)
	UpdateCount() error
}

type Clicks struct {
	repo    Repository
	Counter string `dynamo:"counter,omitempty"`
	Count   int    `dynamo:"count,omitempty"`
}

func New(r Repository) *Clicks {
	return &Clicks{repo: r, Counter: "clicks", Count: 0}
}

func (c *Clicks) GetCount() (int, error) {
	res, err := c.repo.Get()
	if err != nil {
		return 0, err
	}
	return res.Count, nil
}

func (c *Clicks) UpdateCount() error {
	newCount := Clicks{Counter: "clicks", Count: c.Count + 1}
	err := c.repo.Put(newCount)
	if err != nil {
		return err
	}

	return nil
}