package main
 
import "fmt";
import "rsc.io/quote";
import "test2/utils";
import (
    "reflect"
    "math"
)


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

    var x,y int = 4,5;
    fmt.Println("line34 %T1, %x ", x, y)

    var result = math.Sqrt(float64(x*x + y*y));

    fmt.Println(fmt.Sprintf("line34", result))
    fmt.Println(fmt.Sprintf("Type of result is %d and %d", x, y))
    fmt.Println(fmt.Sprintf("Variable string %d content", x))

    fmt.Println(fmt.Sprintf("Type of result is %T\n", result))
    fmt.Println(fmt.Sprintf("Type of result is %T", reflect.ValueOf(result).Kind()))

    fmt.Println(fmt.Sprintf("!... Hello World ...!"))
    fmt.Println(quote.Go())
    testpackage := utils.TestFunction();
    fmt.Println("test package", testpackage);

    sum := 0
    i := 0

    pointer1 := 42
    p := &pointer1
    fmt.Println("test pointer", p);

    d := &pointer1
    fmt.Println("test pointer", d);

    fmt.Println("test pointer 1", *p, *d);
    *p = 100

    fmt.Println("test pointer 2", *p, *d);



    for ; i < 10; i++ {
        sum += i
        fmt.Println("result", sum);
    }
    fmt.Println("result", sum);
    message := string1("call") + string2("call") + string3("call") 
    fmt.Println("Result ", message, array[0], array[1], array[2]);
}

// defer khai báo, chạy khi function mẹ chuẩn bị end

