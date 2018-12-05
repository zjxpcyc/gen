package gen_test

import (
	"testing"

	"github.com/zjxpcyc/gen"
)

func TestMD5(t *testing.T) {
	str := "I'm Yansen"
	exp1 := "e11a48f6707dd0e64c07d6d1daba6b3f"
	salt := "hello"
	exp2 := "a68b8759b177b09277fc1d3fddca70ef"

	res := gen.MD5(str)

	if res != exp1 {
		t.Fatalf("Test md5 no salt fail")
	}

	res = gen.MD5(str, salt)
	if res != exp2 {
		t.Fatalf("Test md5 with salt fail")
	}
}

func TestSHA1(t *testing.T) {
	str := "I'm Yansen"
	exp1 := "238527219095a39db93e54757bcde3529e0560d4"

	res := gen.SHA1(str)

	if res != exp1 {
		t.Fatalf("Test SHA1 fail")
	}
}

func TestHmacSHA256(t *testing.T) {
	str := "I'm Yansen"
	key := "hello"
	exp1 := "0db762fd8b65dbf9869f666d70e1cff39867da227bc922ea1df2703a68541302"

	res := gen.HmacSHA256(str, key)

	if res != exp1 {
		t.Fatalf("Test HmacSHA256 fail")
	}
}
