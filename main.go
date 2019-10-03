package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/chneau/tt"
	"github.com/owulveryck/onnx-go/backend/simple"

	"github.com/owulveryck/onnx-go"
)

func init() {
	log.SetPrefix("[MNIST-TEST] ")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	defer tt.T()()
	// Create a backend receiver
	backend := simple.NewSimpleGraph()
	// Create a model and set the execution backend
	model := onnx.NewModel(backend)

	// read the onnx model
	b, _ := ioutil.ReadFile("py/model.onnx")
	// Decode it into the model
	err := model.UnmarshalBinary(b)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(model)
}
