package main
 
import "fmt";
import "rsc.io/quote"
import "test/utils"
func string1(tesname string) string {
    message := tesname + "test1";
    return message;
}

func string2(tesname string) string {
    message := tesname + "test2";
    return message;
}

func string3(tesname string) string {
    message := tesname + "test3";
    return message;
}

func test4(tesname string) string {
    message := tesname + "test3";
    return message;
}
// Main function
func main() {
    array := [3]string{"test1", "test2", "test3"}
    for i := 0; i < len(array); i++ {
        fmt.Println(array[i])
    }

    fmt.Println("!... Hello World ...!")
    fmt.Println(quote.Go())
    
    message := string1("call") + string2("call") + string3("call") 
    fmt.Println("Result ", message, array[0], array[1], array[2]);
}

