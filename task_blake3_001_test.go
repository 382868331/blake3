package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3001Primary(t *testing.T){
 got:=blake3.New().Size();want:=32;err:=error(nil);if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3001Boundary(t *testing.T){
 h:=blake3.New();h.Write([]byte("x"));got:=len(h.Sum(nil));want:=32;err:=error(nil);if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
