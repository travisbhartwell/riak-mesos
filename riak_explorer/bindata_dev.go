// +build !rel

//go:generate go-bindata -ignore=Makefile -o bindata_generated.go -pkg=riak_explorer -prefix=data/riak_explorer -debug data/
package riak_explorer
