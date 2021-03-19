# Study

## make & new

|        | new(T)            | make(T)               |
| ------ | ----------------- | --------------------- |
| 대상   | 임의의 타입       | slice, map, channel만 |
| 초기화 | 초기화하지 않는다 | 초기화 한다.          |
| 반환값 | \*T               | T                     |

## append

```go
append([]int{}, 1, 2)
// [1,2]
```