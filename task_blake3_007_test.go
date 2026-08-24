package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3007Primary(t *testing.T){
 h:=blake3.New();got,_:=h.Write([]byte("abc"));want:=3;err:=error(nil);if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3007Boundary(t *testing.T){
 h:=blake3.New();got,_:=h.Write(nil);want:=0;err:=error(nil);if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
