package main

import "fmt"

// Dependency Inversion Principle
// 1: HLM(high level module) should not depend on LLM(low level module)
// HLM usually is system level logic like ORM, Persistence DB or some infra level code
// LLM usually is upper level business logic
// If High depends on Low means low change code, high also need to change, so we should break this dependency

// 2: Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// other useful stuff here
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-level module, could be repository level code, like query database
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations,
		Info{parent, Parent, child})
	rs.relations = append(rs.relations,
		Info{child, Child, parent})
}

// high-level module: business logic type
// define a interface as member, which low level type will implement it
// since low level type implement this interface can assign low level object to member and call interface function
// this will invoke low level function through common interface
type Research struct {
	// if we depend on low level type we are breaking DIP
	// relationships Relationships

	browser RelationshipBrowser // low-level
}

func (r *Research) Investigate() {
	// if we depend on low level type we are breaking DIP
	// relations := r.relationships.relations
	// for _, rel := range relations {
	//	if rel.from.name == "John" &&
	//		rel.relationship == Parent {
	//		fmt.Println("John has a child called", rel.to.name)
	//	}
	// }

	// we change to high level use abstraction layer (in go is interface) to invoke low level method
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.Investigate()
}
