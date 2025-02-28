package output_tests_no_ka

import (
	_ "net/http/pprof"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ravinggo/tools/keepalive-gen/output_tests_no_ka/output"
)

func TestKeepalive(t *testing.T) {
	r := require.New(t)
	ud := output.UserData{}
	ud.Init()

	u := ud.GetUser()
	u.Name = "xxx"
	ud.SaveUser()

	ud.Commit()

	u = ud.GetUser()
	if u.Name != "xxx" {
		t.Fatalf(`u.Name != "xxx"`)
	}
	u.Name = "yyy"
	ud.Rollback()

	u = ud.GetUser()
	if u.Name != "xxx" {
		t.Fatalf(`u.Name != "xxx"`)
	}
	u.Name = "yyy"
	ud.SaveUser()
	ud.Rollback()
	u = ud.GetUser()
	if u.Name != "xxx" {
		t.Fatalf(`u.Name != "xxx"`)
	}
	u.Name = "yyy"
	ud.SaveUser()
	ud.Commit()

	u = ud.GetUser()
	if u.Name != "yyy" {
		t.Fatalf(`u.Name != "yyy"`)
	}
	_, ok := ud.GetItem(1)
	r.False(ok)
	originItem := output.Item{
		ID:    1,
		CID:   2,
		Value: 3,
	}
	ud.SaveItem(originItem)
	item, ok := ud.GetItem(1)
	r.True(ok)
	r.Equal(originItem, item)
	ud.Commit()
	item, ok = ud.GetItem(1)
	r.True(ok)
	r.Equal(originItem, item)
}

func DoOnce(b *testing.B) {
	ud := output.UserData{}
	ud.Init()
	for i := 0; i < 100; i++ {
		u := ud.GetUser()
		if u.Age != i {
			b.Fatalf("u.Age=%d,expected:%d", u.Age, i)
		}
		u.Age = i + 1
		ud.SaveUser()
		_, ok := ud.GetItem(int64(i))
		if ok {
			b.Fatalf("ud.GetItem(int64(%d)) found, expected not found", i)
		}
		ud.SaveItem(
			output.Item{
				ID:    int64(i),
				CID:   0,
				Value: 0,
			},
		)
		if i > 0 {
			_, ok := ud.GetItem(int64(i - 1))
			if !ok {
				b.Fatalf("ud.GetItem(int64(%d)) not found, expected found", i-1)
			}
			ud.DeleteItem(int64(i - 1))
		}
		ud.SaveHistory(
			output.History{
				A: i,
				B: 0,
				C: "",
			},
		)
		if i > 0 {
			ud.DeleteHistory(i - 1)
		}
		_, ok = ud.GetHistory(i)
		if !ok {
			b.Fatalf("ud.GetHistory(%d) not found", i)
		}

		_, ok = ud.GetFight(int64(i))
		if ok {
			b.Fatalf("ud.GetFight(int64(%d)) found, expected not found", i)
		}
		ud.SaveFight(
			&output.Fight{
				A: int64(i),
				B: 0,
				C: nil,
			},
		)
		if i > 0 {
			ud.DeleteFight(int64(i - 1))
		}

		_, ok = ud.GetDomain(int64(i))
		if ok {
			b.Fatalf("ud.GetDomain(int64(%d)) found, expected not found", i)
		}
		ud.SaveDomain(
			&output.Domain{
				A: int64(i),
				B: 0,
				C: nil,
			},
		)
		if i > 0 {
			ud.DeleteDomain(int64(i - 1))
		}
		itemCount := 0
		ud.RangeItems(
			func(k int64, v output.Item) bool {
				itemCount++
				return true
			},
		)
		if itemCount != 1 {
			b.Fatalf("ud.RangeItems count != 1")
		}
		if ud.LenItem() != 1 {
			b.Fatal("ud.LenItem() != 1")
		}
		countHistory := 0
		ud.RangeHistorys(
			func(v *output.History) bool {
				countHistory++
				return true
			},
		)
		if countHistory != 1 {
			b.Fatalf("ud.RangeHistorys count != 1")
		}
		err := ud.DoDB(&TestUserDataDB{})
		if err != nil {
			b.Fatal(err)
		}
		ud.Commit()
		if i > 0 {
			if ud.LenItem() != 1 {
				b.Fatal("ud.LenItem() != 1")
			}
		}

		ud.RangeDomains(
			func(v *output.Domain) bool {
				return true
			},
		)
		ud.RangeFights(
			func(v *output.Fight) bool {
				return true
			},
		)
		ud.RangeHistorys(
			func(v *output.History) bool {
				return true
			},
		)
	}
}

func BenchmarkKeepalive(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		DoOnce(b)
	}
}

type TestUserDataDB struct {
}

func (TestUserDataDB) SaveUser(isNew bool, d *output.User) error       { return nil }
func (TestUserDataDB) SaveDomain(isNew bool, d *output.Domain) error   { return nil }
func (TestUserDataDB) DeleteDomain(d *output.Domain) error             { return nil }
func (TestUserDataDB) SaveFight(isNew bool, d *output.Fight) error     { return nil }
func (TestUserDataDB) DeleteFight(d *output.Fight) error               { return nil }
func (TestUserDataDB) SaveHistory(isNew bool, d *output.History) error { return nil }
func (TestUserDataDB) DeleteHistory(d *output.History) error           { return nil }
func (TestUserDataDB) SaveItem(isNew bool, d *output.Item) error       { return nil }
func (TestUserDataDB) DeleteItem(d *output.Item) error                 { return nil }
func (TestUserDataDB) SavePlayer(isNew bool, d *output.Player) error   { return nil }
func (TestUserDataDB) DeletePlayer(d *output.Player) error             { return nil }
