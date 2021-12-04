package aoc

import (
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2020/pkg/util"
)

func Day7(path string) (int, int) {

	lines := util.ReadToStrings(path)

	r := regexp.MustCompile("(([0-9]+) )?([a-z]+ [a-z]+) bags?")
	rules := &rules{bags: make(map[string]*bag)}

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)

		if len(m) == 0 {
			continue
		}

		// Create a new bag
		parentBag := m[0][3]
		b := rules.GetBag(parentBag)
		b.validParent = true

		for i := 1; i < len(m); i++ {

			numChildren := m[i][2]
			childName := m[i][3]

			// Check for 'no other'
			if numChildren == "" {
				continue
			}

			n, err := strconv.Atoi(numChildren)
			if err != nil {
				panic(err)
			}

			child := rules.GetBag(childName)
			b.children[child.name] = n
			child.parents = append(child.parents, b.name)
		}

	}

	// Find parents for "shiny gold"
	count := rules.CountParents("shiny gold")

	// This function will include the top level bag too
	children := rules.CountChildren("shiny gold") - 1

	return count, children
}

type bag struct {
	name        string
	validParent bool
	children    map[string]int
	parents     []string
}

type rules struct {
	bags map[string]*bag
}

func (r *rules) GetBag(bagname string) *bag {
	if _, ok := r.bags[bagname]; !ok {
		b := &bag{name: bagname, children: make(map[string]int), parents: make([]string, 0)}
		r.bags[bagname] = b
	}

	return r.bags[bagname]
}

func (r *rules) CountChildren(b string) int {

	if _, ok := r.bags[b]; !ok {
		return 0
	}

	count := 1
	for k, v := range r.bags[b].children {
		count += r.CountChildren(k) * v
	}

	return count
}

func (r *rules) CountParents(b string) int {

	seen := make(map[string]bool)
	parents := make(map[string]bool)
	look := r.bags[b].parents

	for len(look) > 0 {
		l1 := []string{}

		for _, x := range look {
			if _, ok := seen[x]; ok {
				continue
			} else {
				seen[x] = true
			}

			//Add the parents
			if _, ok := r.bags[x]; ok {
				if r.bags[x].validParent {
					parents[x] = true
				}

				l1 = append(l1, r.bags[x].parents...)
			}
		}

		look = l1
	}

	return len(parents)
}
