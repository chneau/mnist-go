package main

import (
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

type Core struct {
	*gorgonnx.Graph
	*onnx.Model
}

func NewCore(model []byte) (*Core, error) {
	c := &Core{
		Graph: gorgonnx.NewGraph(),
	}
	c.Model = onnx.NewModel(c.Graph)
	err := c.Model.UnmarshalBinary(model)
	if err != nil {
		return nil, err
	}
	return c, nil
}
