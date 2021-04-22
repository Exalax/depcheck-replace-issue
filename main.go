package main

import mw "github.com/grpc-ecosystem/go-grpc-middleware"

func main() {
	_ = mw.WrappedServerStream{}
}
