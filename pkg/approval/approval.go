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
		{ID: "bml1l5i17s989p4h6vbg", Name: "Foo", URL: "http://greeter1.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
		{ID: "bml1l5ssss989p4h6vbg", Name: "Bar", URL: "http://greeter2.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusDenied},
		{ID: "bml1l5i17s98hhhh6vbg", Name: "Baz", URL: "http://greeter3.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusApproved},
		{ID: "bmlhce217s9fvu4h81eg", Name: "Baz", URL: "http://greeter4.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusApproved},
		{ID: "bmlhcea17s9fvu4h81f0", Name: "Greeter", URL: "http://greeter5.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
		{ID: "bmlhcpq17s9fvu4h81mg", Name: "Balade", URL: "http://greeter6.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
		{ID: "bmlhcn217s9fvu4h81lg", Name: "Nature", URL: "http://greeter7.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
		{ID: "bmlhcra17s9fvu4h81ng", Name: "Fire", URL: "http://greeter8.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
		{ID: "bmlhcv217s9fvu4h81rg", Name: "Water", URL: "http://greeter9.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
		{ID: "bmlhd1a17s9fvu4h81ug", Name: "Dance", URL: "http://greeter10.dev.apps.asnieres.devcluster.openshift.com/", Status: StatusUnknown},
	}
}

type Approval struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status Status `json:"status"`
}

type Status string

func List() []Approval {
	mtx.RLock()
	defer mtx.RUnlock()
	return list
}

func Get(id string) *Approval {
	mtx.RLock()
	for _, a := range list {
		if a.ID == id {
			return &a
		}
	}
	mtx.RUnlock()
	return nil
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
