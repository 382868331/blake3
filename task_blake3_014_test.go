package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3014Primary(t *testing.T){
 d:=blake3.New().Digest();buf:=make([]byte,128);d.Read(buf);gotv:=!bytes.Equal(buf[:64],buf[64:]);want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
