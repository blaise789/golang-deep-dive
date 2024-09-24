package main

/*import (
	"encoding/json"
	"fmt"
	"os"
	// "os"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	res1D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	var d map[string]interface{}
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	if err := json.Unmarshal(byt, &d); err != nil {
		panic(err)
	}
	fmt.Println(d)
	num := d["num"]
	fmt.Println(num)
	fmt.Printf("type %T",num)
     enc:=json.NewEncoder(os.Stdout)
	 b:=map[string] int{"apple":5,"lettuce":7}
	 enc.Encode(b)
	 

}
*/