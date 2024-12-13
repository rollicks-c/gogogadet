package funk

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Ordered interface {
	constraints.Ordered | time.Duration
}
