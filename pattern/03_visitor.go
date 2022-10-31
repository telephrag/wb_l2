package pattern

import "fmt"

// types visitor are applied to
type Site interface {
	Accept(v *SiteVisitor)
}

type OrdinarySite string

func (s OrdinarySite) Accept(v *SiteVisitor) { v.Visit(&s) }

type BlockedSite string

func (s BlockedSite) Accept(v *SiteVisitor) { v.VisitVPN(&s) }

type LocalSite string

func (s LocalSite) Accept(v *SiteVisitor) { v.VisitLocal(&s) }

// visitor
type SiteVisitor struct{}

func (v *SiteVisitor) Visit(s *OrdinarySite) {
	fmt.Printf("Visiting %s by normal means.\n", *s)
}

func (v *SiteVisitor) VisitVPN(s *BlockedSite) {
	fmt.Printf("Visiting %s using VPN.\n", *s)
}

func (v *SiteVisitor) VisitLocal(s *LocalSite) {
	fmt.Printf("Visiting %s which is inside local network.\n", *s)
}

// demonstation
func init() {
	fmt.Println("Demonstrating Visitor pattern:")
	defer fmt.Println()

	ddg := OrdinarySite("duckduckgo.com")
	rg := BlockedSite("refactoring.guru")
	lh := LocalSite("localhost:8080")

	sites := []Site{ddg, rg, lh}
	v := &SiteVisitor{}

	for _, s := range sites {
		s.Accept(v)
	}
}
