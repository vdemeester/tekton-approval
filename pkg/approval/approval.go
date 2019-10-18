package approval

import "sync"

var (
	list []Approval
	mtx  sync.RWMutex
	once sync.Once

	StatusUnknown  Status = "Unknown"
	StatusDenied   Status = "Denied"
	StatusApproved Status = "Approved"
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Approval{
		{ID: "i", Name: "Foo", URL: "https://foo.foo", Status: StatusUnknown},
		{ID: "j", Name: "Bar", URL: "https://bar.bar", Status: StatusDenied},
		{ID: "k", Name: "Baz", URL: "https://baz.baz", Status: StatusApproved},
	}
}

type Approval struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status Status `json:"status"`
}

type Status string

func Get() []Approval {
	mtx.RLock()
	defer mtx.RUnlock()
	return list
}

func Add(name, url string, status Status) string {
	return ""
}
