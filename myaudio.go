package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"

	"github.com/hopesea/godub/v2"
	"github.com/hopesea/godub/v2/audioop"
)

func getAudioSeg(audioPath string) *godub.AudioSegment {
	filePath := path.Join(audioPath)
	segment, err := godub.NewLoader().Load(filePath)
	log.Println("================", segment)
	if err != nil {
		fmt.Println(err)
	}
	return segment
}
func GetRawData(segment *godub.AudioSegment) []int32 {
	sampleWitdh := int(segment.SampleWidth())
	//TODO: when channels are more than 1
	rawBytes := segment.RawData()
	rawData, _ := audioop.GetSamples(rawBytes, int(sampleWitdh))
	return rawData
}

func GetWaveData(audioPath string) []int32 {
	segment := getAudioSeg(audioPath)
	rawData := GetRawData(segment)
	frameCount := segment.FrameCount()
	frameRate := segment.FrameRate()
	timePerPoint := 0.01
	samplePerPoint := float64(frameRate) * timePerPoint
	totalPointNum := int(math.Ceil(frameCount / samplePerPoint))
	points := make([]int32, 0, totalPointNum)

	for i := 0; i < totalPointNum; i++ {
		start := i * int(samplePerPoint)
		end := start + int(samplePerPoint)
		if end > int(frameCount) {
			end = int(frameCount)
		}
		samplesPerPoint := rawData[start:end]
		sum := 0.
		for _, n := range samplesPerPoint {
			sum = sum + math.Abs(float64(n))
		}
		point := getMean(samplesPerPoint)
		points = append(points, point)
	}
	return points
}

func getMax(array []int32) int32 {
	max := float64(array[0])
	for _, n := range array {
		max = math.Max(max, math.Abs(float64(n)))
	}
	return int32(max)
}

func getMean(array []int32) int32 {
	sum := 0
	for _, n := range array {
		sum = sum + int(math.Abs(float64(n)))
	}
	return int32(sum / len(array))
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func toFile(data interface{}, destFile string) {
	f, err := os.Create(destFile)
	checkErr(err)
	defer f.Close()
	f.WriteString(fmt.Sprint(data))

}
