package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3011Primary(t *testing.T){
 got:=blake3.New().BlockSize();want:=64;err:=error(nil);if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
