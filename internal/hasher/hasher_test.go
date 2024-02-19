package hasher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	type hash_path struct {
		arg      string
		fullPath string
	}
*/
func TestHasher(t *testing.T) {
	/*
		type test struct {
			input get.args
			options []os.Args
			want string
			wantErr error
			shouldSucceed bool
			context string
		}
		possibleArgs := []values{
			{
				ArgStr: "",
				Str: "outputs/d41d8cd98f00b204e9800998ecf8427e.txt"
			},
			{
				ArgStr: "-vvv",
				Str: "outputs/0cc598961dc9ae41055f83d0950544f3.txt"
			},
			{
				ArgStr: "--json",
				Str: "outputs/31d5e31c2e1e0b3c281f5194f130287e.txt"
			}
		}
		tests := []test{

		}
	*/
	a_1 := []string{""}
	h_1 := Hasher(a_1)
	e_1 := "d41d8cd98f00b204e9800998ecf8427e"

	assert.Equal(t, h_1, e_1)
	//fmt.Printf("Returned Hash: %s\n", h_1)
	//fmt.Printf("Expected Hash: %s\n", e_1)

	a_2 := []string{"--json"}
	h_2 := Hasher(a_2)
	e_2 := "31d5e31c2e1e0b3c281f5194f130287e"

	assert.Equal(t, h_2, e_2)
	//fmt.Printf("Returned Hash: %s\n", h_2)
	//fmt.Printf("Expected Hash:\n%s\n", e_2)

	a_3 := []string{"-vvv"}
	h_3 := Hasher(a_3)
	e_3 := "0cc598961dc9ae41055f83d0950544f3"

	assert.Equal(t, h_3, e_3)
	//fmt.Printf("Returned Hash: %s\n", h_3)
	//fmt.Printf("Expected Hash:\n%s\n", e_3)

}
