package main

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

type EmbledScript struct {
	LuaState *lua.LState
}

func RunScript() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("example/linear.lua"); err != nil {
		panic(err)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("render"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10)); err != nil {
		panic(err)
	}
	ret := L.Get(-1) // returned value
	log.Println(ret.String())
	L.Pop(1)
}

func NewEmbledScript(scriptPath string) *EmbledScript {
	L := lua.NewState()
	if err := L.DoFile(scriptPath); err != nil {
		panic(err)
	}
	return &EmbledScript{
		LuaState: L,
	}
}
