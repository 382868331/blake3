package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3015Primary(t *testing.T){
 d:=blake3.New().Digest();all:=make([]byte,40);d.Read(all);d2:=blake3.New().Digest();_,err:=d2.Seek(7,io.SeekStart);got:=make([]byte,8);d2.Read(got);want:=all[7:15];if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3015Boundary(t *testing.T){
 h:=blake3.New();h.Write([]byte("abc"));base:=h.Digest();all:=make([]byte,50);base.Read(all);d:=h.Digest();_,err:=d.Seek(32,io.SeekStart);got:=make([]byte,10);d.Read(got);want:=all[32:42];if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
