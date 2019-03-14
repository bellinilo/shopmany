package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gianarb/shopmany/frontend/config"
)

const unhealthy = "unhealty"
const healthy = "healthy"

type healthResponse struct {
	Status string
	Checks []check
}

type check struct {
	Error  string
	Status string
	Name   string
}

func NewHealthHandler(config config.Config, hclient *http.Client) *healthHandler {
	return &healthHandler{
		config:  config,
		hclient: hclient,
	}
}

type healthHandler struct {
	config  config.Config
	hclient *http.Client
}

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b := healthResponse{
		Status: unhealthy,
		Checks: []check{},
	}
	w.Header().Add("Content-Type", "application/json")

	itemCheck := checkItem(h.config.ItemHost, h.hclient)
	if itemCheck.Status == healthy {
		b.Status = healthy
	}

	b.Checks = append(b.Checks, itemCheck)

	body, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(500)
	}
	if b.Status == unhealthy {
		w.WriteHeader(500)
	}
	fmt.Fprintf(w, string(body))
}

func checkItem(host string, hclient *http.Client) check {
	c := check{
		Name:   "item",
		Error:  "",
		Status: unhealthy,
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/health", host), nil)
	resp, err := hclient.Do(req)
	if err != nil {
		c.Error = err.Error()
		return c
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		c.Status = healthy
		return c
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Error = err.Error()
		return c
	}
	c.Error = string(b)

	return c
}
