package strategy

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/docker/swarm/cluster"
	"github.com/docker/swarm/scheduler/node"
)

// MyOwnPlacementStrategy randomly places the container into the cluster.
type MyOwnPlacementStrategy struct {
	r *rand.Rand
}

// Initialize a MyOwnPlacementStrategy.
func (p *MyOwnPlacementStrategy) Initialize() error {
	p.r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return nil
}

// Name returns the name of the strategy.
func (p *MyOwnPlacementStrategy) Name() string {
	return "myownstrategy"
}

// RankAndSort randomly sorts the list of nodes.
func (p *MyOwnPlacementStrategy) RankAndSort(config *cluster.ContainerConfig, nodes []*node.Node) ([]*node.Node, error) {
	fmt.Printf("Pass par la Cool")
	for i := len(nodes) - 1; i > 0; i-- {
		j := p.r.Intn(i + 1)
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
	return nodes, nil
}
