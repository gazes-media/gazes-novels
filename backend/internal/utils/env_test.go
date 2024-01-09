package utils

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetenv(t *testing.T) {
	// Test case 1: All environment variables are present
	keys1 := []string{"KEY1", "KEY2", "KEY3"}
	expectedValues1 := map[string]string{"KEY1": "value1", "KEY2": "value2", "KEY3": "value3"}

	// Set environment variables for test case 1
	for key, value := range expectedValues1 {
		os.Setenv(key, value)
	}
	defer unsetEnvVars(keys1)

	result1, err1 := GetEnv(keys1)

	if err1 != nil {
		t.Errorf("Unexpected error in TestGetenv (case 1): %v", err1)
	}

	if !reflect.DeepEqual(result1, expectedValues1) {
		t.Errorf("Unexpected result in TestGetenv (case 1). Expected %v, got %v", expectedValues1, result1)
	}

	// Test case 2: Some environment variables are missing
	keys2 := []string{"KEY1", "MISSING_KEY", "KEY3"}
	expectedValues2 := map[string]string{"KEY1": "value1", "KEY3": "value3"}
	missingKeys2 := []string{"MISSING_KEY"}

	// Set environment variables for test case 2
	for key, value := range expectedValues2 {
		os.Setenv(key, value)
	}
	defer unsetEnvVars(keys2)

	result2, err2 := GetEnv(keys2)

	if err2 == nil {
		t.Error("Expected error in TestGetenv (case 2), but got nil")
	}

	if err2.Error() != fmt.Sprintf("missing environment variables: %v", missingKeys2) {
		t.Errorf("Unexpected error message in TestGetenv (case 2). Expected %v, got %v", fmt.Sprintf("missing environment variables: %v", missingKeys2), err2.Error())
	}

	if !reflect.DeepEqual(result2, expectedValues2) {
		t.Errorf("Unexpected result in TestGetenv (case 2). Expected %v, got %v", expectedValues2, result2)
	}
}

func unsetEnvVars(keys []string) {
	for _, key := range keys {
		os.Unsetenv(key)
	}
}
