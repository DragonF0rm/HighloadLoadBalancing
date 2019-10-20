package router

import (
	"LoadBalancingBackend/cfg"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	responseStatus  = http.StatusOK
	pendingRequests = 0
	fakeConnections = 0
	mutex sync.Mutex
)

func LivenessProbeHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(responseStatus)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	maxConnections := cfg.GetInt("max_connections")
	loadCoefficient := float32(maxConnections - pendingRequests) / float32(maxConnections)

	reqMaxTime := cfg.GetInt("req_max_time")
	reqMinTime := cfg.GetInt("req_min_time")
	sleepMillisCount := int(float32(rand.Intn(reqMaxTime - reqMinTime) + reqMinTime) * loadCoefficient)
	sleepDuration := time.Duration(sleepMillisCount) * time.Millisecond

	time.Sleep(sleepDuration)
	w.WriteHeader(responseStatus)
}

func OverloadHandler(w http.ResponseWriter, r *http.Request) {
	fakeConnBatch := cfg.GetInt("fake_conn_batch")
	fakeConnCount := fakeConnBatch
	fakeConnArg := r.URL.Query().Get("conn")
	if fakeConnArg == "" {
		var err error
		fakeConnCount, err = strconv.Atoi(fakeConnArg)
		if err != nil {
			fakeConnCount = fakeConnBatch
		}
	}
	mutex.Lock()
	fakeConnections += fakeConnCount
	pendingRequests += fakeConnCount
	mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	responseStatus = http.StatusInternalServerError
	w.WriteHeader(http.StatusOK)
}

func RebootHandler(w http.ResponseWriter, r *http.Request) {
	responseStatus = http.StatusOK
	mutex.Lock()
	pendingRequests -= fakeConnections
	fakeConnections = 0
	mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}