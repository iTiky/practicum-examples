package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

type StructObj struct {
	value int
}

func main() {
	debug.SetGCPercent(-1)

	objs := []*StructObj{
		{value: 1},
		{value: 2},
	}

	fmt.Println("1st (initial objs slice):")
	runtime.GC()
	fmt.Printf("  objs (addrs): %v\n", objs)

	fmt.Println("2nd (shrunk objs):")
	objs = objs[:0]
	runtime.GC()
	fmt.Printf("  objs (addrs): %v\n", objs)

	fmt.Println("3rd (restore objs):")
	objs = objs[:2]
	runtime.GC()
	fmt.Printf("  objs (addrs): %v\n", objs)

	fmt.Println("4th (resize objs):")
	objsResized := append(objs, &StructObj{value: 3})
	runtime.GC()
	fmt.Printf("  objs (addrs):         %v\n", objs)
	fmt.Printf("  objsResized (addrs):  %v\n", objsResized)

	fmt.Println("5th (delete 1st obj):")
	objsResized[0] = nil
	runtime.GC()
	fmt.Printf("  objs (addrs):         %v (%d / %d)\n", objs, len(objs), cap(objs))
	fmt.Printf("  objsResized (addrs):  %v (%d / %d)\n", objsResized, len(objsResized), cap(objsResized))

	fmt.Println("6th (delete objs slice):")
	objs = nil
	runtime.GC()
	fmt.Printf("  objs (addrs):         %v (%d / %d)\n", objs, len(objs), cap(objs))
	fmt.Printf("  objsResized (addrs):  %v (%d / %d)\n", objsResized, len(objsResized), cap(objsResized))

	fmt.Println("7th (delete objsResized slice saving pointer to the 1st obj):")
	obj1Ptr := uintptr(unsafe.Pointer(objsResized[1]))
	fmt.Printf("  objsResized[1]: uintptr before GC: %x (%+v)\n", obj1Ptr, (*StructObj)(unsafe.Pointer(obj1Ptr)))
	objsResized = nil
	runtime.GC()
	fmt.Printf("  objsResized[1]: uintptr after GC:  %x (%+v)\n", obj1Ptr, (*StructObj)(unsafe.Pointer(obj1Ptr)))
	fmt.Printf("  objsResized (addrs):               %v (%d / %d)\n", objsResized, len(objsResized), cap(objsResized))
	fmt.Println("  NOTES:")
	fmt.Println("    * we use uintptr as Go knows about the unsafe.Pointer ones and won't clear them up;")
	fmt.Println("    * uintptr after GC: that one can be either 0 / 2 from run to run")

	fmt.Println("8th (initial objsNew slice):")
	objsNew := []*StructObj{
		{value: 3},
	}
	fmt.Printf("  objsNew (addrs): %v (%d / %d)\n", objsNew, len(objsNew), cap(objsNew))

	fmt.Println("9th (delete objsNew 0th obj saving pointer to it):")
	objNew0Ptr := uintptr(unsafe.Pointer(objsNew[0]))
	fmt.Printf("  objsNew[0]: uintptr before GC: %x (%+v)\n", objNew0Ptr, (*StructObj)(unsafe.Pointer(objNew0Ptr)))
	objsNew[0] = nil
	runtime.GC()
	fmt.Printf("  objsNew[0]: uintptr after GC:  %x (%+v)\n", objNew0Ptr, (*StructObj)(unsafe.Pointer(objNew0Ptr)))
	fmt.Printf("  objsNew (addrs):               %v (%d / %d)\n", objsNew, len(objsNew), cap(objsNew))
	fmt.Println("  NOTES:")
	fmt.Println("    * uintptr after GC: that one can be 0 / 3 / random  from run to run")
}
