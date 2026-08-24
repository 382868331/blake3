package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3010Primary(t *testing.T){
 h:=blake3.New();h.Write([]byte("base"));c:=h.Clone();c.Write([]byte("clone"));fresh:=blake3.New();fresh.Write([]byte("base"));gotv:=bytes.Equal(h.Sum(nil),fresh.Sum(nil));want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
