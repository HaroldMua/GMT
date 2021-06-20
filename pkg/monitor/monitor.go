package monitor

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"

	apierror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cache"

	v1 "github.com/HaroldMua/GMT/api/v1"
	"github.com/HaroldMua/GMT/pkg/log"
)

type Monitor struct {
	// CRD Info
	cache cache.Cache
	client client.Client

	nodeName string

	// GPU Info
	cardList v1.CardList
	cardNumber uint
	freeMemorySum  uint64
	totalMemorySum uint64

	updateInterval int64
}

func NewMonitor(interval int64, client client.Client, cache cache.Cache) *Monitor {
	return &Monitor{
		cardList:       make(v1.CardList, 0),
		cardNumber:     0,
		updateInterval: interval,
		client:         client,
		cache:          cache,
	}
}

func Run(m *Monitor) {
	// Initialize CRD and set config
	m.nodeName = os.Getenv("NODE_NAME")
	if err := m.createGmt(); err != nil {
		panic(err)
	}
	m.process()
}

func (m *Monitor) createGmt() error {
	scv := v1.Gmt{
		ObjectMeta: metav1.ObjectMeta{
			Name: m.nodeName,
		},
		Spec: v1.GetSpec{
			UpdateInterval: m.updateInterval,
		},
	}
	err := m.client.Create(context.Background(), &scv)
	if err != nil && !apierror.IsAlreadyExists(err) {
		return err
	}
	return nil
}

func (m *Monitor) process() {
	interval := time.Duration(m.updateInterval) * time.Millisecond
	ticker := time.NewTicker(interval)
	for {
		<- ticker.C

		// update the info of GPU
		m.updateGPU()

		currentGmt := v1.Gmt{}
		err := m.client.Get(context.Background(), client.ObjectKey{
			Name: m.nodeName,
		}, &currentGmt)
		if err != nil {
			log.ErrPrint(err)
			continue
		}

		// update the status of Gmt, if there are no changes in GPU info, don't need update status
		if m.needUpdate(currentGmt.Status) {
			updateGmt := currentGmt.DeepCopy()
			updateGmt.Status = v1.GmtStatus{
				CardList:       m.cardList,
				CardNumber:     m.cardNumber,
				TotalMemorySum: m.totalMemorySum,
				FreeMemorySum:  m.freeMemorySum,
				UpdateTime:     &metav1.Time{
					Time: time.Now(),
				},
			}
			if err := m.client.Update(context.Background(), updateGmt); err != nil {
				log.ErrPrint(err)
			}
		}

	}
}

func (m *Monitor) updateGPU() {

}

func (m *Monitor) needUpdate(status v1.GmtStatus) bool {
	if status.UpdateTime == nil {
		log.Print("CardList is Null, needs update.")
		return true
	}

	if status.TotalMemorySum != m.totalMemorySum {
		log.Print("Total memory changed, needs update.")
		return true
	}

	if status.FreeMemorySum != m.freeMemorySum {
		log.Print("Free memory changed, needs update.")
		return true
	}
	if status.CardNumber != m.cardNumber {
		log.Print("Card Number changed, needs update.")
		return true
	}
	if !reflect.DeepEqual(status.CardList, m.cardList) {
		log.Print("Card List changed, needs update.")
		return true
	}
	return false
}