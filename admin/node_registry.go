package admin

import (
	. "github.com/Huhaokun/let-it-fail/log"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
)

type NodeRegistry struct {
	k8sClient *kubernetes.Clientset
	nodes     map[string]Node
	mu        sync.Mutex
}

func NewNodeRegistry(k8sClient *kubernetes.Clientset) *NodeRegistry {

	r := &NodeRegistry{
		k8sClient: k8sClient,
		nodes:     make(map[string]Node),
	}

	err := r.Init()
	if err != nil {
		log.Fatalf("fail to init node regisry")
	}

	return r
}

func (n *NodeRegistry) Init() error {
	nodeList, err := n.k8sClient.CoreV1().Nodes().List(v1.ListOptions{})
	if err != nil {
		Log.Errorf("fail to list k8s node due to %v", err)
		return err
	}

	for _, node := range nodeList.Items {
		n.addNode(&node)
	}

	wch, err := n.k8sClient.CoreV1().Nodes().Watch(v1.ListOptions{
		Watch: true,
	})
	if err != nil {
		Log.Errorf("fail to watch k8s node due to %v", err)
		return err
	}

	// TODO stop the go-routine
	go func() {
		for {
			event := <-wch.ResultChan()
			n.HandleNodeEvent(&event)
		}
	}()

	return nil
}

func (n *NodeRegistry) HandleNodeEvent(event *watch.Event) {
	// TODO figure it out
	node := event.Object.(*v12.Node)
	if event.Type == watch.Added {
		Log.Infof("add new node %v", node.Name)
		n.addNode(node)
	} else if event.Type == watch.Deleted {
		Log.Infof("delete node %v", node.Name)
	}
}

func (n *NodeRegistry) Get(id string) Node {
	if node, ok := n.nodes[id]; ok {
		return node
	} else {
		return nil
	}
}

func (n *NodeRegistry) List() []Node {
	var nodes []Node
	for _, node := range n.nodes {
		nodes = append(nodes, node)
	}

	return nodes
}

func (n *NodeRegistry) addNode(k8sNode *v12.Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, addr := range k8sNode.Status.Addresses {
		if addr.Type == v12.NodeExternalIP {
			Log.Infof("add node %s to cluster", addr.Address)
			n.nodes[addr.Address] = NewNode(addr.Address, 7999)
			break
		}
	}
}

func (n *NodeRegistry) removeNode(k8sNode *v12.Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, addr := range k8sNode.Status.Addresses {
		if addr.Type == v12.NodeExternalIP {
			Log.Infof("remove node %s from cluster", addr.Address)
			delete(n.nodes, addr.Address)
			break
		}
	}
}
