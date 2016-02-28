// https://developer.mozilla.org/en-US/docs/Web/API/Performance
package performance

import (
	"github.com/flimzy/js/performancenavigation"
	"github.com/flimzy/js/performancetiming"
	"github.com/gopherjs/gopherjs/js"
)

type Performance struct {
	js.Object
}

type PerformanceNavigation struct {
	js.Object
}

func (p *Performance) Navigation() *performancenavigation.PerformanceNavigation {
	n := p.Get("navigation")
	return &performancenavigation.PerformanceNavigation{*n}
}

func (p *Performance) Timing() *performancetiming.PerformanceTiming {
	t := p.Get("timing")
	return &performancetiming.PerformanceTiming{*t}
}
