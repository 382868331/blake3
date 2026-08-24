package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3016Primary(t *testing.T){
 d:=blake3.New().Digest();_,e:=d.Seek(0,io.SeekEnd);got:=e!=nil;want:=true;err:=error(nil);if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3016Boundary(t *testing.T){
 d:=blake3.New().Digest();_,e:=d.Seek(4,io.SeekEnd);got:=e!=nil;want:=true;err:=error(nil);if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
