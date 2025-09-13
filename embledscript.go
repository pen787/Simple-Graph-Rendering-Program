package main

import (
	"image/color"
	"log"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

type EmbledScript struct {
	LuaState *lua.LState
	path     string
}

func (es *EmbledScript) CallLoad() (rl.Color, error) {
	if err := es.LuaState.CallByParam(lua.P{
		Fn:      es.LuaState.GetGlobal("load"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		return color.RGBA{}, err
	}
	ret := es.LuaState.Get(-1)

	R := es.LuaState.GetTable(ret, lua.LString("R"))
	RI, err := strconv.ParseInt(R.String(), 10, 16)
	if err != nil {
		return rl.Color{}, err
	}

	G := es.LuaState.GetTable(ret, lua.LString("G"))
	GI, err := strconv.ParseInt(G.String(), 10, 16)
	if err != nil {
		return rl.Color{}, err
	}
	B := es.LuaState.GetTable(ret, lua.LString("B"))
	BI, err := strconv.ParseInt(B.String(), 10, 16)
	if err != nil {
		return rl.Color{}, err
	}

	es.LuaState.Pop(1)

	return rl.NewColor(uint8(RI), uint8(GI), uint8(BI), uint8(255)), nil
}

func (es *EmbledScript) CallRender(x float32) (float32, error) {
	if err := es.LuaState.CallByParam(lua.P{
		Fn:      es.LuaState.GetGlobal("render"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(x)); err != nil {
		return 0, err
	}
	ret := es.LuaState.Get(-1)
	value, err := strconv.ParseFloat(ret.String(), 32)
	es.LuaState.Pop(1)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

func (es *EmbledScript) Close() {
	es.LuaState.Close()
	log.Println("LuaState closed!")
}

func (es *EmbledScript) DoFile(filepath string) error {
	if err := es.LuaState.DoFile(filepath); err != nil {
		return err
	}
	es.path = filepath
	return nil
}

func (es *EmbledScript) ResetScript() error {
	es.LuaState.Close()
	es.LuaState = lua.NewState(lua.Options{})
	if err := es.DoFile(es.path); err != nil {
		return err
	}
	return nil
}

func NewEmbledScript() *EmbledScript {
	L := lua.NewState(lua.Options{})
	return &EmbledScript{
		LuaState: L,
		path:     "",
	}
}
