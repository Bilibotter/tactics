package tactics

type attrs_ struct {
	health       int
	healthAmp    int // 生命值提升
	speed        int // 攻速
	ability      int // 法强
	castAmp      int // 施法增幅
	recoverAmp   int // 恢复/护盾增幅
	manaReduce   int // 施法时减少的法力值
	manaShrink   int // 施法时减少的法力值百分比
	manaAmp      int // 法力值提升
	dmgTaken     int // 承伤百分比。使用耐久度不方便计算
	armor        int // 抗性
	block        int // 格挡
	armorAmp     int // 抗性增幅
	regeneration int // 最大生命值恢复
	missingRegen int // 基于损失的最大生命值恢复
}

func HP(i int) *attrs_ {
	return &attrs_{health: i}
}

func AS(i int) *attrs_ {
	return &attrs_{speed: i}
}

func AP(i int) *attrs_ {
	return &attrs_{ability: i}
}

func Heal(i int) *attrs_ {
	return &attrs_{recoverAmp: i}
}

func Reduce(i int) *attrs_ {
	return &attrs_{manaReduce: i}
}

func Shrink(i int) *attrs_ {
	return &attrs_{manaShrink: i}
}

func HpAmp(i int) *attrs_ {
	return &attrs_{healthAmp: i}
}

func ManaAmp(i int) *attrs_ {
	return &attrs_{manaAmp: i}
}

// 耐久度
func DR(i int) *attrs_ {
	return &attrs_{dmgTaken: 100 - i}
}

func AR(i int) *attrs_ {
	return &attrs_{armor: i}
}

func BLK(i int) *attrs_ {
	return &attrs_{block: i}
}

func (a *attrs_) Add(attrs *attrs_) {
	a.health += attrs.health
	a.healthAmp += attrs.healthAmp
	a.speed += attrs.speed
	a.ability += attrs.ability
	a.recoverAmp += attrs.recoverAmp
	a.manaReduce += attrs.manaReduce
	a.manaShrink += attrs.manaShrink
	a.manaAmp += attrs.manaAmp
	if attrs.dmgTaken > 0 {
		a.dmgTaken = a.dmgTaken * attrs.dmgTaken / 100
	}
	if a.dmgTaken == 0 {
		a.dmgTaken = attrs.dmgTaken
	}
	a.armor += attrs.armor
	a.block += attrs.block
	a.armorAmp += attrs.armorAmp
	a.regeneration += attrs.regeneration
	a.missingRegen += attrs.missingRegen
}
