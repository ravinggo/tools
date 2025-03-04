// +k8s:deepcopy-gen=package

package output

type User struct {
	ID   int `key_id`
	Name string
	Age  int
}

type Item struct {
	ID    int64 `key_id`
	CID   int64
	Value int64
}

type Player struct {
	PID  int64 `key_id`
	User User
}

type History struct {
	A int `key_id`
	B int64
	C string
}

type Domain struct {
	A int64 `key_id`
	B int
	C []string
}

type Fight struct {
	A int64 `key_id`
	B int
	C []string
}
