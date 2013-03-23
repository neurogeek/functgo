package functgo

import (
    "fmt"
    "reflect"
)

func FFilter(arr interface{}, fn reflect.Value) interface{} {
    arrVal := reflect.Indirect(reflect.ValueOf(arr))
    res := reflect.MakeSlice(arrVal.Type(), arrVal.Len(), arrVal.Cap())

    j := 0

    for i := 1; i < arrVal.Len(); i++ {

        a := fn.Call([]reflect.Value{arrVal.Index(i)})[0]

        if a.Bool() == true {
            res.Index(j).Set(arrVal.Index(i))
            j++
        }
    }

    return res.Slice(0, (j)).Interface()
}

func FReduce(arr interface{}, fn reflect.Value) interface{} {
    arrVal := reflect.Indirect(reflect.ValueOf(arr))
    for i := 1; i < arrVal.Len(); i++ {

        a := []reflect.Value{arrVal.Index(0), arrVal.Index(i)}
        arrVal.Index(0).Set(fn.Call(a)[0])
    }

    return arrVal.Index(0).Interface()
}

func FMap(arr interface{}, fn reflect.Value) interface{} {

    arrVal := reflect.Indirect(reflect.ValueOf(arr))
    res := reflect.MakeSlice(arrVal.Type(), arrVal.Len(), arrVal.Cap())

    for i := 0; i < arrVal.Len(); i++ {
        r := fn.Call([]reflect.Value{arrVal.Index(i)})[0]
        res.Index(i).Set(r)
    }

    return res.Interface()
}
