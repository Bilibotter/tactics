package tactics

type manaLock struct {
	valid_ // 护盾或buff
}

type atkLock struct {
	atkTimes int
}

func newLock(val any) manaLock {
	switch obj := val.(type) {
	case *buff_:
		return manaLock{obj}
	case *shield_:
		return manaLock{obj}
	case int:
		return manaLock{&atkLock{obj}}
	}
	panic("invalid lock type")
}

func (m *manaLock) handle(event Action, _ *Ground) {
	if event == attackA {
		if lock, ok := m.valid_.(*atkLock); ok {
			lock.atkTimes--
		}
	}
}

func (a *atkLock) IsValid() bool {
	return a.atkTimes > 0
}
