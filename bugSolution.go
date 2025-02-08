func myFunc(a, b int) (int, error) {
  if a == 0 {
    return 0, fmt.Errorf("a cannot be zero")
  }
  if b == 0 {
    return 0, fmt.Errorf("b cannot be zero")
  }
  return a / b, nil
}

func main() {
  result, err := myFunc(10, 0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
  result, err = myFunc(10, 2)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
} 