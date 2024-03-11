package main

import (
	"fmt"
	"syscall/js"
)

//	func (a *App) Greet(name string) string {
//		return fmt.Sprintf("Hello %s, It's show time!", name)
//	}
func greet(name string) (string, error) {
	return fmt.Sprintf("Hello %s, It's show time!", name), nil
}

func greetWrapper() js.Func {
	greetFunction := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "Invalid no of arguments passed",
			}
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		greetingOuput := jsDoc.Call("getElementById", "greeting")
		if !greetingOuput.Truthy() {
			result := map[string]any{
				"error": "Unable to get greeting output element",
			}
			return result
		}
		name := args[0].String()
		fmt.Printf("input %s\n", name)
		greeting, err := greet(name)
		if err != nil {
			errStr := fmt.Sprintf("Problem with greet(). Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		greetingOuput.Set("innerText", greeting)
		return nil
	})
	return greetFunction
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("greetPerson", greetWrapper())
	<-make(chan struct{})
}
