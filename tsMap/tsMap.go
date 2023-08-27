package tsMap

import (
	"errors"
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

type MapStrMgr struct {
	ObjMgr map[string]interface{}
	count  int
}

func NewMapStrMgr() *MapStrMgr {
	res := new(MapStrMgr)
	res.ObjMgr = make(map[string]interface{})

	return res
}

func (this *MapStrMgr) AddInt64(key int64, obj interface{}) error {
	k := fmt.Sprintf("%d", key)
	return this.Add(k, obj)
}

func (this *MapStrMgr) Add(key string, obj interface{}) error {
	_, ok := this.ObjMgr[key]
	if ok {
		return errors.New("have obj")
	}
	this.ObjMgr[key] = obj
	this.count++
	return nil
}

func (this *MapStrMgr) GetByInt(key int64) interface{} {
	k := fmt.Sprintf("%d", key)
	return this.Get(k)
}

func (this *MapStrMgr) Get(key string) interface{} {
	obj, ok := this.ObjMgr[key]
	if !ok {
		return nil
	}
	if obj == nil {
		return nil
	}
	return obj
}

func (this *MapStrMgr) Have(key string) bool {
	_, ok := this.ObjMgr[key]
	return ok
}

func (this *MapStrMgr) Del(key string) {
	if this.Have(key) {
		delete(this.ObjMgr, key)
		this.count--
	}
}

func (this *MapStrMgr) Count() (count int) {
	return this.count
}

func (this *MapStrMgr) ToArray() (res []interface{}) {
	res = make([]interface{}, 0)
	for _, v := range this.ObjMgr {
		res = append(res, v)
	}
	return
}

func (this *MapStrMgr) Clear() {
	this.ObjMgr = make(map[string]interface{})
	this.count = 0
}

type GlobalDecimal struct {
	sync.RWMutex
	Decimal map[string]decimal.Decimal
}

func (d *GlobalDecimal) ReadMapGlobalDecimal(key string) decimal.Decimal {
	d.RLock()
	value, _ := d.Decimal[key]
	d.RUnlock()
	return value
}

func (d *GlobalDecimal) ReadAndResetMapGlobalDecimal(key string) decimal.Decimal {
	d.RLock()
	value, _ := d.Decimal[key]
	d.Decimal[key] = decimal.NewFromInt(0)
	d.RUnlock()
	return value
}

func (d *GlobalDecimal) ReadAllMapGlobalDecimal() map[string]decimal.Decimal {
	temp := make(map[string]decimal.Decimal)
	d.RLock()
	for key, value := range d.Decimal {
		// value, _ := t.Timestamp[key]
		temp[key] = value
	}
	d.RUnlock()
	return temp
}

func (d *GlobalDecimal) WriteMapGlobalDecimal(key string, value decimal.Decimal) {
	d.Lock()
	d.Decimal[key] = value
	d.Unlock()
}

type GlobalTimestamp struct {
	sync.RWMutex
	Timestamp map[string]int64
}

func (t *GlobalTimestamp) ReadMapGlobalTimestamp(key string) int64 {
	t.RLock()
	value, _ := t.Timestamp[key]
	t.RUnlock()
	return value
}

func (t *GlobalTimestamp) ReadAndResetMapGlobalTimestamp(key string) int64 {
	t.RLock()
	value, _ := t.Timestamp[key]
	t.Timestamp[key] = 0
	t.RUnlock()
	return value
}

func (t *GlobalTimestamp) ReadAllMapGlobalTimestamp() map[string]int64 {
	temp := make(map[string]int64)
	t.RLock()
	for key, value := range t.Timestamp {
		// value, _ := t.Timestamp[key]
		temp[key] = value
	}
	t.RUnlock()
	return temp
}

func (t *GlobalTimestamp) WriteMapGlobalTimestamp(key string, value int64) {
	t.Lock()
	t.Timestamp[key] = value
	t.Unlock()
}
