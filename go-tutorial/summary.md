
---
# 21.31
- 基本类型
    - int	整型，根据平台可能是32或64位
    - float	浮点型，float32和float64
    - bool	布尔型，true或false
    - complex	复数，complex64和complex128

    - byte	等价 uint8 可以表达ANSCII字符. 0-255
        因为byte可以表达AnSIIC字符，所以可以用来强调一个值是字符类型
        var c byte = 'c' 完全是可以的

    - rune	等价 int32 可以表达Unicode字符. 0-0x10FFFF
    - string	字符串即字节序列，可以转换为[]byte类型即字节切片

- 衍生类型
    - 数组（array）[5]int，长度为5的数组。长度是数组的一部分，所以[5]int和[10]int是不同的类型。
        数组是值类型，赋值和传参会复制整个数组。一般很少直接使用数组。
    - 切片（slice）[]int，长度不固定的数组

    - iota一个特殊的常量生成器，用于在常量声明块中自动生成递增的整数值。它常用于定义枚举类型、位掩码等场景。
    ```go
    const (
        A = iota
        B
        C
    )

    const (
        KB = 1 << (10 * iota)
        MB
        GB
    )
    ```

- 常量（字面量）const
  - 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
    常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。常量可以是字符、字符串、布尔值或数值。

- 变量 var

- 条件控制
  - if switch select