package main

import (
	"reflect"
	"testing"
)

func Test_print(t *testing.T) {
	type args struct {
		item product
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test-1", args{product{"0", "fan"}}},
		{"Test-2", args{product{"1", "tubelight"}}},
		{"Test-3", args{product{"2", "bulb"}}},
		{"Test-4", args{product{"3", "wire"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print(tt.args.item)
		})
	}
}

func Test_productList_get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		products productList
		args     args
		want     product
		want1    int
	}{
		{"Test-1", productList{product{"0", "fan"}, product{"3", "wire"}}, args{"3"}, product{"3", "wire"}, 1},
		{"Test-2", productList{product{"1", "AC"}, product{"3", "wire"}, product{"2", "Stablizer"}}, args{"3"}, product{"3", "wire"}, 1},
		{"Test-2", productList{product{"1", "AC"}, product{"3", "wire"}, product{"2", "Stablizer"}}, args{"1"}, product{"1", "AC"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.products.get(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productList.get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("productList.get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_productList_getAll(t *testing.T) {
	tests := []struct {
		name     string
		products productList
	}{
		{"Test-1", productList{product{"1", "AC"}, product{"3", "wire"}, product{"2", "Stablizer"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.products.getAll()
		})
	}
}

func Test_productList_add(t *testing.T) {
	type args struct {
		item product
	}
	tests := []struct {
		name     string
		products *productList
		args     args
	}{
		{"Test-1", &productList{product{"1", "AC"}, product{"3", "wire"}, product{"2", "Stablizer"}}, args{product{"0", "fan"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.products.add(tt.args.item)
		})
	}
}

func Test_productList_update(t *testing.T) {
	type args struct {
		updatedEntry product
		index        int
	}
	tests := []struct {
		name     string
		products *productList
		args     args
	}{
		{"Test-1", &productList{product{"1", "AC"}, product{"3", "wire"}, product{"4", "Stablizer"}}, args{product{"4", "fan"}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.products.update(tt.args.updatedEntry, tt.args.index)
		})
	}
}

func Test_productList_delete(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name     string
		products *productList
		args     args
	}{
		{"Test-1", &productList{product{"1", "AC"}, product{"3", "wire"}, product{"2", "Stablizer"}}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.products.delete(tt.args.index)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Run Tests"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
