package tactics

import "fmt"

type passive struct {
	attrs_
	trigger  Action
	stack    *attrs_ // 叠加数值
	ground   *Ground
	stacks   int
	maxStack int
	freq     int             // 每多少秒触发一次
	count    int             // 已接收事件数，和freq结合使用
	start    int             // 界定的数值/左边界值
	end      int             // 右边界值
	call     func(g *Ground) // 匹配到事件时调用
	once     int             // once:0表示无限次,1表示只触发一次,2表示已触发
}

func (p *passive) handle(event Event, g *Ground) {
	if !event.Is(p.trigger) {
		return
	}
	if p.maxStack > 0 && p.stacks < p.maxStack && event.Is(p.trigger) {
		if outputLevel >= 3 {
			fmt.Printf("第%d次叠加被动\n", p.stacks+1)
		}
		p.stacks++
		p.Add(p.stack)
	}
	if p.call == nil {
		return
	}
	if p.once == 2 {
		return
	}
	if p.freq != 0 {
		p.count++
		if p.count%p.freq != 0 {
			return
		}
	}
	if p.once == 1 {
		p.once = 2
	}
	if outputLevel >= 3 {
		fmt.Printf("%d秒:触发被动\n", g.CurrenTime)
	}
	p.call(g)
}

func (p *passive) IsValid() bool {
	// 可叠加被动常住生效
	if p.maxStack > 0 {
		return true
	}
	if p.once >= 1 {
		return p.once == 1
	}
	if p.trigger == shieldedA {
		for _, sh := range p.ground.shields {
			if sh.IsValid() {
				return true
			}
		}
		return false
	}
	if p.trigger == healthPercentA {
		if p.end == 0 {
			p.end = 101
		}
		// 未初始化
		if p.ground.maxHealth == 0 {
			return false
		}
		percent := p.ground.currentHealth * 100 / p.ground.maxHealth
		return percent >= p.start && percent <= p.end
	}
	if p.trigger == timeLineA {
		if p.end == 0 {
			p.end = 30
		}
		return p.ground.CurrenTime >= p.start && p.ground.CurrenTime < p.end
	}
	return true
}

func (p *passive) attr() *attrs_ {
	return &p.attrs_
}
