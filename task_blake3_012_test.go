package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3012Primary(t *testing.T){
 h,err:=blake3.NewKeyed(make([]byte,32));got:=h.Size();want:=32;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3012Boundary(t *testing.T){
 h,err:=blake3.NewKeyed(bytes.Repeat([]byte{1},32));h.Write([]byte("x"));got:=len(h.Sum(nil));want:=32;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
