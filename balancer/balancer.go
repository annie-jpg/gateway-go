// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balancer

import (
	"GateWay/config"
	"errors"
)

var (
	NoHostError                = errors.New("no host")
	AlgorithmNotSupportedError = errors.New("algorithm not supported")
)

// Balancer interface is the load balancer for the reverse proxy
type Balancer interface {
	Add(string)
	Remove(string)
	Balance(string) (string, error)
	Inc(string)
	Done(string)
}

// Factory is the factory that generates Balancer,
// and the factory design pattern is used here
type Factory func([]string) Balancer

var factories = make(map[string]Factory)

// Build generates the corresponding Balancer according to the algorithm
func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, AlgorithmNotSupportedError
	}
	return factory(hosts), nil
}

func InitP2b(locations []*config.Location) error {
	for _, l := range locations {
		b, err := Build(l.BalanceMode, l.ProxyPass)
		if err != nil {
			return err
		}
		p := l.Pattern
		if p[len(p)-7:] == "*action" || p[len(p)-7:] == ":action" {
			p = p[:len(p)-7]
		}
		if _, ok := P2bMap[p]; ok {
			return errors.New("Location have same patterns : " + p)
		}
		P2bMap[p] = b
	}
	return nil
}