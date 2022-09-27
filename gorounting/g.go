package main

import "context"

func main() {

}

func Speck(ctx context.Context) error {
	go func() {
		<-ctx.Done()

	}()
	for i := 0; i < 10; i++ {

	}
	return nil
}

func Duck(ctx context.Context) error {
	return nil
}
