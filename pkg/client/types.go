package client

import (
	"encoding/json"
)

type StatusUpdate struct {
	StatusUpdate StatusUpdateResponse `graphql:"statusUpdate(input: $input)"`
}

type StatusUpdateInput struct {
	// The node ID for this status.
	NodeID string `graphql:"nodeID" json:"nodeID"`
	// The namespace ID for this status.
	NamespaceID string `graphql:"namespaceID" json:"namespaceID"`
	// The source for this status.
	Source string `graphql:"source" json:"source"`
	// The data to save in this status.
	Data json.RawMessage `graphql:"data" json:"data"`
}

type StatusUpdateResponse struct {
	Status struct {
		ID                string          `graphql:"id"`
		Data              json.RawMessage `graphql:"data"`
		Source            string          `graphql:"source"`
		StatusNamespaceID string          `graphql:"statusNamespaceID"`

		Metadata struct {
			ID     string `graphql:"id"`
			NodeID string `graphql:"nodeID"`
		} `graphql:"metadata"`
	} `graphql:"status"`
}

// OriginNode is a struct that represents the OriginNode GraphQL type
type OriginNode struct {
	ID         string
	Name       string
	Target     string
	PortNumber int64
	Active     bool
}

// OriginEdges is a struct that represents the OriginEdges GraphQL type
type OriginEdges struct {
	Node OriginNode
}

// Origins is a struct that represents the Origins GraphQL type
type Origins struct {
	Edges []OriginEdges
}

// Pool is a struct that represents the Pool GraphQL type
type Pool struct {
	ID       string
	Name     string
	Protocol string
	Origins  Origins
}

// PortNode is a struct that represents the PortNode GraphQL type
type PortNode struct {
	ID     string
	Name   string
	Number int64
	Pools  []Pool
}

// PortEdges is a struct that represents the PortEdges GraphQL type
type PortEdges struct {
	Node PortNode
}

// Ports is a struct that represents the Ports GraphQL type
type Ports struct {
	Edges []PortEdges
}

// OwnerNode is a struct that represents the OwnerNode GraphQL type
type OwnerNode struct {
	ID string
}

// LocationNode is a struct that represents the LocationNode GraphQL type
type LocationNode struct {
	ID string
}

// LoadBalancer is a struct that represents the LoadBalancer GraphQL type
type LoadBalancer struct {
	ID          string
	Name        string
	Owner       OwnerNode
	Location    LocationNode
	IPAddresses []IPAddress `graphql:"IPAddresses" json:"IPAddresses"`
	Ports       Ports
}

// GetLoadBalancer is a struct that represents the GetLoadBalancer GraphQL query
type GetLoadBalancer struct {
	LoadBalancer LoadBalancer `graphql:"loadBalancer(id: $id)"`
}

// IPAddress is a struct that represents the IPAddress GraphQL type
type IPAddress struct {
	ID       string
	IP       string
	Reserved bool
}

// Readable version of the above:
// type GetLoadBalancer struct {
// 	LoadBalancer struct {
// 		ID    string
//      Owner string
// 		Name  string
//		IPAddressableFragment
// 		Ports struct {
// 			Edges []struct {
// 				Node struct {
// 					Name   string
// 					Number int64
// 					Pools  []struct {
// 						Name     string
// 						Protocol string
// 						Origins  struct {
// 							Edges []struct {
// 								Node struct {
// 									Name       string
// 									Target     string
// 									PortNumber int64
// 									Active     bool
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	} `graphql:"loadBalancer(id: $id)"`
// }
