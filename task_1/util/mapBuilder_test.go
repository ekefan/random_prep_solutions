package util

import (
	"maps"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestBuildNestedMap(t *testing.T) {
	randomJson := randomEntries(4)
	args := []string{"currency"}
	newMap := BuildNestedMap(randomJson, args)
	require.NotEmpty(t, newMap)
	for entry := range maps.Values(newMap) {
		require.NotEmpty(t, entry)
	}
	
	
}

func randomStr(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	builder := strings.Builder{}
	for range length {
		char := chars[rand.Intn(len(chars))]
		builder.WriteByte(char)
	}
	return builder.String()
}

func randomInt(max, min int64) int64 {
	return rand.Int63n((max - min) + min)
}

func randomEntries(numOfEntries int) []map[string]interface{} {
	jsonArray := []map[string]interface{}{}
	country := "country"
	city := "city"
	currency := "currency"
	amount := "amount"
	for range numOfEntries  {
		entry := map[string]interface{}{}
		entry[country] = randomStr(3)
		entry[city] = randomStr(4)
		entry[currency] = randomStr(2)
		entry[amount] = randomInt(150, 50)
		jsonArray = append(jsonArray, entry)
	}
	return jsonArray
}