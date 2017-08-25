package main

import (
	"chainOfResponsibility"
	"fmt"
)

func main() {
	CEO := chainOfResponsibility.NewAuditor("CEO", 10000, nil)
	manager := chainOfResponsibility.NewAuditor("manager", 2000, CEO)
	leader := chainOfResponsibility.NewAuditor("leader", 500, manager)

	fmt.Println(leader.Audit(300))   // leader permitted!
	fmt.Println(leader.Audit(3000))  // CEO permitted!
	fmt.Println(leader.Audit(80000)) // out of limit!
}
