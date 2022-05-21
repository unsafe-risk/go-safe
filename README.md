<h1 align="center">Go Safe</h1>

<div align="center">
  <img src="https://avatars.githubusercontent.com/u/87268138?s=200&v=4" alt="Unsafe Risk" />
</div>

<br>

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
</div>

## Description

Use nil safety value

## Install

```
go get github.com/unsafe-risk/go-safe
```

## Usage

```go
import (
  // ...
  safe "github.com/unsafe-risk/go-safe"
)
```

[https://pkg.go.dev/github.com/unsafe-risk/go-safe](https://pkg.go.dev/github.com/unsafe-risk/go-safe)

- TimeOrNow(t *time.Time) time.Time  
- StrOrZero(s *string) string  
- IOrZero(n *int) int  
- I8OrZero(n *int8) int8  
- I16OrZero(n *int16) int16  
- I32OrZero(n *int32) int32  
- I64OrZero(n *int64) int64  
- UiOrZero(n *uint) uint  
- Ui8OrZero(n *uint8) uint8  
- Ui16OrZero(n *uint16) uint16  
- Ui32OrZero(n *uint32) uint32  
- Ui64OrZero(n *uint64) uint64  
- F32OrZero(n *float32) float32  
- F64OrZero(n *float64) float64  

---

- TimeOrDefault(t *time.Time, def func() time.Time) time.Time  
- StrOrDefault(s *string, def func() string) string  
- IOrDefault(n *int, def func() int) int  
- I8OrDefault(n *int8, def func() int8) int8  
- I16OrDefault(n *int16, def func() int16) int16  
- I32OrDefault(n *int32, def func() int32) int32  
- I64OrDefault(n *int64, def func() int64) int64  
- UiOrDefault(n *uint, def func() uint) uint  
- Ui8OrDefault(n *uint8, def func() uint8) uint8  
- Ui16OrDefault(n *uint16, def func() uint16) uint16  
- Ui32OrDefault(n *uint32, def func() uint32) uint32  
- Ui64OrDefault(n *uint64, def func() uint64) uint64  
- F32OrDefault(n *float32, def func() float32) float32  
- F64OrDefault(n *float64, def func() float64) float64  

---

- StrZero() string  
- IZero() int  
- I8Zero() int8  
- I16Zero() int16  
- I32Zero() int32  
- I64Zero() int64  
- UiZero() uint  
- Ui8Zero() uint8  
- Ui16Zero() uint16  
- Ui32Zero() uint32  
- Ui64Zero() uint64  
- F32Zero() float32  
- F64Zero() float64  