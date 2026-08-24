package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3003Primary(t *testing.T){
 _,e:=blake3.NewKeyed(make([]byte,32));gotv:=e==nil;want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
