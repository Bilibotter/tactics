package tactics

import "fmt"

type valid_ interface {
	IsValid() bool
}

type attribute_ interface {
	attr() *attrs_
}

type attach_ interface {
	attribute_
	valid_
}

// 默认有最大生命值增长的英雄，没有最大生命值增益。
func growMax(i int) func(ground *Ground) {
	return func(ground *Ground) {
		champion := ground.champion_
		growth := champion.healthy() * i / 100
		champion.health += growth
		ground.maxHealth = ground.healthy()
		addHealth(growth, champion)
	}
}

// 考虑有固定生命值增长的英雄，可能有最大生命值增益。
func growHp(hp int, skill ...bool) func(*Ground) {
	return func(ground *Ground) {
		champion := ground.champion_
		if len(skill) > 0 {
			hp = hp * ground.castRate() / 100
		}
		old := ground.healthy()
		champion.health += hp
		ground.maxHealth = ground.healthy()
		growth := ground.healthy() - old
		if outputLevel >= 3 {
			fmt.Printf("%d秒:增长最大生命值%d\n", ground.CurrenTime, hp)
		}
		addHealth(growth, champion)
	}
}

func healLose(i int) func(ground *Ground) {
	return func(ground *Ground) {
		champion := ground.champion_
		lose := champion.healthy() - champion.currentHealth
		heal := lose * i * champion.healAmp() / 100 / 100
		addHealth(heal, champion)
		ground.filter(NewE(healedA, heal))
	}
}

func healMax(i int) func(ground *Ground) {
	return func(ground *Ground) {
		champion := ground.champion_
		heal := champion.healthy() * i / 100 * champion.healAmp() / 100
		addHealth(heal, champion)
		ground.filter(NewE(healedA, heal))
	}
}

func healHp(i int, skill ...bool) func(ground *Ground) {
	return func(ground *Ground) {
		champion := ground.champion_
		heal := i * champion.healAmp() / 100
		if len(skill) > 0 {
			heal = heal * ground.castRate() / 100
		}
		addHealth(heal, champion)
		ground.filter(NewE(healedA, heal))
	}
}

func shield(health int, remain int, skill ...bool) func(ground *Ground) {
	return func(ground *Ground) {
		actual := health * ground.shieldAmp() / 100
		s := &shield_{lastSec: remain, health: actual}
		if len(skill) > 0 {
			s.health = s.health * ground.castRate() / 100
			ground.locks = append(ground.locks, &manaLock{s})
		}
		ground.addShield(s)
		ground.filter_ = append(ground.filter_, s)
		ground.filter(NewE(shieldedA, actual))
	}
}

func shieldMax(percent int, remain int, lock ...bool) func(ground *Ground) {
	return func(ground *Ground) {
		actual := ground.healthy() * percent / 100 * ground.shieldAmp() / 100
		s := &shield_{lastSec: remain, health: actual}
		ground.addShield(s)
		ground.filter_ = append(ground.filter_, s)
		if len(lock) > 0 {
			ground.locks = append(ground.locks, &manaLock{s})
		}
		ground.filter(NewE(shieldedA, actual))
	}
}

func addLockBuff(bf *buff_) func(_ *Ground) {
	return func(ground *Ground) {
		cp := *bf
		addBuffFunc(&cp)(ground)
		ground.locks = append(ground.locks, &manaLock{&cp})
		if outputLevel >= 3 {
			fmt.Printf("%d秒:增加buff %s\n", ground.CurrenTime, bf.key()[:6])
		}
	}
}

func addBuffFunc(bf *buff_) func(ground *Ground) {
	return func(ground *Ground) {
		if buf, ok := ground.buffs[bf.key()]; ok {
			buf.remain = bf.remain
		} else {
			ground.buffs[bf.key()] = bf
			ground.attach = append(ground.attach, bf)
			ground.filter_ = append(ground.filter_, bf)
		}
	}
}

func addHealth(health int, champion *champion_) {
	if outputLevel >= 3 {
		fmt.Printf("增加生命值%d, 当前生命值%d, 最大生命值%d\n", health, champion.currentHealth, champion.healthy())
	}
	champion.currentHealth += health
	// 治疗溢出
	if champion.currentHealth > champion.healthy() {
		champion.currentHealth = champion.healthy()
	}
}
