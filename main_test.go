package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }   
    
    result1 := Add(2, 3)
    expected1 := 5
    if result1 != expected1 {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }    
}

func TestAdd1(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }   
    
    result1 := Add(2, 3)
    expected1 := 6
    if result1 != expected1 {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }    
}
