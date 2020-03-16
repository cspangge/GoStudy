/*
Program Name: First Go Program
Version: 1.0
Description: Go Program basic program structure
Author: pangge
Date Created: 15-03-2020
Date Modified: 15-03-2020
*/
/*
Go is statically typed programming language.
Variables always have a specific type and that type cannot change.
*/
package main

// Preprocessor Statements
import (
    "errors"
    "fmt"
    "math"
)

// User Defined Functions
func welcome() {
    fmt.Println("Hello World")
}

// Data Type Study
func dataTypeStudy() {
    /*
        Built-in data types
        Numbers : Hold numeric values
            Integer:
                Signed Integer
                Type    Size        Description                         Range
                Int8    8 bits      8 bit signed integer                -128 to 127
                Int16   16 bits     16 bit signed integer               -2^15 to 2^15-1
                Int32   32 bits     32 bit signed integer               -2^31 to 2^31-1
                Int64   64 bits     64 bit signed integer, octal,hex    -2^63 to 2^63-1
                Int     Platform Dependent, 32-bit system and 64-bit system
                Unsigned Integer
                Type    Size        Description                         Range
                uint8   8 bits      8 bit unsigned integer              0 to 127
                uint16  16 bits     16 bit unsigned integer             0 to 2^16-1
                uint32  32 bits     32 bit unsigned integer             0 to 2^32-1
                uint64  64 bits     64 bit unsigned integer             0 to 2^64-1
                uint    Platform Dependent, 32-bit system and 64-bit system
                Octal Number:           0 prefix
                Hexadecimal Number:     0x prefix
                Other Type
                Type                Description
                byte                alias for uint8, represented with ASCII value
                rune                alias for int32, used to represent characters, Unicode value encoded in UTF-8 format.
                uintptr             hold memory address pointers
            Float:
                float32:    32-bit
                float64:    64-bit  Default for untyped float value
                var num1 float32 = 3.1415
                var num2 = 3.1415   ==> Type inferred as float64
            Complex Numbers:
                complex64:  numbers with imaginary part, use float32 and float64 represent their real and imaginary parts.
                complex128: numbers with imaginary part, use float32 and float64 represent their real and imaginary parts.
                var num1 = 3 + 7i   ==> Type inferred as complex128
        Boolean
            var flag1 = true
            var flag2 = false
        String
            Hold series or sequence of characters: letters, numbers and special characters.
            var str1 = "Hello World"
            var str2 = `Hello World, this is
            a multi-line text string.`  ==> Doesn't support escape characters.
    */
    var str1 = `Hello World, this is
a multi-line text string. \n doesn't work here.`
    fmt.Println(str1)
}

// Variable Study
func variableStudy() {
    /*
        Declaration Syntax
        var <name> <type>
        or
        var <name> <type> = <expression>
        or
        <name> := <value>
    */
    var str1 string
    str1 = "Hello World"
    var str2 string = "Good Morning"
    str3 := "Good Afternoon"
    fmt.Println(str1)
    fmt.Println(str2)
    fmt.Println(str3)
}

// Operator Study
func operatorStudy() {
    /*
        Go supports the following types of operator:
        1. Arithmetic Operators
            Operator    Name            Example
            +           Addition        a+b
            -           Subtraction     a-b
            *           Multiply        a*b
            /           Division        a/b
            %           Modulus         a%b
    */
    var a int = 50
    var b int = 30
    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Println(a / b)
    fmt.Println(a % b)
    /*
        2. Assignment Operators
            Operator    Description                         Expression
            =           Assignment Operator                 a=b
            +=          add and assign                      a+=b    =>    a=a+b
            -=          subtract and assign                 a-=b    =>    a=a-b
            *=          multiply and assign                 a*=b    =>    a=a*b
            /=          divide and assign                   a/=b    =>    a=a/b
            %=          mod and assign                      a%=b    =>    a=a%b
            <<=         left shift and assign               a<<=5   =>    a=a<<5
            >>=         right shift and assign              a>>=5   =>    a=a>>5
            &=          bitwise AND assign                  a&=5    =>    a=a&5
            ^=          bitwise exclusive OR and assign     a^=5    =>    a=a^5
            |=          bitwise inclusive OR and assign     a|=5    =>    a=a|5
        3. Comparison(Relational) Operators
            Operator    Description                         Example
            >           greater than                        a>b
            <           less than                           a<b
            >=          greater than or equal to            a>=b
            <=          less than or equal to               a<=b
            ==          equal to                            a==b
            !=          not equal to                        a!=b
    */
    a = 20
    b = 50
    if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a is less than b")
    }
    /*
        4. Logical Operators
            Operator    Description                         Example
            &&          Logical AND
            ||          Logical OR
            !           Logical NOT
        5. Bitwise Operators
            Operator    Description
            &           bitwise AND
            |           bitwise OR
            ^           bitwise XOR
            &^          bit clear ( AND NOT )
        6. Other Operators
            Operator    Name                        Description
            &           Address of                  &a generates a pointer to a
            *           Pointer to                  *a denotes the variable pointed to by a
            <-          Receive Operator            <-ch is the values received from channel ch
        7. String concatenation operator
            string can be concatenated using + or += assignment operator.
    */
}

// Constants Study
func constantStudy() {
    /*
        Constants refers to immutable values.
        Syntax:
        const <name> = value
    */
    const pi float64 = 3.1415
    fmt.Println(pi)
}

// Escape Sequence Study
func escapeStudy() {
    /*
        Escape Sequence         Meaning
        \\                      \ character
        \'                      ' character
        \"                      " character
        \?                      ? character
        \a                      Alert or bell
        \b                      Backspace
        \f                      Form feed
        \n                      Newline
        \r                      Carriage return
        \t                      Horizontal tab
        \v                      Vertical tab
        \ooo                    Octal number of one to three digits
        \xhh...                 Hexadecimal number of one ot more digits
    */
}

// Decision Making and Loop Control
func sampleFunction() {
    var x = 25
    /*
        Decision Making
        Statement                   Description
        1. if statement
        if (condition) {
            ...
            // statement executed when condition is true
        }
        or
        if condition {
            ...
        }
    */
    x = 25
    if x > 10 {
        fmt.Println("x is greater than 10")
    }
    /*
        if condition1 && condition2 {
            ...
        }
    */
    if x > 10 && x < 30 {
        fmt.Println("x is greater than 10 and less than 30")
    }
    /*
        2. if ... else ...
        if condition {
            ...
            // statement executed when condition is true
        } else {
            ...
            // statement executed when condition is false
        }
    */
    x = 5
    if x > 10 {
        fmt.Println("x is greater than 10")
    } else {
        fmt.Println("x is less than 10")
    }
    /*
        if condition1 {
            ...
            // statement executed when condition is true
        } else if condition2 {
            ...
            // statement executed when condition is false
        } else {
            ...
            // statement executed when condition is false
        }
    */
    x = 30
    y := 30
    if x > y {
        fmt.Println("x is greater than y")
    } else if x < y {
        fmt.Println("x is less than y")
    } else {
        fmt.Println("x is equal to y")
    }
    /*
        3. nested if statement
        if condition1 {
            if condition2 {
                ...
            } else {
                ...
            }
        }
        4. if with short statement
        if variable_declaration; condition {
            ...
        }
    */
    if z := 20; z%2 == 0 {
        fmt.Println("z is even")
    } else {
        fmt.Println("z is odd")
    }
    /*
        5. switch statement
        switch <expression> {
        case <val1>:
            ...
            // Block executed when expression = val1;
        case <val2>:
            ...
            // Block executed when expression = val2;
        default:
            // Optional
            // Block executed when expression is not equal to val1 and val2
        }
    */
    var dayOfWeek = 5
    switch dayOfWeek {
    case 1:
        fmt.Println("Today is Monday")
    case 2:
        fmt.Println("Today is Tuesday")
    case 3:
        fmt.Println("Today is Wednesday")
    case 4:
        fmt.Println("Today is Thursday")
    case 5:
        fmt.Println("Today is Friday")
    case 6:
        fmt.Println("Today is Saturday")
    case 7:
        fmt.Println("Today is Sunday")
    }
    switch dayOfWeek {
    case 1, 2, 3, 4, 5:
        fmt.Println("It's a weekday")
    case 6, 7:
        fmt.Println("It's a weekend")
    }
    /*
        switch fallthrough statement allows to execute all following statement after a match found
    */
    switch dayOfWeek {
    case 1:
        fmt.Println("Today is Monday")
    case 2:
        fmt.Println("Today is Tuesday")
    case 3:
        fmt.Println("Today is Wednesday")
    case 4:
        fmt.Println("Today is Thursday")
    case 5:
        fmt.Println("Today is Friday")
        fallthrough
    case 6:
        fmt.Println("Today is Saturday")
    case 7:
        fmt.Println("Today is Sunday")
    }
    /*
        switch also support short statement
        switch variable_declaration; variable {
        case <val1>:
            ...
        }

        switch with no expression
        switch {
            case <boolean_expression>:
                ...
            case <boolean_expression>:
                ...
        }
    */
    dayOfWeek = 4
    switch {
    case dayOfWeek == 1:
        fmt.Println("Today is Monday")
    case dayOfWeek == 2:
        fmt.Println("Today is Tuesday")
    case dayOfWeek == 3:
        fmt.Println("Today is Wednesday")
    case dayOfWeek == 4:
        fmt.Println("Today is Thursday")
    case dayOfWeek == 5:
        fmt.Println("Today is Friday")
    case dayOfWeek == 6:
        fmt.Println("Today is Saturday")
    case dayOfWeek == 7:
        fmt.Println("Today is Sunday")
    }
    /*
        Loop
            1. for loop
            for initialization; condition; increment {
                // loop body
            }
            Property            Description
            initialization      Sets a counter variable
            condition           It test the each loop iteration for a condition. If it returns TRUE,
                                the loop continues. If it returns FALSE, the loop ends.
            increment           Increment counter variable
    */
    sum := 0
    for i := 1; i < 5; i++ {
        sum += i
    }
    fmt.Println(sum)
    /*
            while loop
            for condition {
                increment statement
            }
    */
    power := 1
    for power < 5 {
        power *= 2
    }
    fmt.Println(power)
    /*
            2. For Range
            for key, value := range collection {
                //block of statements
            }
    */
    langs := []string{"Go", "C", "C++", "Java"}
    for i, s := range langs {
        fmt.Println(i, s)
    }
    for i, s := range "Apple" {
        fmt.Printf("%U represents %c and it is at position %d\n", s, s, i)
    }
    /*
        The first value is the starting byte index of the rune and the second the rune itself.
    */
    fruits := map[string]string{"A": "Apple", "B": "Banana", "C": "Cherry"}
    for key, value := range fruits {
        fmt.Printf("%s -> %s\n", key, value)
    }

    for key := range fruits {
        fmt.Println("key:", key)
    }
    /*
        Loop Control
        1. break statement
    */
    var count = 0
outer:
    for count <= 10 {
        count++
        if count == 5 {
            break outer
        }
        fmt.Printf("Inside loop %d\n", count)
    }
    fmt.Println("Out of loop")
    /*
            2. continue statement
            for <condition1> {
                // loop body
                if (condition2) {
                    continue
                }
                // loop body
            }
    */
    var ctr = 0
    for ctr < 10 {
        ctr++
        if ctr == 5 {
            println("5 is skipped")
            continue
            // println("This won't be printed too.")
        }
        fmt.Printf("Number is %d\n", ctr)

    }
    fmt.Println("Out of loop")

    var i = 0
    var j = 0
outerfor:
    for i < 3 {
        j = 0
        i++
        for j < 3 {
            j++
            fmt.Printf("i is %d and j is %d\n", i, j)
            if i == 2 {
                fmt.Println("I am Skipped")
                continue outerfor
            }
            fmt.Println("I am Printed")
        }

    }
    /*
        3. goto statement
    */
    var age int
election:
    fmt.Print("Enter your age ")
    _, _ = fmt.Scanln(&age)
    if age <= 17 {
        fmt.Println("You are not eligible to vote.")
        goto election
    } else {
        fmt.Println("You are eligible to vote.")
    }
}

/*
    Function Study
    func function_name( [parameter list] ) [return_types] {
        body of the function
    }
    Returning multiple values from Function
    func function_name( [parameter list] ) [return_types] {
        return var1, var2
    }
    var1, var2 = function_name()

    Call Type                       Description
    Call by value                   This method copies the actual value of an argument into the formal parameter of the
                                    function. In this case, changes made to the parameter inside the function have no
                                    effect on the argument.
    Call by reference               This method copies the address of an argument into the formal parameter. Inside the
                                    function, the address is used to access the actual argument used in the call. This
                                    means that changes made to the parameter affect the argument.
    By default, Go uses call by value to pass arguments.
    func swap(x *int, y *int) {
        var temp int
        temp = *x
        *x = *y
        *y = temp
    }
    swap(&a, &b)

    Function Usage                  Description
    Function as Value               Functions can be created on the fly and can be used as values.
    Function Closures               Functions closures are anonymous functions and can be used in dynamic programming.
    Method                          Methods are special functions with a receiver.

    Function as Value
        getSquareRoot := func(x float64) float64 {
            return math.Sqrt(x)
        }
        fmt.Println(getSquareRoot(9))

    Function Closures
        func getSequence() func() int {
           i:=0
           return func() int {
              i+=1
              return i
           }
        }
        nextNumber := getSequence()
        fmt.Println(nextNumber())       ==> 1
        fmt.Println(nextNumber())       ==> 2

    Method
        func (variable_name variable_data_type) function_name() [return_type]{
            function body
        }

        1. Define a circle
        type Circle struct {
           x,y,radius float64
        }
        2. Define a method for circle
        func(circle Circle) area() float64 {
            return math.Pi * circle.radius * circle.radius
        }
        3. Call
        circle := Circle{x:0, y:0, radius:5}
        fmt.Printf("Circle area: %f", circle.area())
*/
/* function returning the max between two numbers */
func max(num1, num2 int) int {
    /* local variable declaration */
    var result int
    if num1 > num2 {
        result = num1
    } else {
        result = num2
    }
    return result
}
func functionStudy() {
    // Calling a Function
    var num1 = 10
    var num2 = 20
    var result = max(num1, num2)
    fmt.Println(result)
}

// Scope Rules Study
func scopeStudy() {
    /*
        Inside a function or a block (local variables)
        Outside of all functions (global variables)
        In the definition of function parameters (formal parameters)

        Local Variables
            Variables that are declared inside a function or a block are called local variables.
            They can be used only by statements that are inside that function or block of code.
        Global Variables
            Global variables are defined outside of a function, usually on top of the program.
            A global variable can be accessed by any function.

            Global variable declaration
            var g int

            func main() {
                g = 20
                fmt.Println(g)
            }

            Global variable declaration
            var g int = 20

            func main() {
                Local variable declaration
                var g int = 10

                fmt.Printf ("value of g = %d\n",  g)
            }
        Formal Parameters
            Formal parameters are treated as local variables with-in that function and they take preference over
            the global variables.
    */
}

// Strings Study
func stringStudy() {
    /*
            In the Go programming language, strings are slices.
            A string literal holds a valid UTF-8 sequences called runes. A String holds arbitrary bytes.

            Creating Strings
            var greeting = "Hello world!"

            Get Length
            len(greeting)

            Concatenating Strings
            greetings :=  []string{"Hello","world!"}
            fmt.Println(strings.Join(greetings, " "))
    */
}

// Arrays Study
func arrayStudy() {
    /*
        Store a fixed-size sequential collection of elements of the same type.
        An array is used to store a collection of data, but it is often more useful to think of an array as a
        collection of variables of the same type.

        Declaring Arrays
        var variable_name [SIZE] variable_type

        Initializing Arrays
        var variable_name [SIZE] variable_type{initial_val...}
        or
        var variable_name [] variable_type{initial_val...}

        Accessing Array Elements
        variable_name[index]

        Concept                             Description
        Multi-dimensional arrays            Go supports multidimensional arrays. The simplest form of a multidimensional
                                            array is the two-dimensional array.
            Syntax:
            var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
            Two-Dimensional Arrays
            var arrayName [ x ][ y ] variable_type
            Initializing Two-Dimensional Arrays
            a = [3][4]int{
                {0, 1, 2, 3} ,
                {4, 5, 6, 7} ,
                {8, 9, 10, 11}
            }
            Accessing Two-Dimensional Array Elements
            int val = a[2][3]

        Passing arrays to functions         You can pass to the function a pointer to an array by specifying the array's
                                            name without an index.
        Way-1
            void myFunction(param [10]int) {
                ...
            }
        Way-2
            void myFunction(param []int) {
                ...
            }

            func getAverage(arr []int, int size) float32 {
               var i int
               var avg, sum float32

               for i = 0; i < size; ++i {
                  sum += arr[i]
               }

               avg = sum / size
               return avg;
            }
    */
}

// Pointer Study
func pointerStudy() {
    /*
        Pointers
        Every variable is a memory location and every memory location has its address defined which can be accessed
        using ampersand (&) operator, which denotes an address in memory.

        Pointer Variable Declaration
        var var_name *var-type
    */
    //var a int = 10
    //fmt.Printf("Address of a variable: %x\n", &a  )

    var a int = 20 /* actual variable declaration */
    var ip *int    /* pointer variable declaration */

    ip = &a /* store address of a in pointer variable*/

    fmt.Printf("Address of a variable: %x\n", &a)

    /* address stored in pointer variable */
    fmt.Printf("Address stored in ip variable: %x\n", ip)

    /* access the value using the pointer */
    fmt.Printf("Value of *ip variable: %d\n", *ip)

    var ptr *int // Nil Pointers

    fmt.Printf("The value of ptr is : %x\n", ptr)

    if ptr != nil {
        fmt.Println("ptr is not NIL")
    } else {
        fmt.Println("ptr is NIL")
    }

    /*
        Array of pointers
        var ptr [MAX]*int

        Pointer to pointer
        A pointer to a pointer is a form of chain of pointers. Normally, a pointer contains the address of a variable.
        When we define a pointer to a pointer, the first pointer contains the address of the second pointer, which
        points to the location that contains the actual value

        var ptr **int;
    */
    var pptr **int

    /* take the address of var */
    ptr = &a

    /* take the address of ptr using address of operator & */
    pptr = &ptr

    /* take the value using pptr */
    fmt.Printf("Value of a = %d\n", a)
    fmt.Printf("Value available at *ptr = %d\n", *ptr)
    fmt.Printf("Value available at **pptr = %d\n", **pptr)
}

// Structures Study
func structureStudy() {
    /*
        type struct_variable_type struct {
           member definition;
           member definition;
           ...
           member definition;
        }

        variable_name := structure_variable_type {value1, value2...valuen}

        Accessing Structure Members
        To access any member of a structure, we use the member access operator (.). The member access operator is
        coded as a period between the structure variable name and the structure member that we wish to access. You
        would use struct keyword to define variables of structure type.
    */
}

// Slices Study
func sliceStudy() {
    // Slice is an abstraction over Go Array.
    // Array allows you to define variables that can hold several data items of the same kind but it does not provide
    //any inbuilt method to increase its size dynamically or get a sub-array of its own.

    // Defining a Slice
    var numbers []int /* a slice of unspecified size */
    /* numbers == []int{0,0,0,0,0}*/
    numbers = make([]int, 3, 5) /* a slice of length 5 and capacity 5*/
    fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
    // len() function returns the elements presents in the slice
    // cap() function returns the capacity of the slice

    /* create a slice */
    numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
    printSlice(numbers)

    /* print the original slice */
    fmt.Println("numbers ==", numbers)

    /* print the sub slice starting from index 1(included) to index 4(excluded)*/
    fmt.Println("numbers[1:4] ==", numbers[1:4])

    /* missing lower bound implies 0*/
    fmt.Println("numbers[:3] ==", numbers[:3])

    /* missing upper bound implies len(s)*/
    fmt.Println("numbers[4:] ==", numbers[4:])

    numbers1 := make([]int, 0, 5)
    printSlice(numbers1)

    /* print the sub slice starting from index 0(included) to index 2(excluded) */
    number2 := numbers[:2]
    printSlice(number2)

    /* print the sub slice starting from index 2(included) to index 5(excluded) */
    number3 := numbers[2:5]
    printSlice(number3)

    /* append allows nil slice */
    numbers = append(numbers, 10)
    printSlice(numbers)

    /* add one element to slice*/
    numbers = append(numbers, 11)
    printSlice(numbers)

    /* add more than one element at a time*/
    numbers = append(numbers, 12, 13, 14)
    printSlice(numbers)

    /* create a slice numbers1 with double the capacity of earlier slice*/
    numbers3 := make([]int, len(numbers), (cap(numbers))*2)

    /* copy content of numbers to numbers1 */
    copy(numbers3, numbers)
    printSlice(numbers3)

    // Range
    /* print the numbers */
    for i := range numbers {
        fmt.Println("Slice item", i, "is", numbers[i])
    }

    /* create a map*/
    countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

    /* print map using keys*/
    for country := range countryCapitalMap {
        fmt.Println("Capital of", country, "is", countryCapitalMap[country])
    }

    /* print map using key-value*/
    for country, capital := range countryCapitalMap {
        fmt.Println("Capital of", country, "is", capital)
    }
}
func printSlice(x []int) {
    fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}

// Maps Study
func mapStudy() {
    /*
        Go provides another important data type named map which maps unique keys to values. A key is an object that you
        use to retrieve a value at a later date. Given a key and a value, you can store the value in a Map object. After
        the value is stored, you can retrieve it by using its key.

        Declare a variable, by default map will be nil
        var map_variable map[key_data_type]value_data_type

        Define the map as nil map can not be assigned any value
        map_variable = make(map[key_data_type]value_data_type)
    */
    var countryCapitalMap map[string]string
    /* create a map*/
    countryCapitalMap = make(map[string]string)

    /* insert key-value pairs in the map*/
    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"

    /* print map using keys*/
    for country := range countryCapitalMap {
        fmt.Println("Capital of", country, "is", countryCapitalMap[country])
    }

    /* test if entry is present in the map or not*/
    capital, ok := countryCapitalMap["United States"]

    /* if ok is true, entry is present otherwise entry is absent*/
    if ok {
        fmt.Println("Capital of United States is", capital)
    } else {
        fmt.Println("Capital of United States is not present")
    }

    /* delete an entry */
    delete(countryCapitalMap, "France")
    fmt.Println("Entry for France is deleted")

    fmt.Println("Updated map")

    /* print map */
    for country := range countryCapitalMap {
        fmt.Println("Capital of", country, "is", countryCapitalMap[country])
    }
}

// Recursion Study
func recursionStudy() {
    /*
        Recursion is the process of repeating items in a self-similar way.

        func recursion() {
           recursion()
        }
    */
    var i int = 4
    fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

    for i = 0; i < 10; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
}
func factorial(i int) int {
    if i <= 1 {
        return 1
    }
    return i * factorial(i-1)
}
func fibonacci(i int) (ret int) {
    if i == 0 {
        return 0
    }
    if i == 1 {
        return 1
    }
    return fibonacci(i-1) + fibonacci(i-2)
}

// Type Casting Study
func typeCastingStudy() {
    /*
        Type Casting
    */
    var sum int = 17
    var count int = 5
    var mean float32

    mean = float32(sum) / float32(count)
    fmt.Printf("Value of mean : %f\n", mean)
}

// Interfaces Study
func interfaceStudy() {
    /*
        Define an interface
        type interface_name interface {
            method_name1 [return_type]
            method_name2 [return_type]
            method_name3 [return_type]
            ...
            method_namen [return_type]
        }

        Define a struct
        type struct_name struct {
        variables
        }

        Implement interface methods
        func (struct_name_variable struct_name) method_name1() [return_type] {
            method implementation
        }
        ...
        func (struct_name_variable struct_name) method_namen() [return_type] {
            method implementation
        }
    */
    circle := Circle{x: 0, y: 0, radius: 5}
    rectangle := Rectangle{width: 10, height: 5}

    fmt.Printf("Circle area: %f\n", getArea(circle))
    fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}

/* define an interface */
type Shape interface {
    area() float64
}

/* define a circle */
type Circle struct {
    x, y, radius float64
}

/* define a rectangle */
type Rectangle struct {
    width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rect Rectangle) area() float64 {
    return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
    return shape.area()
}

// Error Handling Study
func errorHandlingStudy() {
    /*
        type error interface {
           Error() string
        }
    */
    result, err := Sqrt(-1)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }

    result, err = Sqrt(9)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
func Sqrt(value float64) (float64, error) {
    if value < 0 {
        return 0, errors.New("Math: negative number passed to Sqrt")
    }
    return math.Sqrt(value), nil
}

// The main() function
func main() {
    // Local Variable Declarations
    //welcome()
    //dataTypeStudy()
    //variableStudy()
    //operatorStudy()
    //constantStudy()
    //escapeStudy()
    //sampleFunction()
    //functionStudy()
    //scopeStudy()
    //stringStudy()
    //arrayStudy()
    //pointerStudy()
    //sliceStudy()
    //mapStudy()
    //recursionStudy()
    //typeCastingStudy()
    //interfaceStudy()
    //errorHandlingStudy()

    /*
        How to print type of a variable in Go?
    */
    var a, b, c = 3, 4, "foo"
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("b is of type %T\n", c)
    /*
        The syntax of a for loop in Go programming language is
        for [condition |  ( init; condition; increment ) | Range] {
           statement(s);
        }
        The general form of a function definition in Go programming language is
        func function_name( [parameter list] ) [return_types] {
           body of the function
        }

        In how many ways you can pass parameters to a method?
        Call by value − This method copies the actual value of an argument into the formal parameter of the function.
                        In this case, changes made to the parameter inside the function have no effect on the argument.

        Call by reference − This method copies the address of an argument into the formal parameter. Inside the function,
                            the address is used to access the actual argument used in the call. This means that changes
                            made to the parameter affect the argument.

        What is the difference between actual and formal parameters?
        The parameters sent to the function at calling end are called as actual parameters while at the receiving of
        the function definition called as formal parameters.

        What is a token?
        A Go program consists of various tokens and a token is either a keyword, an identifier, a constant, a string
        literal, or a symbol.

        How to create a map in Go?
        Declare a variable, by default map will be nil
        var map_variable map[key_data_type]value_data_type

        Define the map as nil map can not be assigned any value
        map_variable = make(map[key_data_type]value_data_type)
    */
}
