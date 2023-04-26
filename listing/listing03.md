Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

В первом случае выведется nil, так как значение интерфейса равно nil.
В втором случае выведется false, так как интерфейс, содержащий нулевое значение
не является нулевым, его значение nil, но его тип *fs.PathError.

```
