package main

import "fmt"
import "one"

type TargetClass struct {
    InstanceField int
}

func NewTargetClass() *TargetClass {
    this := new(TargetClass)
    this.InstanceField = 5
    return this
}

func TargetClass_StaticMethod(arg1 string) string {
    return fmt.Sprintf("arg1 = %v, staticField = %v", arg1, TargetClassStaticField)
}

func (this *TargetClass) InstanceMethod() string {
    return fmt.Sprintf("instanceField = %v", this.InstanceField)
}

var TargetClassStaticField string = "hello";

type TestClass struct {
}

func NewTestClass() *TestClass {
    this := new(TestClass)
    return this
}

func (this *TestClass) TestMethod() {
    obj := NewTargetClass()
    //console.log(`instanceMethod (direct): ${obj.instanceMethod()}`);
    //console.log(`staticMethod (direct): ${TargetClass.staticMethod("arg1value")}`);
    //console.log(`instanceField (direct): ${obj.instanceField}`);
    //console.log(`staticField (direct): ${TargetClass.staticField}`);
    cls := one.Reflect_GetClass(obj)
    if cls == nil {
        fmt.Println("cls is null!")
        return
    }
    cls2 := one.Reflect_GetClassByName("TargetClass")
    if cls2 == nil {
        fmt.Println("cls2 is null!")
        return
    }
    
    method1 := cls.GetMethod("instanceMethod")
    if method1 == nil {
        fmt.Println("method1 is null!")
        return
    }
    method1_result := method1.Call(obj, []interface{}{})
    fmt.Println(fmt.Sprintf("instanceMethod: %v", method1_result))
    
    method2 := cls.GetMethod("staticMethod")
    if method2 == nil {
        fmt.Println("method2 is null!")
        return
    }
    method2_result := method2.Call(nil, []interface{}{"arg1value"})
    fmt.Println(fmt.Sprintf("staticMethod: %v", method2_result))
    
    field1 := cls.GetField("instanceField")
    if field1 == nil {
        fmt.Println("field1 is null!")
        return
    }
    field1.SetValue(obj, 6)
    field1_new_val := field1.GetValue(obj)
    fmt.Println(fmt.Sprintf("new instance field value: %v == %v", obj.InstanceField, field1_new_val))
    
    field2 := cls.GetField("staticField")
    if field2 == nil {
        fmt.Println("field2 is null!")
        return
    }
    field2.SetValue(nil, "bello")
    field2_new_val := field2.GetValue(nil)
    fmt.Println(fmt.Sprintf("new static field value: %v == %v", TargetClassStaticField, field2_new_val))
}

func init() {
}

func main() {
    defer func() {
      if r := recover(); r != nil {
          fmt.Print("Exception: ", r)
      }
    }()

    c := NewTestClass()
    c.TestMethod();
}