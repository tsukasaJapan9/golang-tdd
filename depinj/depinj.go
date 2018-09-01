package depinj

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "http world")
}

func main() {
	//Greet(os.Stdout, "taro")
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))

	if err != nil {
		fmt.Println(err)
	}
}