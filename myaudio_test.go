package main

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GetRawData(t *testing.T) {
	GetRawData(getAudioSeg("test_data/312.mp3"))
}

func Test_GetWaveData(t *testing.T) {
	GetWaveData("test_data/312.mp3")
}

func Test_GetOsType(t *testing.T) {
	osType := runtime.GOOS
	fmt.Println("OS Type:", osType)
}

func Test_bytes2Int(t *testing.T) {
	var mySlice = []byte{1, 2}
	num := int(binary.BigEndian.Uint16(mySlice))
	if diff := cmp.Diff(258, num); diff != "" {
		t.Error(diff)
	}
	num2 := int(binary.LittleEndian.Uint16(mySlice))
	if diff := cmp.Diff(513, num2); diff != "" {
		t.Error(diff)
	}
}
