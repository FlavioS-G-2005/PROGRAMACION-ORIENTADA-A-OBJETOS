package main

import "fmt"
type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}
func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}
func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d\n", p.ID, p.Nombre, p.Precio, p.Cantidad)
}
type Inventario struct {
	Productos map[int]Producto
}

func (inv *Inventario) AgregarProducto(p Producto) {
	inv.Productos[p.ID] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	delete(inv.Productos, id)
}

func (inv Inventario) BuscarProducto(id int) (Producto, bool) {
	prod, existe := inv.Productos[id]
	return prod, existe
}

func (inv Inventario) ListarProductos() {
	for _, p := range inv.Productos {
		p.MostrarInfo()
	}
}

func main() {
	inv := Inventario{
		Productos: make(map[int]Producto),
	}
	p1 := NuevoProducto(1, "Pan", 0.50, 20)
	p2 := NuevoProducto(2, "Leche", 1.20, 10)
	inv.AgregarProducto(p1)
	inv.AgregarProducto(p2)
	p1.ActualizarPrecio(0.60)
	p1.ActualizarCantidad(25)
	inv.AgregarProducto(p1) 
	fmt.Println("Inventario actual:")
	inv.ListarProductos()
	prod, encontrada := inv.BuscarProducto(1)
	if encontrada {
		fmt.Println("\nProducto encontrado:")
		prod.MostrarInfo()
	}
	inv.EliminarProducto(2)
	fmt.Println("\nLuego de eliminar la leche:")
	inv.ListarProductos()
}