package tactics

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Keep int

const (
	timeKeep Keep = iota
	atkKeep
)

type buff_ struct {
	attrs_
	Keep
	name   string
	remain int
}

func buff(remain int, attrs ...*attrs_) *buff_ {
	bf := &buff_{}
	for _, attr := range attrs {
		bf.Add(attr)
	}
	bf.remain = remain
	return bf
}

func (b *buff_) key() string {
	segment1, _ := json.Marshal(b.attrs_)
	segment2, _ := json.Marshal(b.Keep)
	segment3, _ := json.Marshal(b.remain)
	// 创建 SHA256 哈希对象
	h := sha256.New()
	// 写入 JSON 数据
	h.Write(segment1)
	h.Write(segment2)
	h.Write(segment3)
	// 获取哈希值
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (b *buff_) IsValid() bool {
	return b.remain > 0
}

func (b *buff_) handle(event Event, _ *Ground) {
	if b.Keep == timeKeep && event.Is(timeGoA) {
		b.remain--
	}
}

func (b *buff_) attr() *attrs_ {
	return &b.attrs_
}
