package consul

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/zhufuyi/sponge/pkg/registry"
)

func newConsulRegistry() *Registry {
	consulClient, err := api.NewClient(&api.Config{})
	if err != nil {
		panic(err)
	}

	r := New(consulClient, WithHealthCheck(true))

	return r
}

func TestRegistry_Register(t *testing.T) {
	r := newConsulRegistry()
	instance := registry.NewServiceInstance("foo", []string{"grpc://127.0.0.1:8282"})

	err := r.Register(context.Background(), instance)
	t.Log(err)

	_, err = r.ListServices()
	t.Log(err)

	_, err = r.GetService(context.Background(), "foo")
	t.Log(err)

	_, err = r.Watch(context.Background(), "foo")
	t.Log(err)

	go func() {
		r.resolve(newServiceSet())
	}()

	err = r.Deregister(context.Background(), instance)
	t.Log(err)

	time.Sleep(time.Millisecond * 100)
}