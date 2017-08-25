package chainOfResponsibility

type auditor struct {
	name      string
	authority int
	superior  *auditor
}

func (auditor auditor) Audit(request int) string {
	if request <= auditor.authority {
		return auditor.name + " permitted!"
	} else {
		if auditor.superior == nil {
			return "out of limit!"
		} else {
			return auditor.superior.Audit(request)
		}
	}
}

func NewAuditor(name string, authority int, superior *auditor) *auditor {
	auditor := &auditor{name: name, authority: authority, superior: superior}

	return auditor
}
