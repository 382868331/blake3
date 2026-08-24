package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3018Primary(t *testing.T){
 buf:=make([]byte,32);got,err:=blake3.New().Digest().Read(buf);want:=32;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3018Boundary(t *testing.T){
 buf:=make([]byte,1);got,err:=blake3.New().Digest().Read(buf);want:=1;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
