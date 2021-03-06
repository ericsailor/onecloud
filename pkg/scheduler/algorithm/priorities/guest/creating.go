package guest

import (
	"yunion.io/x/onecloud/pkg/scheduler/algorithm/priorities"
	"yunion.io/x/onecloud/pkg/scheduler/core"
)

type CreatingPriority struct {
	priorities.BasePriority
}

func (p *CreatingPriority) Name() string {
	return "creating"
}

func (p *CreatingPriority) Clone() core.Priority {
	return &CreatingPriority{}
}

func (p *CreatingPriority) Map(u *core.Unit, c core.Candidater) (core.HostPriority, error) {

	h := priorities.NewPriorityHelper(p, u, c)

	hc, err := p.HostCandidate(c)
	if err != nil {
		return core.HostPriority{}, err
	}

	if hc.CreatingGuestCount > 0 {
		score := -int(hc.CreatingGuestCount)
		h.SetFrontScore(score)
	}

	return h.GetResult()
}
