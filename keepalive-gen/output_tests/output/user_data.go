// +k8s:deepcopy-gen=package

package output

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ravinggo/tools/keepalive-gen/keepalive"
)

type User struct {
	ID   int `bson:"_id" key_id`
	Name string
	Age  int
}

type Item struct {
	ID    int64 `bson:"_id" key_id`
	CID   int64
	Value int64
}

type Player struct {
	PID  int64 `bson:"_id" key_id`
	User User
}

type History struct {
	A int `bson:"_id" key_id`
	B int64
	C string
}

type Domain struct {
	A bson.ObjectID `bson:"_id" key_id`
	B int
	C []string
}

type Fight struct {
	A int64 `bson:"_id" key_id`
	B int
	C []string
}

// +k8s:deepcopy-gen=false
//
// +game:keepalive-gen=true
type UserData struct {
	keepalive.KeepAlive
	user      User
	items     map[int64]Item
	playerMap map[int64]*Player
	histories []History
	domains   []Domain
	fights    []*Fight
}
