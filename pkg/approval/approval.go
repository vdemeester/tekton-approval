package approval

import (
	"sync"

	"github.com/rs/xid"
)

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
		{ID: "bml1l5i17s989p4h6vbg", Name: "Foo", URL: "https://foo.foo", Status: StatusUnknown},
		{ID: "bml1l5ssss989p4h6vbg", Name: "Bar", URL: "https://bar.bar", Status: StatusDenied},
		{ID: "bml1l5i17s98hhhh6vbg", Name: "Baz", URL: "https://baz.baz", Status: StatusApproved},
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

func Add(name, url string, status Status) Approval {
	a := Approval{
		ID:     xid.New().String(),
		Name:   name,
		URL:    url,
		Status: status,
	}
	mtx.Lock()
	list = append(list, a)
	mtx.Unlock()
	return a
}
func Update(id string, a Approval) Approval {
	for i, aa := range list {
		if aa.ID == id {
			mtx.Lock()
			list[i] = a
			mtx.Unlock()
			return a
		}
	}
	return a
}
